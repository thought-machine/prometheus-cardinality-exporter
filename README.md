# Prometheus Cardinality Exporter

A simple Prometheus exporter for exposing the cadinality of metrics Prometheus has scraped.

There is a similar project out there, however it is poorly written and not maintained:
https://github.com/artemsre/prometheus-cardinality-exporter

## Design

This is a simple Golang webserver which exposes a single endpoint, `/metrics`.

The Prometheus API is queried every 6 hours by default.

## Metrics

There are 4 types of metric exposed:

```
1. cardinality_exporter_label_value_count_by_label_name{label=""}
2. cardinality_exporter_memory_in_bytes_by_label_name{label=""}
3. cardinality_exporter_series_count_by_label_pair{label_pair=""}
4. cardinality_exporter_series_count_by_metric_name{metric=""}
```
