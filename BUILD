subinclude("///third_party/subrepos/pleasings//docker")

subinclude("///third_party/subrepos/pleasings//k8s")

go_binary(
    name = "prometheus-cardinality-exporter",
    srcs = ["main.go"],
    static=False,
    deps = [
        "//cardinality",
        "//third_party/go:prometheus",
        "//third_party/go/kubernetes:apimachinery",
        "//third_party/go/kubernetes:client-go",
        "//third_party/go:logrus",
        "//third_party/go:backoff",
        "//third_party/go:go-flags",
    ],
)

docker_image(
    name = "prometheus-cardinality-exporter_alpine",
    srcs = [
        ":prometheus-cardinality-exporter",
    ],
    dockerfile = "Dockerfile-prometheus-cardinality-exporter",
    image = "prometheus-cardinality-exporter",

)

k8s_config(
    name = "k8s",
    srcs = [
        "k8s/deployment.yaml",
        "k8s/service.yaml",
        "k8s/service-account.yaml",
        "k8s/clusterrole.yaml",
        "k8s/clusterrolebinding.yamls",
    ],
    containers = [
        ":prometheus-cardinality-exporter_alpine",
    ],
)
