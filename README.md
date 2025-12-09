# Prometheus Cardinality Exporter

![Go Version](https://img.shields.io/github/go-mod/go-version/thought-machine/prometheus-cardinality-exporter)
![License](https://img.shields.io/github/license/thought-machine/prometheus-cardinality-exporter)
![Docker Pulls](https://img.shields.io/docker/pulls/thoughtmachine/prometheus-cardinality-exporter)

A simple Prometheus exporter for exposing the cardinality of metrics Prometheus has scraped. It queries the target Prometheus API at `/api/v1/status/tsdb` to provide granular insights into label usage, series counts, and memory consumption.

This tool is critical for identifying **high-cardinality metrics** that may be causing performance degradation or OOM (Out of Memory) kills in your monitoring infrastructure.

This was originally started as an intern project by [Harry Fallows](https://github.com/harryfallows)  during his internship at [Thought Machine](https://www.thoughtmachine.net/). 

## Features

* **Granular Cardinality Metrics**: Export label value counts, memory usage by label, and series counts by metric name.
* **Kubernetes Service Discovery**: Automatically discover and scrape Prometheus pods in your cluster.
* **Multi-Instance Support**: Monitor multiple Prometheus instances from a single exporter.
* **Auth Compatible**: Specific support for Basic Auth and Bearer Token setups for secured Prometheus instances.

---

## 🚀 Quick Start

### Docker

**Docker images**

Distroless docker images are available at:
`ghcr.io/thought-machine/prometheus-cardinality-exporter:$COMMIT`

> **Note:** Images are **no longer** uploaded to Docker Hub. See [Docker Hub](https://hub.docker.com/r/thoughtmachine/prometheus-cardinality-exporter).


Run the exporter locally, pointing it to a Prometheus instance running on `localhost:9090`.

```bash
docker run -p 9090:9090 thoughtmachine/prometheus-cardinality-exporter \
  --proms=[http://host.docker.internal:9090](http://host.docker.internal:9090) \
  --port=9090 \
  --freq=1
```

Access metrics at: `http://localhost:9090/metrics`


**Binary**
```bash
# Clone and run
git clone [https://github.com/thought-machine/prometheus-cardinality-exporter.git](https://github.com/thought-machine/prometheus-cardinality-exporter.git)
cd prometheus-cardinality-exporter
go run . --proms=http://localhost:9090
```

## 📊 Exposed Metrics

The exporter exposes the following metrics:

| Metric Name | Description |
| :--- | :--- |
| `cardinality_exporter_label_value_count_by_label_name` | Count of unique values for a specific label name. Useful for finding labels with too many values (e.g., `user_id` or `pod_name`). |
| `cardinality_exporter_memory_in_bytes_by_label_name` | Memory used by a specific label name (sum of the length of all values). |
| `cardinality_exporter_series_count_by_label_pair` | Number of series associated with a specific label key-value pair. |
| `cardinality_exporter_series_count_by_metric_name` | Number of series per metric name. Useful for identifying the "heaviest" metrics in your TSDB. |


## ⚙️ Configuration

The exporter is configured via command-line flags.
`(go run . [OPTIONS])`

| Short | Flag | Description | Default |
| :--- | :--- | :--- | :--- |
| `-s` | `--selector` | Label selector for K8s service discovery (e.g., `app=prometheus`). | |
| `-n` | `--namespaces` | Comma-separated K8s namespaces to discover services in. | |
| `-i` | `--proms` | Manual list of Prometheus URLs to scrape. | |
| `-d` | `--service_discovery` | Enable Kubernetes service discovery (replaces `--proms`). | `false` |
| `-p` | `--port` | Port to expose the exporter metrics on. | `9090` |
| `-f` | `--freq` | Frequency (in hours) to query the target Prometheus TSDB API. | |
| `-r` | `--regex` | Regex to filter discovered service names. | |
| `-a` | `--auth` | Path to YAML file containing auth credentials. | |
| `-L` | `--stats-limit` | Limit the number of items fetched from TSDB stats. | `10` |
| `-l` | `--log.level` | Log level (`debug`, `info`, `warn`, `error`, `fatal`). | `info` |


## 🔐 Authentication Guide

If your target Prometheus instances require authentication (e.g., Basic Auth or Bearer Token) to access the `/api/v1/status/tsdb` endpoint, you must provide a credential configuration file using the `--auth` flag.

The structure of this file depends on how you are discovering your Prometheus targets.

#### 1. Using Manual List (`--proms`)
When manually specifying Prometheus URLs, map the full URL to the Authorization header value. 
```yaml
"<prometheus-url>": "<full-authorization-header-value>"
```

#### 2. Using Service Discovery (`--service_discovery`)

When using Kubernetes service discovery, you map credentials using specific **identifiers**. The exporter checks for credentials in the following order of precedence:

1. **Sharded Instance Level** (Most specific)
2. **Prometheus Instance Level**
3. **Namespace Level** (Least specific)
4. **No Auth** (If no match found)
   
Naming Convention: `<namespace>[_<prometheus-instance-name>[_<sharded-instance-name>]]`

Example Configuration:

```yaml
# 1. Namespace Level
# Apply to ALL Prometheus instances found in "monitoring-ns"
"monitoring-ns": "Bearer eyJhbGciOiJ..."

# 2. Instance Level
# Apply specifically to the "main-prom" instance in "default" namespace
"default_main-prom": "Basic YWRtaW46cGFzc3dvcmQ="

# 3. Sharded Instance Level
# Apply to a specific shard of a Prometheus instance
"default_main-prom_shard-0": "Basic 987654321"
```

> ⚠️ **Note:** You must provide the full value for the Authorization header (e.g., including `Basic` or `Bearer` prefix).
> * Correct: "Basic YWRtaW46..." or "Bearer eyJ..."
> * Incorrect: "YWRtaW46..."

## ☸️ Kubernetes Deployment
To deploy in Kubernetes with **Service Discovery ** enabled, the exporter needs RBAC permissions to list Services and Pods.

**1. RBAC Permissions**

Create a `ServiceAccount`, `ClusterRole`, and `ClusterRoleBinding`.

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cardinality-exporter
  namespace: monitoring
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cardinality-exporter-role
rules:
- apiGroups: [""]
  resources: ["services", "pods", "endpoints"]
  verbs: ["get", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cardinality-exporter-binding
subjects:
- kind: ServiceAccount
  name: cardinality-exporter
  namespace: monitoring
roleRef:
  kind: ClusterRole
  name: cardinality-exporter-role
  apiGroup: rbac.authorization.k8s.io
```

**2. Deployment Manifest**

Deploy the exporter using the service account created above. Ensure you set the `--namespaces` and `--selector flags` to match your Prometheus installation.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: prometheus-cardinality-exporter
  namespace: monitoring
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cardinality-exporter
  template:
    metadata:
      labels:
        app: cardinality-exporter
    spec:
      serviceAccountName: cardinality-exporter
      containers:
      - name: exporter
        image: thoughtmachine/prometheus-cardinality-exporter:latest
        args:
        - "--service_discovery"
        - "--namespaces=monitoring"
        - "--selector=app.kubernetes.io/name=prometheus"
        - "--freq=1"
        ports:
        - containerPort: 9090
```

## 🔨 Building 

Build Binary
```bash
go build ./...
```

Build Docker Image
```bash
docker build -f Dockerfile-builder . -t prometheus-cardinality-exporter
```

## 🧪 Testing 
```bash
go test ./...
```

## 🚨 Linting 
```bash
go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0 run
```

