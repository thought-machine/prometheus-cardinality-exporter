subinclude("///third_party/subrepos/pleasings//k8s")

k8s_config(
    name = "k8s",
    srcs = [
        "clusterrole.yaml",
        "clusterrolebinding.yaml",
        "deployment.yaml",
        "service.yaml",
        "service-account.yaml",
    ],
    containers = [
        "//:prometheus-cardinality-exporter_alpine",
    ],
    visibility = [
        "//...",
    ],
)
