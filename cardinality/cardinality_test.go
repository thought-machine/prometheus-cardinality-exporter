package cardinality

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

	"github.com/thought-machine/prometheus-cardinality-exporter/cardinality/mock_cardinality"
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

		// Mock json response
		response := &http.Response{
			Status:     tt.responseStatus,
			StatusCode: tt.responseStatusCode,
			Body:       io.NopCloser(bytes.NewBufferString(tt.json)),
		}
		ts.MockPrometheusClient.EXPECT().Do(AuthHeaderCorrect(tt.expectedAuthHeaderValue)).Return(response, nil)
		err := tt.prometheusInstance.FetchTSDBStatus(ts.MockPrometheusClient)

		assert.Equal(ts.T(), tt.incomingTSDBStatus, tt.prometheusInstance.LatestTSDBStatus)
		assert.Equal(ts.T(), err, nil)

		// reset the LatestTSDBStatus, so that it doesn't affect later tests
		tt.prometheusInstance.LatestTSDBStatus = *new(TSDBStatus)

	}
}

// This function tests the ExposeTSDBStatus function on all test cases
func (ts *CardinalitySuite) TestExposeTSDBStatus() {

	for _, tt := range cardinalityTests {

		// Iterate over each of the mock input metrics and check the prometheus.GetMetricWith() function is called with each of them
		// Also check that old metrics are deleted

		tt.prometheusInstance.LatestTSDBStatus = tt.incomingTSDBStatus

		SeriesCountByMetricNameGauge := &PrometheusCardinalityMetric{GaugeVec: ts.MockSeriesCountByMetricNameGauge}
		SeriesCountByMetricNameGauge.expectMetricUpdates(tt.prometheusInstance.TrackedLabels.SeriesCountByMetricNameLabels, tt.incomingTSDBStatus.Data.SeriesCountByMetricName, tt.prometheusInstance.InstanceName, tt.prometheusInstance.ShardedInstanceName, tt.prometheusInstance.Namespace, "metric")

		LabelValueCountByLabelNameGauge := &PrometheusCardinalityMetric{GaugeVec: ts.MockLabelValueCountByLabelNameGauge}
		LabelValueCountByLabelNameGauge.expectMetricUpdates(tt.prometheusInstance.TrackedLabels.LabelValueCountByLabelNameLabels, tt.incomingTSDBStatus.Data.LabelValueCountByLabelName, tt.prometheusInstance.InstanceName, tt.prometheusInstance.ShardedInstanceName, tt.prometheusInstance.Namespace, "label")

		MemoryInBytesByLabelNameGauge := &PrometheusCardinalityMetric{GaugeVec: ts.MockMemoryInBytesByLabelNameGauge}
		MemoryInBytesByLabelNameGauge.expectMetricUpdates(tt.prometheusInstance.TrackedLabels.MemoryInBytesByLabelNameLabels, tt.incomingTSDBStatus.Data.MemoryInBytesByLabelName, tt.prometheusInstance.InstanceName, tt.prometheusInstance.ShardedInstanceName, tt.prometheusInstance.Namespace, "label")

		SeriesCountByLabelValuePairGauge := &PrometheusCardinalityMetric{GaugeVec: ts.MockSeriesCountByLabelValuePairGauge}
		SeriesCountByLabelValuePairGauge.expectMetricUpdates(tt.prometheusInstance.TrackedLabels.SeriesCountByLabelValuePairLabels, tt.incomingTSDBStatus.Data.SeriesCountByLabelValuePair, tt.prometheusInstance.InstanceName, tt.prometheusInstance.ShardedInstanceName, tt.prometheusInstance.Namespace, "label_pair")

		//Call the ExposeTSDBStatus function to check that the correct functions are called
		err := tt.prometheusInstance.ExposeTSDBStatus(SeriesCountByMetricNameGauge, LabelValueCountByLabelNameGauge, MemoryInBytesByLabelNameGauge, SeriesCountByLabelValuePairGauge)
		assert.Equal(ts.T(), err, nil)

		// reset the LatestTSDBStatus, so that it doesn't affect later tests
		tt.prometheusInstance.LatestTSDBStatus = *new(TSDBStatus)
	}
}

// This function was introduced to reduce clutter, it is used to set the EXPECTed calls to each GaugeVec
func (mockMetric *PrometheusCardinalityMetric) expectMetricUpdates(trackedLabels [10]string, incomingMetrics [10]labelValuePair, prometheusInstance string, shardedInstance string, namespace string, nameOfLabel string) {

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
		err = tt.prometheusInstance.FetchTSDBStatus(&http.Client{})
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
	responseStatus          string
	responseStatusCode      int
	prometheusInstance      PrometheusCardinalityInstance
	incomingTSDBStatus      TSDBStatus
	expectedMetrics         map[string]bool
	expectedAuthHeaderValue string
}{
	{
		`{"status":"success", "data":{"seriesCountByMetricName":[],"labelValueCountByLabelName":[],"memoryInBytesByLabelName":[],"seriesCountByLabelValuePair":[]}}`,
		"200 OK",
		200,
		PrometheusCardinalityInstance{
			Namespace:           "namespace",
			InstanceName:        "prometheus-test",
			ShardedInstanceName: "prometheus-shard",
		},
		TSDBStatus{
			Status: "success",
			Data: TSDBData{
				[10]labelValuePair{},
				[10]labelValuePair{},
				[10]labelValuePair{},
				[10]labelValuePair{},
			},
		},
		make(map[string]bool),
		"",
	},
	{
		`{"status":"success", "data":{"seriesCountByMetricName":[{"name":"label0","value":0}],"labelValueCountByLabelName":[{"name":"label1","value":1}],"memoryInBytesByLabelName":[{"name":"label2","value":2}],"seriesCountByLabelValuePair":[{"name":"label3=label3value","value":3}]}}`,
		"200 OK",
		200,
		PrometheusCardinalityInstance{
			Namespace:           "namespace-1",
			InstanceName:        "prometheus-test-1",
			ShardedInstanceName: "prometheus-shard-1",
			TrackedLabels: TrackedLabelNames{
				SeriesCountByMetricNameLabels:     [10]string{"YeOldeMetric", "MetricMcOutOfDate"},
				LabelValueCountByLabelNameLabels:  [10]string{},
				MemoryInBytesByLabelNameLabels:    [10]string{},
				SeriesCountByLabelValuePairLabels: [10]string{"StraightOuttaDateMetric"},
			},
		},
		TSDBStatus{
			Status: "success",
			Data: TSDBData{
				[10]labelValuePair{
					{Label: "label0", Value: 0},
				},
				[10]labelValuePair{
					{Label: "label1", Value: 1},
				},
				[10]labelValuePair{
					{Label: "label2", Value: 2},
				},
				[10]labelValuePair{
					{Label: "label3=label3value", Value: 3},
				},
			},
		},
		map[string]bool{
			`cardinality_exporter_series_count_by_metric_name_total{instance_namespace="namespace-1",metric="label0",scraped_instance="prometheus-test-1",sharded_instance="prometheus-shard-1"} 0`:                      true,
			`cardinality_exporter_label_value_count_by_label_name_total{instance_namespace="namespace-1",label="label1",scraped_instance="prometheus-test-1",sharded_instance="prometheus-shard-1"} 1`:                   true,
			`cardinality_exporter_memory_by_label_name_bytes{instance_namespace="namespace-1",label="label2",scraped_instance="prometheus-test-1",sharded_instance="prometheus-shard-1"} 2`:                              true,
			`cardinality_exporter_series_count_by_label_value_pair_total{instance_namespace="namespace-1",label_pair="label3=label3value",scraped_instance="prometheus-test-1",sharded_instance="prometheus-shard-1"} 3`: true,
		},
		"",
	},
	{
		`{"status":"success", "data":{"seriesCountByMetricName":[{"name":"label4","value":4},{"name":"label5","value":5}],"labelValueCountByLabelName":[{"name":"label6","value":6}],"memoryInBytesByLabelName":[{"name":"label7","value":7}],"seriesCountByLabelValuePair":[]}}`,
		"200 OK",
		200,
		PrometheusCardinalityInstance{
			Namespace:           "namespace-2",
			InstanceName:        "prometheus-test-2",
			ShardedInstanceName: "prometheus-shard-2",
			AuthValue:           "Basic YWRtaW46cGFzc3dvcmQ=",
			TrackedLabels: TrackedLabelNames{
				SeriesCountByMetricNameLabels:     [10]string{},
				LabelValueCountByLabelNameLabels:  [10]string{"OAM", "GreatGrandmetric"},
				MemoryInBytesByLabelNameLabels:    [10]string{"DeadMetric"},
				SeriesCountByLabelValuePairLabels: [10]string{},
			},
		},
		TSDBStatus{
			Status: "success",
			Data: TSDBData{
				[10]labelValuePair{
					{Label: "label4", Value: 4},
					{Label: "label5", Value: 5},
				},
				[10]labelValuePair{
					{Label: "label6", Value: 6},
				},
				[10]labelValuePair{
					{Label: "label7", Value: 7},
				},
				[10]labelValuePair{},
			},
		},
		map[string]bool{
			`cardinality_exporter_series_count_by_metric_name_total{instance_namespace="namespace-2",metric="label4",scraped_instance="prometheus-test-2",sharded_instance="prometheus-shard-2"} 4`:    true,
			`cardinality_exporter_series_count_by_metric_name_total{instance_namespace="namespace-2",metric="label5",scraped_instance="prometheus-test-2",sharded_instance="prometheus-shard-2"} 5`:    true,
			`cardinality_exporter_label_value_count_by_label_name_total{instance_namespace="namespace-2",label="label6",scraped_instance="prometheus-test-2",sharded_instance="prometheus-shard-2"} 6`: true,
			`cardinality_exporter_memory_by_label_name_bytes{instance_namespace="namespace-2",label="label7",scraped_instance="prometheus-test-2",sharded_instance="prometheus-shard-2"} 7`:            true,
		},
		"Basic YWRtaW46cGFzc3dvcmQ=",
	},
}
