package(default_visibility = ["PUBLIC"])

subinclude("///third_party/subrepos/pleasings//remote")

go_get(
    name = "go-flags",
    get = "github.com/jessevdk/go-flags",
    licences = ["bsd-3-clause"],
    revision = "c0795c8afcf41dd1d786bebce68636c199b3bb45",
)
go_get(
    name = 'errors',
    get = 'github.com/pkg/errors',
    licences = ['bsd-2-clause'],

)
go_get(
    name = 'gofrs.uuid',
    get = 'github.com/gofrs/uuid',
    licences = ['mit'],

)
go_get(
    name = 'grpc',
    get = 'google.golang.org/grpc',
    licences = ['apache-2.0'],
    deps = [
        ':oauth2',
        ':x_sys',
    ],

)
go_get(
    name = 'genproto_protobuf',
    get = 'google.golang.org/genproto/protobuf',
    licences = ['apache-2.0'],
    deps = [':protobuf'],

)
go_get(
    name = 'oauth2',
    get = 'golang.org/x/oauth2',
    licences = ['bsd-3-clause'],
    deps = [
        ':cloud-metadata',
        ':x_net',
    ],

)
go_get(
    name = "protobuf",
    get = "github.com/golang/protobuf/...",
    licences = ["bsd-3-clause"],
    revision = "v1.4.2",
    strip = [
        "proto/proto3_proto",
        "conformance",
    ],
    deps = [
        ":protobuf-v2",
    ],
)

go_get(
    name = 'rpcstatus',
    get = 'google.golang.org/genproto/googleapis/rpc',
    licences = ['apache-2.0'],
    deps = [':protobuf'],

)
go_get(
    name = 'x_net',
    get = "golang.org/x/net/...",
    revision = "65e2d4e15006aab9813ff8769e768bbf4bb667a0",
    licences = ['bsd-3-clause'],
    deps = [
        ':x_crypto',
        ':x_text',
    ],

)
go_get(
    name = "x_sys",
    get = "golang.org/x/sys/...",
    install = [
        "unix",
        "cpu",
    ],
    licences = ["bsd-3-clause"],
    revision = "95b1ffbd15a57cc5abb3f04402b9e8ec0016a52c",
)

go_get(
    name = 'grpc-middleware',
    get = 'github.com/grpc-ecosystem/go-grpc-middleware',
    licences = ['apache-2.0'],
    deps = [
        ':grpc',
        ':opentracing',
        ':x_net',
    ],

)
go_get(
    name = 'opentracing',
    get = 'github.com/opentracing/opentracing-go',
    licences = ['apache-2.0'],
    deps = [':x_net'],

)
go_get(
    name = 'genproto_api',
    get = 'google.golang.org/genproto/googleapis/api',
    licences = ['apache-2.0'],
    deps = [':protobuf'],

)
go_get(
    name = 'cmp',
    get = 'github.com/google/go-cmp/cmp',
    licences = ['bsd-2-clause'],

)
go_get(
    name = 'protoc-gen-go',
    get = 'protoc-gen-go',
    licences = ['bsd-3-clause'],
    deps = [':protobuf'],

)
go_get(
    name = 'logrus',
    get = 'github.com/sirupsen/logrus',
    licences = ['mit'],
    deps = [
        ':x_crypto',
        ':x_sys',
    ],

)
go_get(
    name = 'x_crypto',
    get = 'golang.org/x/crypto/...',
    revision = "e84da0312774c21d64ee2317962ef669b27ffb41",
    licences = ['bsd-3-clause'],
    deps = [':x_sys'],

)

go_get(
    name = "prometheus",
    get = "github.com/prometheus/client_golang/...",
    licences = ["apache-2.0"],
    revision = "v1.0.0",
    deps = [
        ":client_model",
        ":json-iterator",
        ":perks",
        ":procfs",
        ":prometheus_common",
        ":protobuf",
        ":x_net",
    ],
)

go_get(
    name = 'json-iterator',
    get = 'github.com/json-iterator/go',
    licences = ['MIT'],
    deps = [
        ':concurrent',
        ':reflect2',
    ],

)

go_get(
    name = 'procfs',
    get = 'github.com/prometheus/procfs',
    licences = ['apache-2.0'],
    deps = [':x_sync'],

)
go_get(
    name = 'goimports',
    get = 'goimports',
    licences = ['bsd-3-clause'],
    deps = [':x_tools'],

)


go_get(
    name = 'testify',
    get = 'github.com/stretchr/testify/...',
    licences = ['MIT'],
    deps = [
        ':difflib',
        ':go-spew',
        ':objx',
        ':yaml.v3',
    ],

)

go_get(
    name = "difflib",
    get = "github.com/pmezard/go-difflib/...",
    licences = ["bsd-3-clause"],
    revision = "v1.0.0",
)
go_get(
    name = 'go-spew',
    get = 'github.com/davecgh/go-spew/spew',
    licences = ['isc'],

)
go_get(
    name = 'objx',
    get = 'github.com/stretchr/objx',
    licences = ['mit'],

)
go_get(
    name = 'yaml.v3',
    get = 'gopkg.in/yaml.v3',
    licences = ['apache-2.0'],

)

go_get(
    name = 'glog',
    get = 'github.com/golang/glog',
    licences = ['apache-2.0'],

)
go_get(
    name = "gogo-protobuf",
    get = "github.com/gogo/protobuf",
    install = [
        "proto",
        "sortkeys",
    ],
    licences = ["bsd-3-clause"],
    revision = "342cbe0a04158f6dcb03ca0079991a51a4248c02",
    strip = ["proto/proto3_proto"],
)
go_get(
    name = 'inf.v0',
    get = 'gopkg.in/inf.v0',
    licences = ['bsd-3-clause'],

)
go_get(
    name = 'lru',
    get = 'github.com/hashicorp/golang-lru',
    licences = ['mpl-2.0'],

)
go_get(
    name = 'pflag',
    get = 'github.com/spf13/pflag',
    licences = ['bsd-3-clause'],

)
go_get(
    name = 'github-com-docker-spdystream',
    get = 'github.com/docker/spdystream',
    deps = [':websocket'],

)
go_get(
    name = 'github-com-ghodss-yaml',
    get = 'github.com/ghodss/yaml',
    deps = [':yaml.v2'],

)
go_get(
    name = 'github-com-google-gofuzz',
    get = 'github.com/google/gofuzz',

)
go_get(
    name = 'k8s-io-kube-openapi',
    get = 'k8s.io/kube-openapi',
    deps = [
        ':yaml.v2',
        ':github-com-emicklei-go-restful',
        ':github-com-go-openapi-spec',
        ':github-com-googleapis-gnostic',
    ],

)
go_get(
    name = 'x_time_rate',
    get = 'golang.org/x/time/rate',
    licences = ['bsd-3-clause'],
    deps = [':x_net'],

)
go_get(
    name = 'api',
    get = 'k8s.io/api',
    deps = [':apimachinery'],

)
go_get(
    name = 'github-com-gregjones-httpcache',
    get = 'github.com/gregjones/httpcache',
    deps = [
        ':gomemcache',
        ':github-com-garyburd-redigo',
        ':github-com-peterbourgon-diskv',
        ':github-com-syndtr-goleveldb',
    ],

)
go_get(
    name = 'github-com-howeyc-gopass',
    get = 'github.com/howeyc/gopass',
    deps = [':x_crypto'],

)
go_get(
    name = 'github-com-imdario-mergo',
    get = 'github.com/imdario/mergo',

)
go_get(
    name = 'github-com-juju-ratelimit',
    get = 'github.com/juju/ratelimit',

)

go_get(
    name = "x_text",
    get = "golang.org/x/text/...",
    revision = "v0.3.0",
    strip = [
        "cmd",
        "message/pipeline",
    ],
)

go_get(
    name = "x_mod",
    get = "golang.org/x/mod",
    install = [
        "semver",
    ],
    licences = ["bsd-3-clause"],
    revision = "e5e73c1b9c72835114eb6daab038373d39515006",
)


go_get(
    name = "x_tools",
    get = "golang.org/x/tools",
    install = [
        "cmd/...",
        "imports",
        "go/packages",
    ],
    licences = ["bsd-3-clause"],
    revision = "5bcca83a78812bd91ce8cb29be4fc4521cdc8f6f",
    deps = [
        ":x_crypto",
        ":x_errors",
        ":x_mod",
        ":x_net",
        ":x_sync",
    ],
)

go_get(
    name="backoff",
    get = "github.com/cenkalti/backoff",
    licences = ["mit"],
    revision="31cc31bb63269a3c813d1f26e0ab058452b7d803",
    deps = [
        ":go-spew",
        ":difflib",
        ":x_sys",
        ":testify",
    ],
)


go_get(
    name = "websocket",
    get = "github.com/gorilla/websocket",
    licences = ["bsd-2-clause"],
    revision = "v1.4.2",
)

go_get(
    name = "yaml.v2",
    get = "gopkg.in/yaml.v2",
    revision = "f7716cbe52baa25d2e9b0d0da546fcf909fc16b4",
)

go_get(
    name = "concurrent",
    get = "github.com/modern-go/concurrent/...",
    revision = "1.0.3",
)

go_get(
    name = "cloud-metadata",
    get = "cloud.google.com/go/compute/metadata",
    licences = ["apache-2.0"],
)

go_get(
    name = "reflect2",
    get = "github.com/modern-go/reflect2/...",
    licences = ["apache-2.0"],
    revision = "v1.0.1",
    deps = [
        ":concurrent",
    ],
)

go_get(
    name = "gomemcache",
    get = "github.com/bradfitz/gomemcache",
    install = [
        "memcache",
    ],
    licences = ["apache-2.0"],
    revision = "1952afaa557dc08e8e0d89eafab110fb501c1a2b",
)

go_get(
    name = "x_sync",
    get = "golang.org/x/sync/...",
    licences = ["bsd-3-clause"],
    revision = "fd80eb99c8f653c847d294a001bdf2a3a6f768f5",
    deps = [
        ":x_net",
    ],
)

go_get(
    name = "golang_protobuf_extensions",
    get = "github.com/matttproud/golang_protobuf_extensions/...",
    licences = ["apache-2.0"],
    revision = "c12348ce28de40eed0136aa2b644d0ee0650e56c",
    deps = [
        ":protobuf",
    ],
)

go_get(
    name = "kingpin",
    get = "gopkg.in/alecthomas/kingpin.v2",
    licences = ["mit"],
    revision = "v2.2.6",
    deps = [
        ":template",
        ":units",
    ],
)


go_get(
    name = "template",
    get = "github.com/alecthomas/template",
    licences = ["bsd-3-clause"],
    revision = "a0175ee3bccc567396460bf5acd36800cb10c49c",
)

go_get(
    name = "units",
    get = "github.com/alecthomas/units",
    licences = ["mit"],
    revision = "2efee857e7cfd4f3d0138cc3cbb1b4966962b93a",
)

go_get(
    name = "snappy",
    get = "github.com/golang/snappy",
    licences = ["bsd-3-clause"],
    revision = "553a641470496b2327abcac10b36396bd98e45c9",
)

go_get(
    name = "uuid",
    get = "github.com/pborman/uuid",
    licences = ["bsd-3-clause"],
    revision = "cccd189d45f7ac3368a0d127efb7f4d08ae0b655",
)

go_get(
    name = "perks",
    get = "github.com/beorn7/perks/...",
    licences = ["mit"],
    revision = "4c0e84591b9aa9e6dcfdf3e020114cd81f89d5f9",
)

go_get(
    name = "protobuf-v2",
    get = "google.golang.org/protobuf/...",
    install = [
        "encoding/...",
        "internal/...",
        "proto",
        "reflect/...",
        "runtime/...",
        "types/...",
    ],
    licences = ["bsd-3-clause"],
    revision = "v1.24.0",
    deps = [
        ":cmp",
    ],
)

go_get(
    name = "client_model",
    get = "github.com/prometheus/client_model/...",
    licences = ["apache-2.0"],
    revision = "fd36f4220a901265f90734c3183c5f0c91daa0b8",
    deps = [
        ":protobuf",
    ],
)

go_get(
    name = "prometheus_common",
    get = "github.com/prometheus/common",
    install = [
        "expfmt",
        "model",
        "internal/...",
        "log",
    ],
    licences = ["apache-2.0"],
    revision = "v0.6.0",
    deps = [
        ":client_model",
        ":golang_protobuf_extensions",
        ":kingpin",
        ":logrus",
        ":protobuf",
    ],
)

go_get(
    name = "mockgen",
    binary = True,
    get = [],
    install = ["github.com/golang/mock/mockgen"],
    licences = ["apache-2.0"],
    deps = [":mock"],
)

go_get(
    name = "mock",
    get = "github.com/golang/mock/...",
    licences = ["apache-2.0"],
    revision = "93f6609a15b7de76bd49259f1f9a6b58df358936",
    deps = [
        ":x_net",
    ],
)
