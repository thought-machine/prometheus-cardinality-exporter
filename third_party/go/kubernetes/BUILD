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
    revision = "cd0308977f05a0de033c0e8ce1860b317e360dca",
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
    revision = "bc6354cbbc295e925e4c611ffe90c1f287ee54db",
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go:websocket",
    ],
)

go_get(
    name = "github-com-emicklei-go-restful",
    get = "github.com/emicklei/go-restful/...",
    revision = "26b41036311f2da8242db402557a0dbd09dc83da",
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
    revision = "v1.6.0",
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
    revision = "3a0015ad55fa9873f41605d3e8f28cd279c32ab2",
    deps = [
        ":github-com-go-openapi-swag",
    ],
)

go_get(
    name = "github-com-go-openapi-jsonreference",
    get = "github.com/go-openapi/jsonreference",
    revision = "3fb327e6747da3043567ee86abd02bb6376b6be2",
    deps = [
        ":github-com-go-openapi-jsonpointer",
        ":github-com-puerkitobio-purell",
    ],
)

go_get(
    name = "github-com-go-openapi-spec",
    get = "github.com/go-openapi/spec",
    revision = "370d9e047557906322be8396e77cb0376be6cb96",
    visibility = ["PUBLIC"],
    deps = [
        ":github-com-go-openapi-jsonpointer",
        ":github-com-go-openapi-jsonreference",
    ],
)

go_get(
    name = "github-com-go-openapi-swag",
    get = "github.com/go-openapi/swag",
    revision = "811b1089cde9dad18d4d0c2d09fbdbf28dbd27a5",
    deps = [
        ":github-com-mailru-easyjson",
        "//third_party/go:yaml.v2",
    ],
)

go_get(
    name = "github-com-google-btree",
    get = "github.com/google/btree",
    revision = "e89373fe6b4a7413d7acd6da1725b83ef713e6e4",
)

go_get(
    name = "github-com-google-gofuzz",
    get = "github.com/google/gofuzz",
    revision = "24818f796faf91cd76ec7bddd72458fbced7a6c1",
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
    revision = "6d7ae43a9ae94f90ed9912252530d97039049c66",
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go:protobuf",
        "//third_party/go:yaml.v2",
    ],
)

go_get(
    name = "github-com-gregjones-httpcache",
    get = "github.com/gregjones/httpcache/...",
    revision = "9cad4c3443a7200dd6400aef47183728de563a38",
    visibility = ["PUBLIC"],
    deps = [
        ":github-com-garyburd-redigo",
        ":github-com-peterbourgon-diskv",
        ":github-com-syndtr-goleveldb",
        "//third_party/go:gomemcache",
    ],
)

go_get(
    name = "github-com-json-iterator-go",
    get = "github.com/json-iterator/go",
    revision = "f2b4162afba35581b6d4a50d3b8f34e33c144682",
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
    revision = "8b799c424f57fa123fc63a99d6383bc6e4c02578",
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
    revision = "975f53781597ed779763b7b65566e74c4004d8de",
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
    revision = "714f901b98fdb3aa954b4193d8cbd64a28d80cad",
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
    revision = "7c29201646fa3de8506f701213473dd407f19646",
    visibility = ["PUBLIC"],
)
