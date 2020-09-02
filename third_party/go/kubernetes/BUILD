KUBERNETES_VERSION = "1.12.6"

#########################
# Top Level K8s targets
#########################

go_get(
    name = "client-go",
    get = "k8s.io/client-go/...",
    install = [
        "discovery",
        "kubernetes",
        "kubernetes/scheme",
        "kubernetes/typed/...",
        "pkg/apis/clientauthentication/...",
        "pkg/version",
        "plugin/pkg/client/auth/exec",
        "rest",
        "rest/watch",
        "tools/clientcmd/...",
        "tools/metrics",
        "tools/reference",
        "transport",
        "util/cert",
        "util/connrotation",
        "util/flowcontrol",
        "util/integer",
    ],
    repo = "github.com/kubernetes/client-go",
    revision = "kubernetes-%s" % KUBERNETES_VERSION,
    strip = [
        "plugin/pkg/client/auth/openstack",
        "plugin/pkg/client/auth/azure",
    ],
    visibility = ["PUBLIC"],
    deps = [
        ":api",
        ":apimachinery",
        ":github-com-gregjones-httpcache",
        ":github-com-howeyc-gopass",
        ":github-com-imdario-mergo",
        ":github-com-juju-ratelimit",
        "//third_party/go:oauth2",
        "//third_party/go:x_time_rate",
    ],
)

##########################
# Internal K8s libraries
##########################

go_get(
    name = "api",
    get = "k8s.io/api/...",
    repo = "github.com/kubernetes/api",
    revision = "kubernetes-%s" % KUBERNETES_VERSION,
    visibility = ["PUBLIC"],
    deps = [
        ":apimachinery",
    ],
)

go_get(
    name = "apimachinery",
    get = "k8s.io/apimachinery/...",
    repo = "github.com/kubernetes/apimachinery",
    revision = "kubernetes-%s" % KUBERNETES_VERSION,
    strip = [
        "pkg/util/proxy",
        "pkg/util/jsonmergepatch",
        "pkg/runtime/serializer/recognizer/testing",
    ],
    visibility = ["PUBLIC"],
    deps = [
        ":github-com-docker-spdystream",
        ":github-com-ghodss-yaml",
        ":github-com-google-gofuzz",
        ":github-com-json-iterator-go",
        ":k8s-io-kube-openapi",
        "//third_party/go:glog",
        "//third_party/go:go-spew",
        "//third_party/go:gogo-protobuf",
        "//third_party/go:inf.v0",
        "//third_party/go:lru",
        "//third_party/go:pflag",
        "//third_party/go:uuid",
    ],
)

go_get(
    name = "k8s-io-kube-openapi",
    get = "k8s.io/kube-openapi",
    install = [
        "pkg/common",
        "pkg/util/proto",
    ],
    repo = "github.com/kubernetes/kube-openapi",
    revision = "addea2498afe5a6d58f8bdcd9ae51363d12f12ef",
    visibility = ["PUBLIC"],
    deps = [
        ":github-com-emicklei-go-restful",
        ":github-com-go-openapi-spec",
        ":github-com-googleapis-gnostic",
        "//third_party/go:yaml.v2",
    ],
)

# ################
# Dependencies
# ################

go_get(
    name = "github-com-docker-spdystream",
    get = "github.com/docker/spdystream/...",
    revision = "6480d4af844c189cf5dd913db24ddd339d3a4f85",
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go:websocket",
    ],
)

go_get(
    name = "github-com-emicklei-go-restful",
    get = "github.com/emicklei/go-restful/...",
    revision = "4d17d1531f3a98ecae5584d389e98fcb4904a324",
    strip = [
        "examples",
    ],
    visibility = ["PUBLIC"],
)

go_get(
    name = "github-com-garyburd-redigo",
    get = "github.com/garyburd/redigo",
    install = [
        "internal",
        "redis",
    ],
    revision = "v1.6.2",
)

go_get(
    name = "github-com-ghodss-yaml",
    get = "github.com/ghodss/yaml",
    revision = "v1.0.0",
    deps = [
        "//third_party/go:yaml.v2",
    ],
)

go_get(
    name = "github-com-go-openapi-jsonpointer",
    get = "github.com/go-openapi/jsonpointer",
    revision = "v0.19.3",
    deps = [
        ":github-com-go-openapi-swag",
    ],
)

go_get(
    name = "github-com-go-openapi-jsonreference",
    get = "github.com/go-openapi/jsonreference",
    revision = "v0.19.4",
    deps = [
        ":github-com-go-openapi-jsonpointer",
        ":github-com-puerkitobio-purell",
    ],
)

go_get(
    name = "github-com-go-openapi-spec",
    get = "github.com/go-openapi/spec",
    revision = "v0.19.9",
    visibility = ["PUBLIC"],
    deps = [
        ":github-com-go-openapi-jsonpointer",
        ":github-com-go-openapi-jsonreference",
    ],
)

go_get(
    name = "github-com-go-openapi-swag",
    get = "github.com/go-openapi/swag",
    revision = "v0.19.9",
    deps = [
        ":github-com-mailru-easyjson",
        "//third_party/go:yaml.v2",
    ],
)

go_get(
    name = "github-com-google-btree",
    get = "github.com/google/btree",
    revision = "479b5e81b0a93ec038d201b0b33d17db599531d3",
)

go_get(
    name = "github-com-google-gofuzz",
    get = "github.com/google/gofuzz",
    revision = "v1.1.0",
    visibility = ["PUBLIC"],
)

go_get(
    name = "github-com-googleapis-gnostic",
    get = "github.com/googleapis/gnostic",
    install = [
        "compiler",
        "extensions",
        "OpenAPIv2",
    ],
    revision = "v0.4.0",
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go:pretty",
        "//third_party/go:protobuf",
        "//third_party/go:yaml.v2",

    ],
)

go_get(
    name = "github-com-gregjones-httpcache",
    get = "github.com/gregjones/httpcache/...",
    revision = "901d90724c7919163f472a9812253fb26761123d",
    visibility = ["PUBLIC"],
    deps = [
        ":github-com-garyburd-redigo",
        ":github-com-peterbourgon-diskv",
        ":github-com-syndtr-goleveldb",
        ":github-com-gomodule-redigo-redis",
        "//third_party/go:gomemcache",
    ],
)

go_get(
    name = "github-com-gomodule-redigo-redis",
    get = "github.com/gomodule/redigo/redis",
    revision = "v1.8.2",
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go:testify",
    ],
)

go_get(
    name = "github-com-json-iterator-go",
    get = "github.com/json-iterator/go",
    revision = "v1.1.10",
    visibility = ["PUBLIC"],
    deps = [
        ":github-com-modern-go-reflect2",
    ],
)

go_get(
    name = "github-com-juju-ratelimit",
    get = "github.com/juju/ratelimit",
    revision = "1.0.1",
)

go_get(
    name = "github-com-mailru-easyjson",
    get = "github.com/mailru/easyjson",
    install = [
        "buffer",
        "jlexer",
        "jwriter",
    ],
    deps = [
        ":github-com-josharian-intern",
    ],
    revision = "v0.7.6",
)

go_get(
    name = "github-com-josharian-intern",
    get = "github.com/josharian/intern",
    revision = "v1.0.0",

)
go_get(
    name = "github-com-modern-go-concurrent",
    get = "github.com/modern-go/concurrent",
    revision = "1.0.3",
)

go_get(
    name = "github-com-modern-go-reflect2",
    get = "github.com/modern-go/reflect2",
    revision = "1.0.0",
    deps = [
        ":github-com-modern-go-concurrent",
    ],
)

go_get(
    name = "github-com-onsi-ginkgo",
    get = "github.com/onsi/ginkgo",
    revision = "v1.4.0",
)

go_get(
    name = "github-com-onsi-gomega",
    get = "github.com/onsi/gomega",
    revision = "v1.3.0",
    deps = [
        "//third_party/go:protobuf",
        "//third_party/go:x_net",
        "//third_party/go:yaml.v2",
    ],
)

go_get(
    name = "github-com-peterbourgon-diskv",
    get = "github.com/peterbourgon/diskv",
    revision = "v2.0.1",
    visibility = ["//third_party/go/..."],
    deps = [
        ":github-com-google-btree",
    ],
)

go_get(
    name = "github-com-puerkitobio-purell",
    get = "github.com/PuerkitoBio/purell",
    revision = "v1.1.1",
    deps = [
        ":github-com-puerkitobio-urlesc",
    ],
)

go_get(
    name = "github-com-puerkitobio-urlesc",
    get = "github.com/PuerkitoBio/urlesc",
    revision = "de5bf2ad457846296e2031421a34e2568e304e35",
    deps = [
        "//third_party/go:x_net",
    ],
)

go_get(
    name = "github-com-syndtr-goleveldb",
    get = "github.com/syndtr/goleveldb",
    install = [
        "leveldb",
    ],
    revision = "5c35d600f0caac04c20d52438103f1a7aa612598",
    deps = [
        ":github-com-onsi-ginkgo",
        ":github-com-onsi-gomega",
        "//third_party/go:snappy",
    ],
)

go_get(
    name = "github-com-howeyc-gopass",
    get = "github.com/howeyc/gopass",
    revision = "bf9dde6d0d2c004a008c27aaee91170c786f6db8",
    deps = [
        "//third_party/go:x_crypto",
    ],
)

go_get(
    name = "github-com-imdario-mergo",
    get = "github.com/imdario/mergo",
    revision = "v0.3.11",
    visibility = ["PUBLIC"],
)
