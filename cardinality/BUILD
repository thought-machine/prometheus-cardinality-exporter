go_library(
    name = "cardinality",
    srcs = [
        "cardinality.go",
        "metrics.go",
    ],
    visibility = ["//..."],
    deps = [
        "//third_party/go:logrus",
        "//third_party/go:prometheus",
    ],
)

go_test(
    name = "cardinality_test",
    srcs = [
        "cardinality_test.go",
    ],
    deps = [
        ":cardinality",
        "//cardinality/mock_cardinality:mock_cardinality",
        "//third_party/go:testify",
    ],
)
