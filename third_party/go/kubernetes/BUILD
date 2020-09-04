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
    hashes = [
        "0b8b8805aef82c9b95f858a735f6c935d286c2e20a28d42c35455e38b3d3b32d",
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
    hashes = [
        "4094e6f14afea4269bc3f22ee0e784bda87c7e125282bd5c4afc194bac6ca6b3",
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
    hashes = [
        "aca0057b54be25950970357eb50a1a99c88a85ca91e893749557b73ef5b0f927",
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
    hashes = [
        "461de3035a36400e1da865c7829940455021bf65c5c2ef7427c985f856659d81",
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
    hashes = [
        "844b52f12edfe7d802820f6d61395e740e7f8a99ca1ea3dc98e0e29d116b6ae6",
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
    hashes = [
        "20312757e5d8d322fa032d0cf6a7bfb7ffdc47dfa601d2b6c7165ed75acf201a",
    ],
)

go_get(
    name = "github-com-garyburd-redigo",
    get = "github.com/garyburd/redigo",
    install = [
        "internal",
        "redis",
    ],
    revision = "v1.6.2",
    hashes = [
        "0f8f1e3c5f3b8c9c095f7459f376b6d9c406b20930b4f2d087a9e3ac75283574",
    ],
)

go_get(
    name = "github-com-ghodss-yaml",
    get = "github.com/ghodss/yaml",
    revision = "v1.0.0",
    deps = [
        "//third_party/go:yaml.v2",
    ],
    hashes = [
        "29db4b0ee39b89bcf365ccc80d167c799f776d264157aef90ef920e89df59ed8",
    ],
)

go_get(
    name = "github-com-go-openapi-jsonpointer",
    get = "github.com/go-openapi/jsonpointer",
    revision = "v0.19.3",
    deps = [
        ":github-com-go-openapi-swag",
    ],
    hashes = [
        "1493e329a656f1e1fb555909a495f6d7332996171c856918494497fb89a700ad",
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
    hashes = [
        "28f9615a40c6e5b9c33c4bdb13ec48c7bc2c31001bb99567a6329eca12d2ec74",
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
    hashes = [
        "1877db714cc82d1f50875a48f57ef041cb19fb62d07da35814e8240ae14f80ab",
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
    hashes = [
        "b1632d2677e085d7ddc3f2481f8636d704883d9f878bd9351a0cf9bcdaad51c7",
    ],
)

go_get(
    name = "github-com-google-btree",
    get = "github.com/google/btree",
    revision = "479b5e81b0a93ec038d201b0b33d17db599531d3",
    hashes = [
        "db602d714288fa8da4e35e217830bf3fe440239237896ac8feb8815306a94678",
    ],
)

go_get(
    name = "github-com-google-gofuzz",
    get = "github.com/google/gofuzz",
    revision = "v1.1.0",
    visibility = ["PUBLIC"],
    hashes = [
        "cdcc183fcf44937e0f3740b208be885e5c6e0ed7ef00257642c7e81cf79cdccb",
    ],
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
    hashes = [
        "ea6baba693c7343e382e59cdb13c3a0da1006ec5c65c01368a8106dafbff44b1",
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
    hashes = [
        "95bd4a4abe4b661417e6d901f9fc366425315dd6f230dbe93c431bdb90840add",
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
    hashes = [
        "25642190fd903186dc647a63bfaab254a870134e755124a3ff8560ac5ee4a8f6",
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
    hashes = [
        "18fd7d82bd6cb62c2adb6d39fe9199e58a72904d21032210ffd22aed07d0d6f3",
    ],
)

go_get(
    name = "github-com-juju-ratelimit",
    get = "github.com/juju/ratelimit",
    revision = "1.0.1",
    hashes = [
        "b935cab54e34e430d65be2ce2c4ba1e8dff23197c6324d25c00c49de9ec586ab",
    ],
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
    hashes = [
        "5c001c676ff0f733cfc972922cdbd575b5f28c04d6903ecb45dccf431bb12198",
    ],
)

go_get(
    name = "github-com-josharian-intern",
    get = "github.com/josharian/intern",
    revision = "v1.0.0",

    hashes = [
        "f41f402061cc99ffd37392923e1e26abe1fa772ff4a86f6200874149f0939613",
    ],
)
go_get(
    name = "github-com-modern-go-concurrent",
    get = "github.com/modern-go/concurrent",
    revision = "1.0.3",
    hashes = [
        "810908664f8b0073fcc52d8771586936d53bf4f5bff99e5760a2b7d46b2423e0",
    ],
)

go_get(
    name = "github-com-modern-go-reflect2",
    get = "github.com/modern-go/reflect2",
    revision = "1.0.0",
    deps = [
        ":github-com-modern-go-concurrent",
    ],
    hashes = [
        "0029d29b14c35dc01c053c7e390af08c3f7025578c19bb782009eed4f7718cf1",
    ],
)

go_get(
    name = "github-com-onsi-ginkgo",
    get = "github.com/onsi/ginkgo",
    revision = "v1.4.0",
    hashes = [
        "3165362a976871d3e13346b04516fdb9ffbad22747eac20901d3a2cb660d1aa2",
    ],
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
    hashes = [
        "f7f33e3af40628150971584e860f3c6243ab93f6ba5aa6794a9fbed29b88512d",
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
    hashes = [
        "67379f92bfc57efb16d46eb8438b405039125a470b84d2e2a5938421d04fcf10",
    ],
)

go_get(
    name = "github-com-puerkitobio-purell",
    get = "github.com/PuerkitoBio/purell",
    revision = "v1.1.1",
    deps = [
        ":github-com-puerkitobio-urlesc",
    ],
    hashes = [
        "98cdbe3e0ffc8d86783bb75fabb6b3f896148f6e9c45ff3b06bb8b4dd3505ed5",
    ],
)

go_get(
    name = "github-com-puerkitobio-urlesc",
    get = "github.com/PuerkitoBio/urlesc",
    revision = "de5bf2ad457846296e2031421a34e2568e304e35",
    deps = [
        "//third_party/go:x_net",
    ],
    hashes = [
        "b44ea68a1b048562e23599e93ff16fb7ebf77bd12a7d2e7099ee00e91f5c9a2d",
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
    hashes = [
        "b70b72a06c618925aee4d8ba2bf7b6aa6f8479258aa13ec3c35cb10600e646b7",
    ],
)

go_get(
    name = "github-com-howeyc-gopass",
    get = "github.com/howeyc/gopass",
    revision = "bf9dde6d0d2c004a008c27aaee91170c786f6db8",
    deps = [
        "//third_party/go:x_crypto",
    ],
    hashes = [
        "857cfaedfb18d71de10e19c5dceddbd017a6619210bec75a4933bc15d0a8b6bc",
    ],
)

go_get(
    name = "github-com-imdario-mergo",
    get = "github.com/imdario/mergo",
    revision = "v0.3.11",
    visibility = ["PUBLIC"],
    hashes = [
        "6fe5b18a68c6dd0d3f06925d5403a843f824f50f1b63356d7e179511211500af",
    ],
)
