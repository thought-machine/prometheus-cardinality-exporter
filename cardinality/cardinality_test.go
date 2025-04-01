package cardinality

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github/thought-machine/prometheus-cardinality-exporter/cardinality/mock_cardinality"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// CardinalitySuite used to mock objects
type CardinalitySuite struct {
	suite.Suite
	MockController                       *gomock.Controller
	MockPrometheusClient                 *mock_cardinality.MockPrometheusClient
	MockSeriesCountByMetricNameGauge     *mock_cardinality.MockPrometheusGaugeVec
	MockLabelValueCountByLabelNameGauge  *mock_cardinality.MockPrometheusGaugeVec
	MockMemoryInBytesByLabelNameGauge    *mock_cardinality.MockPrometheusGaugeVec
	MockSeriesCountByLabelValuePairGauge *mock_cardinality.MockPrometheusGaugeVec
}

func TestCardinalitySuite(t *testing.T) {
	suite.Run(t, new(CardinalitySuite))
}

// Set up each of the required mocks
func (ts *CardinalitySuite) SetupTest() {
	ts.MockController = gomock.NewController(ts.T())
	ts.MockPrometheusClient = mock_cardinality.NewMockPrometheusClient(ts.MockController)
	ts.MockSeriesCountByMetricNameGauge = mock_cardinality.NewMockPrometheusGaugeVec(ts.MockController)
	ts.MockLabelValueCountByLabelNameGauge = mock_cardinality.NewMockPrometheusGaugeVec(ts.MockController)
	ts.MockMemoryInBytesByLabelNameGauge = mock_cardinality.NewMockPrometheusGaugeVec(ts.MockController)
	ts.MockSeriesCountByLabelValuePairGauge = mock_cardinality.NewMockPrometheusGaugeVec(ts.MockController)
}

func (ts *CardinalitySuite) TearDownTest() {
	defer ts.MockController.Finish()
}

type authHeaderMatcher struct {
	expectedAuthHeaderValue string
}

func (m authHeaderMatcher) Matches(y interface{}) bool {

	authHeaders := y.(*http.Request).Header["Authorization"]
	if m.expectedAuthHeaderValue != "" {
		if len(authHeaders) > 1 || authHeaders[0] != m.expectedAuthHeaderValue {
			return false
		}
	}
	return true
}

func (m authHeaderMatcher) String() string {
	if m.expectedAuthHeaderValue == "" {
		return "contains no authorization header value"
	}
	return fmt.Sprintf("contains authorization header value %s", m.expectedAuthHeaderValue)
}

func AuthHeaderCorrect(expectedAuthHeaderValue string) gomock.Matcher {
	return authHeaderMatcher{expectedAuthHeaderValue}
}

// This tests the FetchTSDBStatus function on all of the test cases
func (ts *CardinalitySuite) TestFetchTSDBStatus() {
	for _, tt := range cardinalityTests {
		// Create a mock API server
		mux := http.NewServeMux()
		mockAPI := func(w http.ResponseWriter, r *http.Request) {
			if tt.expectedAuthHeaderValue != "" {
				if len(r.Header["Authorization"]) > 1 || r.Header["Authorization"][0] != tt.expectedAuthHeaderValue {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("401 Unauthorized"))
					return
				}
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(tt.json))
		}
		mux.HandleFunc("/api/v1/status/tsdb", mockAPI)
		mockAPIServer := &http.Server{Handler: mux}
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			log.Error(err)
		}
		mockAPIPort := listener.Addr().(*net.TCPAddr).Port
		log.Infof("Serving test API for test on port: %d.", mockAPIPort)
		go mockAPIServer.Serve(listener)
		time.Sleep(1 * time.Millisecond) // This is here to give the Serve time to set up

		// Fetch metrics from test API
		tt.prometheusInstance.InstanceAddress = fmt.Sprintf("http://localhost:%v", mockAPIPort)
		err = tt.prometheusInstance.FetchTSDBStatus(&http.Client{}, 20)
		assert.Equal(ts.T(), err, nil)

		// Verify the data was correctly unmarshaled
		assert.Equal(ts.T(), tt.prometheusInstance.LatestTSDBStatus.Status, "success")
		assert.Equal(ts.T(), len(tt.prometheusInstance.LatestTSDBStatus.Data.SeriesCountByMetricName), 2)
		assert.Equal(ts.T(), len(tt.prometheusInstance.LatestTSDBStatus.Data.LabelValueCountByLabelName), 2)
		assert.Equal(ts.T(), len(tt.prometheusInstance.LatestTSDBStatus.Data.MemoryInBytesByLabelName), 2)
		assert.Equal(ts.T(), len(tt.prometheusInstance.LatestTSDBStatus.Data.SeriesCountByLabelValuePair), 2)

		// Reset the LatestTSDBStatus
		tt.prometheusInstance.LatestTSDBStatus = TSDBStatus{}
	}
}

// This function tests the ExposeTSDBStatus function on all test cases
func (ts *CardinalitySuite) TestExposeTSDBStatus() {
	for _, tt := range cardinalityTests {
		// Create a mock API server
		mux := http.NewServeMux()
		mockAPI := func(w http.ResponseWriter, r *http.Request) {
			if tt.expectedAuthHeaderValue != "" {
				if len(r.Header["Authorization"]) > 1 || r.Header["Authorization"][0] != tt.expectedAuthHeaderValue {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("401 Unauthorized"))
					return
				}
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(tt.json))
		}
		mux.HandleFunc("/api/v1/status/tsdb", mockAPI)
		mockAPIServer := &http.Server{Handler: mux}
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			log.Error(err)
		}
		mockAPIPort := listener.Addr().(*net.TCPAddr).Port
		log.Infof("Serving test API for test on port: %d.", mockAPIPort)
		go mockAPIServer.Serve(listener)
		time.Sleep(1 * time.Millisecond) // This is here to give the Serve time to set up

		// Fetch metrics from test API
		tt.prometheusInstance.InstanceAddress = fmt.Sprintf("http://localhost:%v", mockAPIPort)
		err = tt.prometheusInstance.FetchTSDBStatus(&http.Client{}, 20)
		assert.Equal(ts.T(), err, nil)

		// Expose metrics
		err = tt.prometheusInstance.ExposeTSDBStatus(&SeriesCountByMetricNameGauge, &LabelValueCountByLabelNameGauge, &MemoryInBytesByLabelNameGauge, &SeriesCountByLabelValuePairGauge)
		assert.Equal(ts.T(), err, nil)

		// Reset the LatestTSDBStatus
		tt.prometheusInstance.LatestTSDBStatus = TSDBStatus{}
	}
}

// This function was introduced to reduce clutter, it is used to set the EXPECTed calls to each GaugeVec
func (mockMetric *PrometheusCardinalityMetric) expectMetricUpdates(trackedLabels []string, incomingMetrics []labelValuePair, prometheusInstance string, shardedInstance string, namespace string, nameOfLabel string) {

	// Iterate over each metric and apply checks to see whether GetMetricWith is called
	for _, metric := range incomingMetrics {
		gauge := prometheus.NewGauge(prometheus.GaugeOpts{})
		if metric.Label != "" {
			(mockMetric.GaugeVec).(*mock_cardinality.MockPrometheusGaugeVec).EXPECT().GetMetricWith(prometheus.Labels{nameOfLabel: metric.Label, "scraped_instance": prometheusInstance, "sharded_instance": shardedInstance, "instance_namespace": namespace}).Return(gauge, nil)
		}
	}

	// Iterate over each of the trackedLabels to check if they are no longer tracked, if so expect a call to Delete
	for _, oldMetric := range trackedLabels {
		found := false
		for _, newMetric := range incomingMetrics {
			if oldMetric == newMetric.Label {
				found = true
				break
			}
		}
		if !found && oldMetric != "" {
			(mockMetric.GaugeVec).(*mock_cardinality.MockPrometheusGaugeVec).EXPECT().Delete(prometheus.Labels{nameOfLabel: oldMetric, "scraped_instance": prometheusInstance, "sharded_instance": shardedInstance, "instance_namespace": namespace}).Return(true)
		}
	}
}

// E2E test
// 1. Creates a /metrics endpoint
// 2. Creates another endpoints to act as the Prometheus API
// 3. Calls the FetchTSDBStatus function to call the mock API
// 4. Calls the ExposeTSDBStatus function to expose the fetched metrics on the /metrics endpoint
// 5. Scrapes the /metrics endpoint and checks that the result is as expected
func (ts *CardinalitySuite) TestE2E() {

	for _, tt := range cardinalityTests {

		// Create /metrics endpoint on next available port
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		metricsServer := &http.Server{Handler: mux}
		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			panic(err)
		}
		metricsServerPort := listener.Addr().(*net.TCPAddr).Port
		log.Infof("Serving /metrics endpoint for E2E test on port: %d.", metricsServerPort)
		go metricsServer.Serve(listener)

		// Set up test API on next available port
		JSONResponse := json.RawMessage(tt.json)
		mockAPI := func(w http.ResponseWriter, r *http.Request) {
			if tt.expectedAuthHeaderValue != "" {
				if len(r.Header["Authorization"]) > 1 || r.Header["Authorization"][0] != tt.expectedAuthHeaderValue {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("401 Unauthorized"))
					return
				}
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(JSONResponse)
		}
		mux.HandleFunc("/api/v1/status/tsdb", mockAPI)
		mockAPIServer := &http.Server{Handler: mux}
		listener, err = net.Listen("tcp", ":0")
		if err != nil {
			log.Error(err)
		}
		mockAPIPort := listener.Addr().(*net.TCPAddr).Port
		log.Infof("Serving test API for E2E test on port: %d.", mockAPIPort)
		go mockAPIServer.Serve(listener)
		time.Sleep(1 * time.Millisecond) // This is here to give the Serve time to set up

		// Fetch metrics from test API
		tt.prometheusInstance.LatestTSDBStatus = *new(TSDBStatus)
		tt.prometheusInstance.InstanceAddress = fmt.Sprintf("http://localhost:%v", mockAPIPort)
		err = tt.prometheusInstance.FetchTSDBStatus(&http.Client{}, 20)
		if err != nil {
			log.WithError(err).Warningf("Error fetching Prometheus status: %v", err)
		}

		// Expose test metrics on /metrics
		tt.prometheusInstance.ExposeTSDBStatus(&SeriesCountByMetricNameGauge, &LabelValueCountByLabelNameGauge, &MemoryInBytesByLabelNameGauge, &SeriesCountByLabelValuePairGauge)

		// Perform GET request on /metrics
		apiURL := fmt.Sprintf("http://localhost:%v/metrics", metricsServerPort)
		request, err := http.Get(apiURL)
		if err != nil {
			panic(fmt.Sprintf("Can't connect to %v: %v ", apiURL, err))
		}
		defer request.Body.Close()

		// Read the body of the response from /metrics GET request
		body, err := io.ReadAll(request.Body)
		if err != nil {
			panic(fmt.Sprintf("Can't read from socket: %v", err))
		}

		// Check that all expected metrics are found
		bodyString := string(body)
		scanner := bufio.NewScanner(strings.NewReader(bodyString))
		for scanner.Scan() {
			if tt.expectedMetrics[scanner.Text()] {
				delete(tt.expectedMetrics, scanner.Text())
			}
		}

		// Assert that there are no expected metrics unaccounted for
		assert.Equal(ts.T(), 0, len(tt.expectedMetrics))
		metricsServer.Shutdown(context.Background()) // Shutdown the server at the end of the function
		mockAPIServer.Shutdown(context.Background()) // Shutdown the server at the end of the function
	}
}

// Test cases
var cardinalityTests = []struct {
	json                    string
	prometheusInstance      PrometheusCardinalityInstance
	expectedAuthHeaderValue string
	expectedMetrics         map[string]bool
}{
	{
		json: `{
			"status": "success",
			"data": {
				"seriesCountByMetricName": [
					{"name": "metric1", "value": 1},
					{"name": "metric2", "value": 2}
				],
				"labelValueCountByLabelName": [
					{"name": "label1", "value": 1},
					{"name": "label2", "value": 2}
				],
				"memoryInBytesByLabelName": [
					{"name": "label1", "value": 1},
					{"name": "label2", "value": 2}
				],
				"seriesCountByLabelValuePair": [
					{"name": "label1=value1", "value": 1},
					{"name": "label2=value2", "value": 2}
				]
			}
		}`,
		prometheusInstance: PrometheusCardinalityInstance{
			Namespace:           "test",
			InstanceName:        "test-instance",
			ShardedInstanceName: "test-sharded",
			TrackedLabels: TrackedLabelNames{
				SeriesCountByMetricNameLabels:     make([]string, 0),
				LabelValueCountByLabelNameLabels:  make([]string, 0),
				MemoryInBytesByLabelNameLabels:    make([]string, 0),
				SeriesCountByLabelValuePairLabels: make([]string, 0),
			},
		},
		expectedAuthHeaderValue: "",
		expectedMetrics: map[string]bool{
			`cardinality_exporter_series_count_by_metric_name_total{instance_namespace="test",metric="metric1",scraped_instance="test-instance",sharded_instance="test-sharded"} 1`:                true,
			`cardinality_exporter_series_count_by_metric_name_total{instance_namespace="test",metric="metric2",scraped_instance="test-instance",sharded_instance="test-sharded"} 2`:                true,
			`cardinality_exporter_label_value_count_by_label_name_total{instance_namespace="test",label="label1",scraped_instance="test-instance",sharded_instance="test-sharded"} 1`:              true,
			`cardinality_exporter_label_value_count_by_label_name_total{instance_namespace="test",label="label2",scraped_instance="test-instance",sharded_instance="test-sharded"} 2`:              true,
			`cardinality_exporter_memory_by_label_name_bytes{instance_namespace="test",label="label1",scraped_instance="test-instance",sharded_instance="test-sharded"} 1`:                         true,
			`cardinality_exporter_memory_by_label_name_bytes{instance_namespace="test",label="label2",scraped_instance="test-instance",sharded_instance="test-sharded"} 2`:                         true,
			`cardinality_exporter_series_count_by_label_value_pair_total{instance_namespace="test",label_pair="label1=value1",scraped_instance="test-instance",sharded_instance="test-sharded"} 1`: true,
			`cardinality_exporter_series_count_by_label_value_pair_total{instance_namespace="test",label_pair="label2=value2",scraped_instance="test-instance",sharded_instance="test-sharded"} 2`: true,
		},
	},
}
