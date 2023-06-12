# Prometheus Cardinality Exporter

A simple Prometheus exporter for exposing the cardinality of metrics Prometheus has scraped.

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

(```go run . [OPTIONS]```)

| Short Flag | Long Flag           | Description                                                                          |
|------------|---------------------|--------------------------------------------------------------------------------------|
| -s        | --selector         | Selector for service discovery                                                       |
| -n        | --namespaces       | Namespaces to find services in                                                       |
| -i        | --proms            | Prometheus instance links to export metrics from                                     |
| -d        | --service_discovery | Use kubernetes service discovery instead of manually specifying Prometheus instances |
| -p        | --port             | Port on which to serve metrics                                                       |
| -f        | --freq             | Frequency in hours with which to query the Prometheus API                            |
| -r        | --regex            | If any found services don’t match the regex, they are ignored                        |
| -a        | --auth             | Location of YAML file where Prometheus instance authorisation credentials can be found. For instances that don't appear in the file, it is assumed that no authorisation is required to access them. |
| -l        | --log.level        | Level for logging. Options (in order of verbosity): [debug, info, warn, error, fatal].|

## Exposing Metrics

### Dealing with auth'd Prometheus instances
Some Prometheus instances will not let the exporter access the ```/api/v1/status/tsdb``` endpoint without providing some authorisation credentials. To access these instances, you must provide the authorisation credentials required. The solution to this depends on whether you are using the ```--proms``` or ```--service_discovery``` flag:
- With ```--proms```:
    - Use the ```--auth``` flag to specify a YAML file mapping ```--proms``` instances to the values required.
    - Example: \<my-prometheus\>:\<my-Authorization-header-value\>).
- With ```--service_discovery```:
    - Use the ```--auth``` flag to specify a YAML file mapping instance identifiers to the values required.
    - Identifiers can be at the namespace level, the Prometheus instance level, or the sharded instance level.
    - The naming convention is: ```<namespace>[_<prometheus-instance-name>[_<sharded-instance-name>]]``` (square brackets means optional).
    - Examples (k8s/secret.yaml provides an example Kubernetes Secret):
        - ```my-namespace: "Bearer 123456789"``` - specifies that requests to Prometheus instances in namespace "my-namespace" should include the header "Authorization: Bearer 123456789".
        - ```my-namespace_my-prometheus-instance: "Basic 123456789"``` - specifies that requests to the Prometheus instance "my-prometheus-instance" in namespace "my-namespace" should include the header "Authorization: Basic 123456789".
        - ```my-namespace_my-prometheus-instance_my-sharded-instance: "Basic 987654321"``` - specifies that requests to sharded instance "my-sharded-instance" with the Prometheus instance name "my-prometheus-instance" in namespace "my-namespace" should include the header "Authorization: Basic 987654321".
    - When looking for authorisation credentials, the exporter will look in this order:
        1. sharded instance level
        1. Prometheus instance level
        1. namespace level
        1. nothing

In both cases you must specify the exact value of the Authorization header, since the request to ```/api/v1/status/tsdb``` will include the header: ```Authorization: <your-provided-value>```. k8s/secret.yaml provides an example of the ```--service_discovery``` ```--auth``` file.

### Installing on a cluster
See k8s/README.md for running on kubernetes

#### Docker images

Distroless docker images are available at thoughtmachine/prometheus-cardinality-exporter:$COMMIT_distroless

Docker images based on Alpine are available at thoughtmachine/prometheus-cardinality-exporter:$COMMIT

See  https://hub.docker.com/r/thoughtmachine/prometheus-cardinality-exporter

### Running Locally
```go run . --port=<port-to-serve-on> --proms=<prometheus-instance-to-expose> [--proms=<prometheus-instance-to-expose>...] --freq=<frequency-to-ping-api>```

### Running Within a Kubernetes Cluster (with service discovery)
#### In order to deploy to a kubernetes cluster:

Tweak and apply the files in k8s/

#### Make sure you alter the k8s/deployment.yaml such that it contains the options that you require:
In the example below, all of the possible flags that can be used with the ```--service_discovery``` option are included.\
NOTE: not all flags are required, for example, you do not need the ```--auth``` flag if none of your Prometheus instances require authorization to access.

```args: ["-c", "/home/app/prometheus-cardinality-exporter  --auth=<prometheus-api-auth-values-filepath> --port=<port-to-serve-on> --service_discovery --freq=<frequency-to-ping-api> --selector=<service-selector> --regex=<regex-for-prometheus-instances> --namespaces=<namespace-of-prometheus-instances> [--namespaces=<namespace-of-prometheus-instances>...]]```

## Building
```go build ./...```

If you'd prefer to use docker to build and run all tests use

```docker build -f Dockerfile-builder . --rm=false```

## Testing
```go test ./...```

## Linting

```go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.1 run```
