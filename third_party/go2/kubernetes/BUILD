KUBERNETES_VERSION = "1.12.6"

#########################
# Top Level K8s targets
#########################

go_get(
    name = "client-go",
    get = "k8s.io/client-go/...",
    install = [
        "discovery",
        "discovery/fake",
        "dynamic",
        "kubernetes",
        "kubernetes/fake",
        "kubernetes/scheme",
        "kubernetes/typed/...",
        "pkg/apis/clientauthentication/...",
        "pkg/version",
        "plugin/pkg/client/auth",
        "plugin/pkg/client/auth/exec",
        "plugin/pkg/client/auth/gcp",
        "plugin/pkg/client/auth/oidc",
        "rest",
        "rest/watch",
        "testing",
        "third_party/forked/golang/template",
        "tools/auth",
        "tools/cache",
        "tools/clientcmd/...",
        "tools/metrics",
        "tools/pager",
        "tools/reference",
        "tools/remotecommand",
        "transport",
        "transport/spdy",
        "util/buffer",
        "util/cert",
        "util/connrotation",
        "util/exec",
        "util/flowcontrol",
        "util/homedir",
        "util/integer",
        "util/jsonpath",
        "util/workqueue",
    ],
    patch = "client_go_client_auth.patch",
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

go_get(
    name = "kubernetes",
    get = "k8s.io/kubernetes/...",
    install = [
        "pkg/credentialprovider/...",
    ],
    licences = ["apache-2.0"],
    repo = "github.com/kubernetes/kubernetes",
    revision = "v%s" % KUBERNETES_VERSION,
    strip = [
        "build",
        "cluster",
        "cmd",
        "Godeps",
        "hack",
        "pkg/cloudprovider",
        "pkg/credentialprovider/azure",
        "pkg/credentialprovider/rancher",
        "pkg/util",
        "staging",
        "test",
        "third_party",
        "translations",
        "vendor",
    ],
    visibility = ["PUBLIC"],
    deps = [
        ":api",
        ":github-com-aws-aws-sdk-go",
        "//third_party/go:docker",
        "//third_party/go:oauth2",
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
        "pkg/labels/selector",
        "pkg/runtime/converter",
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
        ":klog",
    ],
    install = [
        "pkg/apis/meta/v1",
    ],
)

go_get(
    name = "k8s-io-kube-openapi",
    get = "k8s.io/kube-openapi",
    hashes = [
        "41c14baeb7244770b28792e043001dc9e995fb88",
    ],
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

################
# It is dark, you are likely to be eaten by a Grue.
################

go_get(
    name = "github-com-docker-spdystream",
    get = "github.com/docker/spdystream/...",
    hashes = [
        "ad15daf271373e8e87571c072f5c642ca6d3a307",
    ],
    revision = "bc6354cbbc295e925e4c611ffe90c1f287ee54db",
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go:websocket",
    ],
)

go_get(
    name = "github-com-emicklei-go-restful",
    get = "github.com/emicklei/go-restful/...",
    hashes = [
        "41e3ec8d6acd328c55fdaa3a96148355c9f40db5",
    ],
    revision = "26b41036311f2da8242db402557a0dbd09dc83da",
    strip = [
        "examples",
    ],
    visibility = ["PUBLIC"],
)

go_get(
    name = "github-com-garyburd-redigo",
    get = "github.com/garyburd/redigo",
    hashes = [
        "b1bd685ba9d1897ade9eff11256f1cc31f7309a0",
    ],
    install = [
        "internal",
        "redis",
    ],
    revision = "v1.6.0",
)

go_get(
    name = "github-com-ghodss-yaml",
    get = "github.com/ghodss/yaml",
    hashes = [
        "d576c60b16310e8cf3c05b2b44b6e365c4895943",
    ],
    revision = "v1.0.0",
    deps = [
        "//third_party/go:yaml.v2",
    ],
)

go_get(
    name = "github-com-go-openapi-jsonpointer",
    get = "github.com/go-openapi/jsonpointer",
    hashes = [
        "2dd36bdcf7fd3a586fb36ead212df393bfde86c4",
    ],
    revision = "3a0015ad55fa9873f41605d3e8f28cd279c32ab2",
    deps = [
        ":github-com-go-openapi-swag",
    ],
)

go_get(
    name = "github-com-go-openapi-jsonreference",
    get = "github.com/go-openapi/jsonreference",
    hashes = [
        "b9759d0f247ec178a858f0d9deb9340fee6b61a1",
    ],
    revision = "3fb327e6747da3043567ee86abd02bb6376b6be2",
    deps = [
        ":github-com-go-openapi-jsonpointer",
        ":github-com-puerkitobio-purell",
    ],
)

go_get(
    name = "github-com-go-openapi-spec",
    get = "github.com/go-openapi/spec",
    hashes = [
        "4993babd05e875980f2d5f67cf23421b861d087c",
    ],
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
    hashes = [
        "46425726c71c328017ee3bf880383051dfc7df2d",
    ],
    revision = "811b1089cde9dad18d4d0c2d09fbdbf28dbd27a5",
    deps = [
        ":github-com-mailru-easyjson",
        "//third_party/go:yaml.v2",
    ],
)

go_get(
    name = "github-com-google-btree",
    get = "github.com/google/btree",
    hashes = [
        "95d593d70de28ba1f7eab3cefd8b520626ef407a",
    ],
    revision = "e89373fe6b4a7413d7acd6da1725b83ef713e6e4",
)

go_get(
    name = "github-com-google-gofuzz",
    get = "github.com/google/gofuzz",
    hashes = [
        "addd2b7950c6abe71efe83cc46b9f604cf37758c",
    ],
    revision = "24818f796faf91cd76ec7bddd72458fbced7a6c1",
    visibility = ["PUBLIC"],
)

go_get(
    name = "github-com-googleapis-gnostic",
    get = "github.com/googleapis/gnostic",
    hashes = [
        "8f42a28088c75946d3628733430d3b349bc9a9ec",
    ],
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
    hashes = [
        "4ae7eda2212ca8d72fab96c92226b0552cd2d114",
    ],
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
    hashes = [
        "9b2996576f26227b56000680e731d5af5801cc79",
    ],
    revision = "f2b4162afba35581b6d4a50d3b8f34e33c144682",
    visibility = ["PUBLIC"],
    deps = [
        ":github-com-modern-go-reflect2",
    ],
)

go_get(
    name = "github-com-juju-ratelimit",
    get = "github.com/juju/ratelimit",
    hashes = [
        "1576093395bfacdf9c3eeff28a603d6cf7befdc8",
    ],
    revision = "1.0.1",
)

go_get(
    name = "github-com-mailru-easyjson",
    get = "github.com/mailru/easyjson",
    hashes = [
        "9dfeae5c56742932357aa448c6a3fb2b5d632a3a",
    ],
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
    hashes = [
        "88d5021c434289a7397e515c33f711d5869be401",
    ],
    revision = "1.0.3",
)

go_get(
    name = "github-com-modern-go-reflect2",
    get = "github.com/modern-go/reflect2",
    hashes = [
        "8be05e035d11917cbc07a5750cdc56c034a6665b",
    ],
    revision = "1.0.0",
    deps = [
        ":github-com-modern-go-concurrent",
    ],
)

go_get(
    name = "github-com-onsi-ginkgo",
    get = "github.com/onsi/ginkgo",
    hashes = [
        "b0b48584513fa42323b0b17be8136635d3407474",
    ],
    revision = "v1.4.0",
)

go_get(
    name = "github-com-onsi-gomega",
    get = "github.com/onsi/gomega",
    hashes = [
        "be5fbc1892177cb90bb9508c5f3ee44a13183871",
    ],
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
    hashes = [
        "20893369c7df764a3947ff13ee7bf1341259311d",
    ],
    revision = "v2.0.1",
    visibility = ["//third_party/go/..."],
    deps = [
        ":github-com-google-btree",
    ],
)

go_get(
    name = "github-com-puerkitobio-purell",
    get = "github.com/PuerkitoBio/purell",
    hashes = [
        "66568fc693294a7be83e59602f677cfdf89e353d",
    ],
    revision = "975f53781597ed779763b7b65566e74c4004d8de",
    deps = [
        ":github-com-puerkitobio-urlesc",
    ],
)

go_get(
    name = "github-com-puerkitobio-urlesc",
    get = "github.com/PuerkitoBio/urlesc",
    hashes = [
        "1212b2e893eb3cd8f71e98cba2ec7739af5c1f15",
    ],
    revision = "de5bf2ad457846296e2031421a34e2568e304e35",
    deps = [
        "//third_party/go:x_net",
    ],
)

go_get(
    name = "github-com-syndtr-goleveldb",
    get = "github.com/syndtr/goleveldb",
    hashes = [
        "00145e2abc6e42aa7ccf5e6e024b605fdecedb1b",
    ],
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
    hashes = [
        "f28680add98cc3f90b0013c1a94997bdf4f6e67c",
    ],
    revision = "bf9dde6d0d2c004a008c27aaee91170c786f6db8",
    deps = [
        "//third_party/go:x_crypto",
    ],
)

go_get(
    name = "github-com-imdario-mergo",
    get = "github.com/imdario/mergo",
    hashes = [
        "d7ae51a08549f255dbf0315d7ba8a50bf6711df9",
    ],
    revision = "7c29201646fa3de8506f701213473dd407f19646",
    visibility = ["PUBLIC"],
)

go_get(
    name = "github-com-aws-aws-sdk-go",
    get = "github.com/aws/aws-sdk-go",
    hashes = [
        "8f2187e52158dd0a0705f886035a8e871d227b2b",
    ],
    install = [
        "service/ecr",
    ],
    licences = ["apache-2.0"],
    revision = "v1.21.7",
)

go_get(
    name = "po-lint",
    binary = True,
    get = "github.com/coreos/prometheus-operator",
    install = ["cmd/po-lint"],
    licences = ["apache-2.0"],
    revision = "v0.38.0",
    deps = [
        ":apimachinery",
        "//third_party/go:yaml",
    ],
)

go_get(
    name = "klog",
    get = "k8s.io/klog",
    revision = "4f12f780d4781eb828bc6520ae87d7516f30f0b2",
    deps = [
        ":logr",
    ],
)

go_get(
    name = "logr",
    get = "github.com/go-logr/logr",
    revision = "ee2de8da5be655b0a1755b264167fcca1d3193f5",
)
