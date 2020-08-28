package cardinality

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// SeriesCountByMetricNameGauge provides a list of metrics names and their series count
	SeriesCountByMetricNameGauge = PrometheusCardinalityMetric{
		GaugeVec: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem: "cardinality_exporter",
				Name:      "series_count_by_metric_name_total",
				Help:      "A list of metrics names and their series count.",
			},
			[]string{
				"metric",
				"scraped_instance",
				"sharded_instance",
				"instance_namespace",
			},
		),
	}

	// LabelValueCountByLabelNameGauge provides a list of the label names and their value count
	LabelValueCountByLabelNameGauge = PrometheusCardinalityMetric{
		GaugeVec: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem: "cardinality_exporter",
				Name:      "label_value_count_by_label_name_total",
				Help:      "A list of the label names and their value count.",
			},
			[]string{
				"label",
				"scraped_instance",
				"sharded_instance",
				"instance_namespace",
			},
		),
	}

	// MemoryInBytesByLabelNameGauge provides a list of the label names and memory used in bytes
	// Memory usage is calculated by adding the length of all values for a given label name
	MemoryInBytesByLabelNameGauge = PrometheusCardinalityMetric{
		GaugeVec: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem: "cardinality_exporter",
				Name:      "memory_by_label_name_bytes",
				Help:      "A list of the label names and memory used in bytes. Memory usage is calculated by adding the length of all values for a given label name.",
			},
			[]string{
				"label",
				"scraped_instance",
				"sharded_instance",
				"instance_namespace",
			},
		),
	}

	// SeriesCountByLabelValuePairGauge provides a list of label value pairs and their series count
	SeriesCountByLabelValuePairGauge = PrometheusCardinalityMetric{
		GaugeVec: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem: "cardinality_exporter",
				Name:      "series_count_by_label_value_pair_total",
				Help:      "A list of label/value pairs and their series count.",
			},
			[]string{
				"label_pair",
				"scraped_instance",
				"sharded_instance",
				"instance_namespace",
			},
		),
	}
)

func init() {
	prometheus.MustRegister(SeriesCountByMetricNameGauge.GaugeVec)
	prometheus.MustRegister(LabelValueCountByLabelNameGauge.GaugeVec)
	prometheus.MustRegister(MemoryInBytesByLabelNameGauge.GaugeVec)
	prometheus.MustRegister(SeriesCountByLabelValuePairGauge.GaugeVec)
}
