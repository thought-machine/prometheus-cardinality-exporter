package main

import (
	"context"
	"fmt"
	"github/thought-machine/prometheus-cardinality-exporter/cardinality"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/cenkalti/backoff"
	"github.com/jessevdk/go-flags"
	logging "github.com/sirupsen/logrus"
)

var log = logging.WithFields(logging.Fields{})

var opts struct {
	Selector              string   `long:"selector" short:"s" default:"app=prometheus" help:"Selector for Service Discovery."`
	Namespaces            []string `long:"namespaces" short:"n"  help:"Namespaces for Service Discovery."`
	PrometheusInstances   []string `long:"proms" short:"i" help:"Prometheus instance links. Mutually exclusive to the service discover flag."`
	PromAPIAuthValuesFile string   `long:"auth" short:"a" help:"Location of YAML file where Prometheus instance authorisation credentials can be found. For instances that don't appear in the file, it is assumed that no authorisation is required to access them."`
	ServiceDiscovery      bool     `long:"service_discovery" short:"d" help:"Service discovery flag, use service discovery to find new instances of Prometheus within a cluster. Mutually exclusive to the prometheus instance link flag."`
	Port                  int      `long:"port" short:"p" default:"9090" help:"Port on which to serve."`
	Frequency             float32  `long:"freq" short:"f" default:"6" help:"Frequency in hours with which to query the Prometheus API."`
	ServiceRegex          string   `long:"regex" short:"r" default:"prometheus-[a-zA-Z0-9_-]+" help:"If any found services don't match the regex, they are ignored."`
	LogLevel              string   `long:"log.level" short:"l" default:"info" help:"Level for logging. Options (in order of verbosity): [debug, info, warn, error, fatal]."`
	StatsLimit            int      `long:"stats-limit" short:"L" default:"10" help:"Limit the number of items fetched from the TSDB statistics."`
}

func collectMetrics() {

	// Number of times to retry before fetching the data before giving up.
	// If the number of retries is exhausted, it will wait until the next time it has to query the Prometheus API.
	var numRetries uint64 = 3
	sleepTime, err := time.ParseDuration(fmt.Sprintf("%0.4fh", opts.Frequency))
	if err != nil {
		log.Errorf("Cannot parse frequency variable %v: %v", opts.Frequency, err)
	}

	// Map of prometheus instance identifiers to their authorisation credentials, used for accessing the TSDB API
	var promAPIAuthValues map[string]string

	// This is a data structure that allows for the storage of the names prometheus instances and their sharded instances
	// Sharded instances are specified because a service may have several endpoints
	// Ignoring this would result in kubernetes selecting only one endpoint per API call, which could lead to inconsistent metric reporting
	// Each sharded instance also stores it's address (which can change), the latest cardinality info, and the current tracked labels
	cardinalityInfoByInstance := make(map[string]*cardinality.PrometheusCardinalityInstance)

	if opts.PromAPIAuthValuesFile != "" {
		filename, err := filepath.Abs(opts.PromAPIAuthValuesFile)
		if err != nil {
			log.Errorf("Failed to obtain the filepath of the Prometheus API authorisation values file provided: %v.", err.Error())
		} else {
			fileContents, err := os.ReadFile(filename)
			if err != nil {
				log.Errorf("Failed to read Prometheus API authorisation values file provided: %v.", err.Error())
			} else {
				err = yaml.Unmarshal(fileContents, &promAPIAuthValues)
				if err != nil {
					log.Errorf("Failed to read Prometheus API authorisation values file into the appropriate data structure: %v. Check the format of your file!", err.Error())
				}
			}
		}
		if len(promAPIAuthValues) == 0 {
			log.Errorf("Skipping the authorisation component to continue collecting metrics from Prometheus instances that don't require authorisation. This will result in no metrics from secured Prometheus instances.")
		}
	}

	if !opts.ServiceDiscovery { // Prometheus instances defined by arguments

		// precompile required regex
		regexC, err := regexp.Compile(`https?:\/\/[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+[a-zA-Z0-9_.-]*\/?`)
		if err != nil {
			log.Fatalf("invalid regex: %+v", err)
		}

		// In this case the name of the sharded instance is the same as the name of the prometheus instance
		// This is because is not possible to distinguish between them based on addresses given as arguments
		for _, prometheusInstanceAddress := range opts.PrometheusInstances {

			// Check the address matches a familiar pattern: http(s)://<instance name>.<anything else>(/)
			matched := regexC.MatchString(prometheusInstanceAddress)
			if !matched {
				log.Fatalf("%v is not a valid prometheus instance address.", prometheusInstanceAddress)
			}

			// Get the name of the prometheus instance from the link
			splitByDots := strings.Split(prometheusInstanceAddress, ".")
			splitInstanceName := strings.Split(splitByDots[0], "/")
			instanceName := splitInstanceName[len(splitInstanceName)-1]
			namespace := splitByDots[1]

			instanceID := namespace + "_" + instanceName

			// Add the prometheus instance to the data structure
			cardinalityInfoByInstance[instanceID] = &cardinality.PrometheusCardinalityInstance{
				Namespace:           namespace,
				InstanceName:        instanceName,
				ShardedInstanceName: instanceName,
				InstanceAddress:     prometheusInstanceAddress,
				AuthValue:           promAPIAuthValues[prometheusInstanceAddress],
				TrackedLabels: cardinality.TrackedLabelNames{
					SeriesCountByMetricNameLabels:     make([]string, 0, opts.StatsLimit),
					LabelValueCountByLabelNameLabels:  make([]string, 0, opts.StatsLimit),
					MemoryInBytesByLabelNameLabels:    make([]string, 0, opts.StatsLimit),
					SeriesCountByLabelValuePairLabels: make([]string, 0, opts.StatsLimit),
				},
			}
		}
	}

	for {
		if opts.ServiceDiscovery {

			// Obtains the cluster config of the cluster we are currently in
			config, err := rest.InClusterConfig()
			if err != nil {
				log.Fatalf("Error obtaining the current cluster config: %v", err.Error())
			}

			// Creates the clientset
			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				log.Fatalf("Error creating the clientset from the cluster config: %v", err.Error())
			}

			// If namespaces are specified as arguments use them, if not use service discovery
			var namespaceList []string
			if len(opts.Namespaces) == 0 {
				// Accesses the API to list all namespaces in the cluster
				namespaces, _ := clientset.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
				for _, namespaceObj := range namespaces.Items {
					namespaceList = append(namespaceList, namespaceObj.ObjectMeta.GetName())
				}
			} else {
				namespaceList = opts.Namespaces
			}

			for _, namespace := range namespaceList {

				// Accesses the API to list all endpoints and services which match the label selector in the given namespace
				endpointsList, _ := clientset.CoreV1().Endpoints(namespace).List(context.TODO(), v1.ListOptions{LabelSelector: opts.Selector})

				if err != nil {
					log.Fatalf("Error obtaining endpoints matching selector (%v) in namespace (%v): %v", namespace, opts.Selector, err.Error())
				}

				// Iterate over all of the endpoints and add them to the data structure
				for _, endpoints := range endpointsList.Items { // This loop represents a service

					prometheusInstanceName := endpoints.ObjectMeta.GetName()
					//If the instance name doesn't start with the chosen prefix, it is ignored
					if matched, _ := regexp.MatchString(opts.ServiceRegex, prometheusInstanceName); !matched {
						continue
					}

					for _, endpointSubset := range endpoints.Subsets { // This loop represents groups of endpoints within a service

						for _, address := range endpointSubset.Addresses { // This loop represents each individual endpoint
							shardedInstanceName := address.TargetRef.Name // Name of sharded instance e.g. prometheus-kubernetes-0
							instanceID := namespace + "_" + prometheusInstanceName + "_" + shardedInstanceName

							if _, ok := cardinalityInfoByInstance[instanceID]; !ok {
								// Add a newly found endpoint to the data structure
								cardinalityInfoByInstance[instanceID] = &cardinality.PrometheusCardinalityInstance{
									Namespace:           namespace,
									InstanceName:        prometheusInstanceName,
									ShardedInstanceName: shardedInstanceName,
									InstanceAddress:     "http://" + address.IP + ":9090",
									TrackedLabels: cardinality.TrackedLabelNames{
										SeriesCountByMetricNameLabels:     make([]string, 0, opts.StatsLimit),
										LabelValueCountByLabelNameLabels:  make([]string, 0, opts.StatsLimit),
										MemoryInBytesByLabelNameLabels:    make([]string, 0, opts.StatsLimit),
										SeriesCountByLabelValuePairLabels: make([]string, 0, opts.StatsLimit),
									},
								}
							} else {
								// If the endpoint is already known, update it's address
								cardinalityInfoByInstance[instanceID].InstanceAddress = "http://" + address.IP + ":9090"
							}

							if authValue, ok := promAPIAuthValues[instanceID]; ok { // Check for Prometheus API credentials for sharded instance
								cardinalityInfoByInstance[instanceID].AuthValue = authValue
							} else if authValue, ok := promAPIAuthValues[namespace+"_"+prometheusInstanceName]; ok { // Check for Prometheus API credentials for instance
								cardinalityInfoByInstance[instanceID].AuthValue = authValue
							} else if authValue, ok := promAPIAuthValues[namespace]; ok { // Check for Prometheus API credentials for namespace
								cardinalityInfoByInstance[instanceID].AuthValue = authValue
							}
						}
					}
				}
			}
		}

		// Iterates over all prometheus instances and runs cardinality exporter logic
		for instanceID, instance := range cardinalityInfoByInstance {

			prometheusClient := &http.Client{}

			log.Infof("Fetching current Prometheus status, from Prometheus instance: %v. Sharded instance: %v. Namespace: %v.", instance.InstanceName, instance.ShardedInstanceName, instance.Namespace)

			if instance.AuthValue != "" {
				log.Info("Including Authorization header.")
			}

			// Fetch the data from Prometheus
			err := backoff.Retry(func() error {
				return cardinalityInfoByInstance[instanceID].FetchTSDBStatus(prometheusClient, opts.StatsLimit)
			}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), numRetries))
			if err != nil {
				log.WithError(err).Warningf("Error fetching Prometheus status: %v", err)
				delete(cardinalityInfoByInstance, instanceID)
				continue
			}

			// Expose data on /metrics
			err = backoff.Retry(func() error {
				return cardinalityInfoByInstance[instanceID].ExposeTSDBStatus(&cardinality.SeriesCountByMetricNameGauge, &cardinality.LabelValueCountByLabelNameGauge, &cardinality.MemoryInBytesByLabelNameGauge, &cardinality.SeriesCountByLabelValuePairGauge)
			}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), numRetries))
			if err != nil {
				log.WithError(err).Warningf("Error exposing Prometheus metrics: %v", err)
			}
		}

		// Sleep until next metric update
		log.Debugf("Sleeping for %0.4f hours.", opts.Frequency)
		time.Sleep(sleepTime)
	}
}

func main() {
	_, err := flags.Parse(&opts)

	// Exit gracefully if help flag used
	if err != nil && flags.WroteHelp(err) {
		os.Exit(0)
	} else if err != nil {
		log.Fatalf("%+v", err)
	}

	if len(opts.PrometheusInstances) > 0 && opts.ServiceDiscovery {
		log.Fatal("Cannot parse Prometheus Instances (--proms) AND use Service Discovery (--service_discovery), these options are mutually exclusive.")
	} else if len(opts.PrometheusInstances) > 0 {
		log.Info("Obtaining metics from prometheus instances specified as arguments.")
	} else if opts.ServiceDiscovery {
		log.Info("Obtaining metrics from services found with service discovery.")
	} else {
		log.Fatal("Service Discovery has not been selected (--service_discovery) and no Prometheus Instances (--proms) have been passed, therefore there are no Prometheus Instances to connect to.")
	}

	logLevel, err := logging.ParseLevel(opts.LogLevel)
	if err != nil {
		log.Warnf("Invalid log level \"%s\", setting log level to \"info\".", opts.LogLevel)
		logLevel = logging.InfoLevel
	}
	logging.SetLevel(logLevel)

	log.Infof("Serving on port: %d", opts.Port)
	log.Infof("Serving Prometheus metrics on /metrics")
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	log.Infof("Starting Prometheus cardinality metric collection.")
	go collectMetrics()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", opts.Port), nil))
}
