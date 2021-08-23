package cardinality

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	logging "github.com/sirupsen/logrus"
)

var log = logging.WithFields(logging.Fields{})

// PrometheusClient interface for mock
type PrometheusClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// PrometheusGaugeVec interface for mock
type PrometheusGaugeVec interface {
	GetMetricWith(labels prometheus.Labels) (prometheus.Gauge, error)
	Delete(labels prometheus.Labels) bool
	Collect(ch chan<- prometheus.Metric)
	Describe(ch chan<- *prometheus.Desc)
}

// PrometheusCardinalityMetric used to apply methods to a PrometheusGaugeVec (updateMetric)
type PrometheusCardinalityMetric struct {
	GaugeVec PrometheusGaugeVec
}

// Struct for retaining a single label value pair
type labelValuePair struct {
	Label string `json:"name"`
	Value uint64 `json:"value"`
}

// TSDBData contains the metric updates
type TSDBData struct {
	SeriesCountByMetricName     [10]labelValuePair `json:"seriesCountByMetricName"`
	LabelValueCountByLabelName  [10]labelValuePair `json:"labelValueCountByLabelName"`
	MemoryInBytesByLabelName    [10]labelValuePair `json:"memoryInBytesByLabelName"`
	SeriesCountByLabelValuePair [10]labelValuePair `json:"seriesCountByLabelValuePair"`
}

// TSDBStatus : a struct to hold data returned by the Prometheus API call
type TSDBStatus struct {
	Status string   `json:"status"`
	Data   TSDBData `json:"data"`
}

// TrackedLabelNames : a struct to keep track of which metrics we are currently tracking
type TrackedLabelNames struct {
	SeriesCountByMetricNameLabels     [10]string
	LabelValueCountByLabelNameLabels  [10]string
	MemoryInBytesByLabelNameLabels    [10]string
	SeriesCountByLabelValuePairLabels [10]string
}

// PrometheusCardinalityInstance stores all that is required to know about  prometheus instance
// inc. it's name, it's address, the latest api call results, and the labels currently being tracked
type PrometheusCardinalityInstance struct {
	Namespace           string
	InstanceName        string
	InstanceAddress     string
	ShardedInstanceName string
	AuthValue           string
	LatestTSDBStatus    TSDBStatus
	TrackedLabels       TrackedLabelNames
}

// FetchTSDBStatus saves tracked TSDB status metrics in the struct pointed to by the "data" parameter
func (promInstance *PrometheusCardinalityInstance) FetchTSDBStatus(prometheusClient PrometheusClient) error {

	// Create a GET request to the Prometheus API
	apiURL := promInstance.InstanceAddress + "/api/v1/status/tsdb"
	request, err := http.NewRequest("GET", apiURL, nil)

	if promInstance.AuthValue != "" {
		request.Header.Add("Authorization", promInstance.AuthValue)
	}

	if err != nil {
		return fmt.Errorf("Cannot create GET request to %v: %v", apiURL, err)
	}

	// Perform GET request
	res, err := prometheusClient.Do(request)
	if err != nil {
		return fmt.Errorf("Can't connect to %v: %v ", apiURL, err)
	}
	defer res.Body.Close()

	// Check the response and either log it, if 2xx, or return an error
	responseStatusLog := fmt.Sprintf("Request to %s returned status %s.", apiURL, res.Status)
	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		return errors.New(responseStatusLog)
	}
	log.Debug(responseStatusLog)

	// Read the body of the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Can't read from socket: %v", err)
	}

	// Parse the JSON response body into a struct
	err = json.Unmarshal(body, &promInstance.LatestTSDBStatus)
	if err != nil {
		return fmt.Errorf("Can't parse json: %v", err)
	}
	return nil
}

// ExposeTSDBStatus expose TSDB status to /metrics
func (promInstance *PrometheusCardinalityInstance) ExposeTSDBStatus(seriesCountByMetricNameGauge, labelValueCountByLabelNameGauge, memoryInBytesByLabelNameGauge, seriesCountByLabelValuePairGauge *PrometheusCardinalityMetric) (err error) {

	promInstance.TrackedLabels.SeriesCountByMetricNameLabels, err = seriesCountByMetricNameGauge.updateMetric(promInstance.LatestTSDBStatus.Data.SeriesCountByMetricName, promInstance.TrackedLabels.SeriesCountByMetricNameLabels, promInstance.InstanceName, promInstance.ShardedInstanceName, promInstance.Namespace, "metric")
	if err != nil {
		return err
	}
	promInstance.TrackedLabels.LabelValueCountByLabelNameLabels, err = labelValueCountByLabelNameGauge.updateMetric(promInstance.LatestTSDBStatus.Data.LabelValueCountByLabelName, promInstance.TrackedLabels.LabelValueCountByLabelNameLabels, promInstance.InstanceName, promInstance.ShardedInstanceName, promInstance.Namespace, "label")
	if err != nil {
		return err
	}
	promInstance.TrackedLabels.MemoryInBytesByLabelNameLabels, err = memoryInBytesByLabelNameGauge.updateMetric(promInstance.LatestTSDBStatus.Data.MemoryInBytesByLabelName, promInstance.TrackedLabels.MemoryInBytesByLabelNameLabels, promInstance.InstanceName, promInstance.ShardedInstanceName, promInstance.Namespace, "label")
	if err != nil {
		return err
	}
	promInstance.TrackedLabels.SeriesCountByLabelValuePairLabels, err = seriesCountByLabelValuePairGauge.updateMetric(promInstance.LatestTSDBStatus.Data.SeriesCountByLabelValuePair, promInstance.TrackedLabels.SeriesCountByLabelValuePairLabels, promInstance.InstanceName, promInstance.ShardedInstanceName, promInstance.Namespace, "label_pair")
	if err != nil {
		return err
	}

	return nil

}

// Updates the given metric with new values and deletes ones which are no longer being reported
func (Metric *PrometheusCardinalityMetric) updateMetric(newLabelsValues [10]labelValuePair, trackedLabels [10]string, prometheusInstance string, shardedInstance string, namespace string, nameOfLabel string) (newTrackedLabels [10]string, err error) {

	for idx, labelValuePair := range newLabelsValues {
		if labelValuePair.Label == "" {
			break
		}
		metricGauge, err := Metric.GaugeVec.GetMetricWith(prometheus.Labels{nameOfLabel: labelValuePair.Label, "scraped_instance": prometheusInstance, "sharded_instance": shardedInstance, "instance_namespace": namespace})
		if err != nil {
			return trackedLabels, fmt.Errorf("Error updating metric with label name %v: %v", labelValuePair.Label, err)
		}
		metricGauge.Set(float64(labelValuePair.Value))
		newTrackedLabels[idx] = labelValuePair.Label
	}

	for _, oldLabel := range trackedLabels {
		found := false
		for _, newLabelVP := range newLabelsValues {
			if oldLabel == newLabelVP.Label {
				found = true
				break
			}
		}
		if !found && oldLabel != "" {
			Metric.GaugeVec.Delete(prometheus.Labels{nameOfLabel: oldLabel, "scraped_instance": prometheusInstance, "sharded_instance": shardedInstance, "instance_namespace": namespace})
		}
	}

	return newTrackedLabels, nil

}
