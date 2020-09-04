# Prometheus Cardinality Exporter

A simple Prometheus exporter for exposing the cadinality of metrics Prometheus has scraped.

This was originally started as an intern project by [Harry Fallows](https://github.com/harryfallows) during his internship at [Thought Machine](https://thoughtmachine.net/).

## Design

This is a Prometheus exporter for exposing metrics related to the cardinality of metrics and labels scraped by other Prometheus instances.
The metrics are obtained by calling the Promtheus API at the ```/api/v1/status/tsdb``` endpoint ([docs](https://prometheus.io/docs/prometheus/latest/querying/api/)).

## Metrics

There are 4 types of metric exposed:

```
1. cardinality_exporter_label_value_count_by_label_name{label=""}: These metrics report label names and their respective value counts.
2. cardinality_exporter_memory_in_bytes_by_label_name{label=""}: These metrics report label names and their respective memory used in bytes. Memory usage is calculated by adding the length of all values for a given label name.
3. cardinality_exporter_series_count_by_label_pair{label_pair=""}: This will provide a list of label value pairs and their series count.
4. cardinality_exporter_series_count_by_metric_name{metric=""}: These metrics report metric names and their respective series counts.
```
## Options

(```plz run //:prometheus-cardinality-exporter -- [OPTIONS]```)

| Short Flag | Long Flag           | Description                                                                          |
|------------|---------------------|--------------------------------------------------------------------------------------|
| -s        | --selector=         | Selector for service discovery                                                       |
| -n        | --namespaces=       | Namespaces to find services in                                                       |
| -i        | --proms=            | Prometheus instance links to export metrics from                                     |
| -d        | --service_discovery | Use kubernetes service discovery instead of manually specifying Prometheus instances |
| -p        | --port=             | Port on which to serve metrics                                                       |
| -f        | --freq=             | Frequency in hours with which to query the Prometheus API                            |
| -r        | --regex=            | If any found services don’t match the regex, they are ignored                        |

## Exposing Metrics

### Installing on a cluster
See k8s/README.md for running on kubernetes

Docker images are available at thoughtmachine/prometheus-cardinality-exporter:$COMMIT
See  https://hub.docker.com/r/thoughtmachine/prometheus-cardinality-exporter


### Running Locally
```plz run //:prometheus-cardinality-exporter -- --port=<port-to-serve-on> --proms=<prometheus-instance-to-expose> [--proms=<prometheus-instance-to-expose>...] --freq=<frequency-to-ping-api>```

### Running Within a Kubernetes Cluster (with service discovery)
#### In order to deploy to a kubernetes cluster, run:
```plz run //k8s:k8s_push```
#### Make sure you alter the k8s/deployment.yaml such that it contains the options that you require:
```args: ["-c", "/home/app/prometheus-cardinality-exporter  --port=<port-to-serve-on> --service_discovery --freq=<frequency-to-ping-api> --selector=<service-selector> --regex=<regex-for-prometheus-instances> --namespaces=<namespace-of-prometheus-instances> [--namespaces=<namespace-of-prometheus-instances>...]]```

## Building
```plz build //...```

If you'd prefer to use docker to build and run all tests use

```docker build -f Dockerfile-builder . --rm=false```

## Testing
#### If you want to test any changes to cardinality/cardinality.go, you can run:
```plz run //cardinality:cardinality_test```
