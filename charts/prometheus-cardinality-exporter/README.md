# Prometheus Cardinality Exporter Helm Chart

A Helm chart for deploying [Prometheus Cardinality Exporter](https://github.com/thought-machine/prometheus-cardinality-exporter) - a monitoring tool that queries Prometheus instances to expose metrics about high-cardinality metrics that can cause performance degradation or OOM errors.

## Installation

```bash
helm install prometheus-cardinality-exporter ./charts/prometheus-cardinality-exporter
```

## Configuration

### Image

| Parameter | Description | Default |
|-----------|-------------|---------|
| `image.repository` | Image repository | `ghcr.io/thought-machine/prometheus-cardinality-exporter` |
| `image.tag` | Image tag | `""` (uses Chart.appVersion) |
| `image.pullPolicy` | Image pull policy | `IfNotPresent` |
| `imagePullSecrets` | Image pull secrets | `[]` |

### Service Account and RBAC

| Parameter | Description | Default |
|-----------|-------------|---------|
| `serviceAccount.create` | Create a service account | `true` |
| `serviceAccount.annotations` | Service account annotations | `{}` |
| `serviceAccount.name` | Service account name | `""` (generated) |
| `rbac.create` | Create RBAC resources | `true` |
| `rbac.clusterScope` | Use ClusterRole (true) or Role (false) | `true` |

**Note:** ClusterRole is required to discover Prometheus instances across all namespaces. Use Role if you only need to discover Prometheus in the release namespace.

### Service

| Parameter | Description | Default |
|-----------|-------------|---------|
| `service.type` | Service type | `ClusterIP` |
| `service.port` | Service port | `9090` |
| `service.annotations` | Service annotations | `{}` |

### Application Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `config.serviceDiscovery` | Enable Kubernetes service discovery | `true` |
| `config.selector` | Label selector for discovering Prometheus | `app=prometheus` |
| `config.namespaces` | Namespaces to discover (empty = all) | `[]` |
| `config.prometheusUrls` | Manual list of Prometheus URLs | `[]` |
| `config.frequency` | Query frequency in hours | `6` |
| `config.statsLimit` | Max items from TSDB stats | `10` |
| `config.logLevel` | Log level (debug/info/warn/error/fatal) | `info` |
| `config.regex` | Regex filter for service names | `""` |

### Authentication

| Parameter | Description | Default |
|-----------|-------------|---------|
| `auth.enabled` | Enable authentication | `false` |
| `auth.existingSecret` | Use existing secret | `""` |
| `auth.credentials` | Auth credentials map | `{}` |

**Authentication credentials format:**

For service discovery mode, keys can be (in order of precedence):
- `namespace_prometheus-name_shard-index` (most specific)
- `namespace_prometheus-name`
- `namespace` (least specific)

For manual mode (prometheusUrls), keys are the full URL.

Example:
```yaml
auth:
  enabled: true
  credentials:
    monitoring: "Bearer token123"
    monitoring_prometheus-main: "Basic YWRtaW46cGFzcw=="
```

### Prometheus Operator (ServiceMonitor)

| Parameter | Description | Default |
|-----------|-------------|---------|
| `serviceMonitor.enabled` | Create ServiceMonitor | `false` |
| `serviceMonitor.namespace` | ServiceMonitor namespace | `""` (release namespace) |
| `serviceMonitor.interval` | Scrape interval | `30s` |
| `serviceMonitor.scrapeTimeout` | Scrape timeout | `10s` |
| `serviceMonitor.labels` | Additional labels | `{}` |

### Resources

| Parameter | Description | Default |
|-----------|-------------|---------|
| `resources.requests.cpu` | CPU request | `100m` |
| `resources.requests.memory` | Memory request | `100Mi` |
| `resources.limits.memory` | Memory limit | `700Mi` |

### Pod Configuration

| Parameter | Description | Default |
|-----------|-------------|---------|
| `replicaCount` | Number of replicas | `1` |
| `podAnnotations` | Pod annotations | `{}` |
| `podSecurityContext.runAsUser` | Pod security context runAsUser | `1000` |
| `podSecurityContext.runAsNonRoot` | Pod security context runAsNonRoot | `true` |
| `securityContext.readOnlyRootFilesystem` | Container read-only root filesystem | `true` |
| `securityContext.allowPrivilegeEscalation` | Container allow privilege escalation | `false` |
| `nodeSelector` | Node selector | `{}` |
| `tolerations` | Tolerations | `[]` |
| `affinity` | Affinity | `{}` |

## Examples

### Basic installation with service discovery

```bash
helm install prometheus-cardinality-exporter ./charts/prometheus-cardinality-exporter \
  --set config.serviceDiscovery=true \
  --set config.selector="app=prometheus"
```

### Installation with manual Prometheus URLs

```bash
helm install prometheus-cardinality-exporter ./charts/prometheus-cardinality-exporter \
  --set config.serviceDiscovery=false \
  --set config.prometheusUrls[0]="http://prometheus:9090"
```

### Installation with namespace-scoped RBAC

```bash
helm install prometheus-cardinality-exporter ./charts/prometheus-cardinality-exporter \
  --set rbac.clusterScope=false \
  --set config.namespaces[0]="monitoring"
```

### Installation with ServiceMonitor (Prometheus Operator)

```bash
helm install prometheus-cardinality-exporter ./charts/prometheus-cardinality-exporter \
  --set serviceMonitor.enabled=true \
  --set serviceMonitor.interval=60s
```

### Installation with authentication

```bash
helm install prometheus-cardinality-exporter ./charts/prometheus-cardinality-exporter \
  --set auth.enabled=true \
  --set auth.credentials.monitoring="Bearer mytoken"
```

## Exposed Metrics

The exporter exposes four metric types:

| Metric | Description |
|--------|-------------|
| `cardinality_exporter_series_count_by_metric_name_total` | Series count per metric name |
| `cardinality_exporter_label_value_count_by_label_name_total` | Unique label value counts |
| `cardinality_exporter_memory_by_label_name_bytes` | Memory consumption per label |
| `cardinality_exporter_series_count_by_label_value_pair_total` | Series count by label/value pairs |

All metrics include labels: `metric/label`, `scraped_instance`, `sharded_instance`, and `instance_namespace`.
