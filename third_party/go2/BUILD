package(default_visibility = ["PUBLIC"])

go_get(
    name = "version",
    get = "github.com/hashicorp/go-version",
    licences = ["mpl-2.0"],
    revision = "192140e6f3e645d971b134d4e35b5191adb9dfd3",
)

go_get(
    name = "aws-iam-authenticator",
    get = "sigs.k8s.io/aws-iam-authenticator/...",
    install = [
        "pkg/...",
        "vendor/...",
    ],
    licences = ["apache-2.0"],
    revision = "v0.5.0",
    deps = [
        ":aws-sdk-go",
    ],
)

go_get(
    name = "aws-sdk-go",
    get = "github.com/aws/aws-sdk-go/...",
    install = [
        "aws/...",
        "internal/...",
        "private/...",
        "service/...",
        "vendor/...",
    ],
    licences = ["apache-2.0"],
    revision = "v1.33.5",
    deps = [
        ":go-jmespath",
        ":x_net",
    ],
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
    name = "units",
    get = "github.com/alecthomas/units",
    licences = ["mit"],
    revision = "2efee857e7cfd4f3d0138cc3cbb1b4966962b93a",
)

go_get(
    name = "template",
    get = "github.com/alecthomas/template",
    licences = ["bsd-3-clause"],
    revision = "a0175ee3bccc567396460bf5acd36800cb10c49c",
)

go_get(
    name = "confluent-kafka",
    get = "github.com/confluentinc/confluent-kafka-go/kafka",
    licences = ["apache-2.0"],
    patch = "confluent_no_pkgconfig.patch",
    revision = "v1.1.0",
    visibility = [
        "//common/go/kafka/...",
        "//services/kafka/kafkagenie/...",
        "//vault/infrastructure/postmerge_tests/kafka/tests/...",
        "//vault/infrastructure/topicadmin/...",
        "//vault/kernel/data_loader/common/sat/...",
    ],
)

go_get(
    name = "handlers",
    get = "github.com/gorilla/handlers",
    licences = ["bsd-2-clause"],
    revision = "v1.3.0",
)

go_get(
    name = "cronexpr",
    get = "github.com/thought-machine/cronexpr",
    licences = ["apache-2.0"],
    revision = "1.1.1",
)

# From version f8560614 (2019-04-22), this has a circular dependency on
# golang.org/x/net because of acme/autocert, which we don't use; resolve this by
# stripping out acme/autocert and patching out that dependency in the metadata
go_get(
    name = "x_crypto",
    get = "golang.org/x/crypto/...",
    licences = ["bsd-3-clause"],
    patch = "x_crypto_no_net.patch",
    revision = "123391ffb6de907695e1066dc40c1ff09322aeb6",  # 2020-07-28
    strip = ["acme/autocert"],
    deps = [":x_sys"],
)

go_get(
    name = "errors",
    get = "github.com/pkg/errors",
    licences = ["bsd-2-clause"],
    revision = "v0.9.1",
)

go_get(
    name = "scuttle",
    binary = True,
    get = "github.com/redboxllc/scuttle",
    licences = ["mit"],
    revision = "v1.0.1",
)

go_get(
    name = "glog",
    get = "github.com/golang/glog",
    licences = ["apache-2.0"],
    revision = "23def4e6c14b4da8ac2ed8007337bc5eb5007998",
)

go_get(
    name = "go-bindata",
    binary = True,
    get = "github.com/kevinburke/go-bindata/...",
    licences = ["CC0 1.0 Universal"],
    revision = "46eb4c183bfc1ebb527d9d19bcded39476302eb8",
    strip = ["testdata"],
)

go_get(
    name = "cleanhttp",
    get = "github.com/hashicorp/go-cleanhttp",
    licences = ["MPL-2.0"],
    revision = "eda1e5db218aad1db63ca4642c8906b26bcf2744",
)

go_get(
    name = "go-flags",
    get = "github.com/jessevdk/go-flags",
    licences = ["bsd-3-clause"],
    revision = "c0795c8afcf41dd1d786bebce68636c199b3bb45",
)

go_get(
    name = "resize",
    get = "github.com/nfnt/resize",
    licences = ["isc"],
    revision = "891127d8d1b52734debe1b3c3d7e747502b6c366",
)

go_get(
    name = "grpc-middleware",
    get = "github.com/grpc-ecosystem/go-grpc-middleware/...",
    licences = ["apache-2.0"],
    # TODO(TM-16410): change the revision version to 1.3 once it's officially released
    revision = "1bcd9c1034fe33f4492454ca8864118a832ca9a0",
    strip = [
        "testing",
        "logging",
        "auth",
        "tags/logrus",
        "tags/zap",
    ],
    deps = [
        ":grpc",
        ":opentracing",
        ":x_net",
    ],
)

go_get(
    name = "grpc-prometheus",
    get = "github.com/grpc-ecosystem/go-grpc-prometheus",
    licences = ["apache-2.0"],
    revision = "34abd90a014618f61222a1b0a7b7eb834a2d0dc3",
    deps = [
        ":grpc",
        ":prometheus",
        ":x_net",
    ],
)

go_get(
    name = "gziphandler",
    get = "github.com/NYTimes/gziphandler",
    licences = ["apache-2.0"],
    revision = "v1.1.1",
)

go_get(
    name = "go-jmespath",
    get = "github.com/jmespath/go-jmespath",
    licences = ["apache-2.0"],
    revision = "3433f3ea46d9f8019119e7dd41274e112a2359a9",
    strip = ["fuzz"],
)

go_get(
    name = "go-license",
    get = "github.com/ryanuber/go-license",
    licences = ["mit"],
    revision = "3878fc8f33ea4637cf02cf8a0bf926288a246ceb",
)

go_get(
    name = "logging",
    get = "github.com/op/go-logging",
    licences = ["bsd-3-clause"],
    revision = "dfaf3dff9b631bc4236201d90d41ee0de9202889",
)

go_get(
    name = "semver",
    get = "github.com/coreos/go-semver/semver",
    licences = ["apache-2.0"],
    revision = "8ab6407b697782a06568d4b7f1db25550ec2e4c6",
)

go_get(
    name = "go-spew",
    get = "github.com/davecgh/go-spew/spew",
    licences = ["isc"],
    revision = "v1.1.0",
)

go_get(
    name = "go.uuid",
    get = "github.com/satori/go.uuid",
    licences = ["mit"],
    revision = "0aa62d5ddceb50dbcb909d790b5345affd3669b6",
    visibility = ["//third_party/go/..."],
)

go_get(
    name = "gofrs.uuid",
    get = "github.com/gofrs/uuid",
    licences = ["mit"],
    revision = "v3.2.0",
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
    name = "grpc-gateway",
    get = "github.com/grpc-ecosystem/grpc-gateway/...",
    licences = ["bsd-3-clause"],
    revision = "v1.14.6",
    strip = ["examples"],
    deps = [
        ":genproto_api",
        ":genproto_protobuf",
        ":glog",
        ":grpc",
        ":protobuf",
        ":x_net",
        ":yaml",
    ],
)

go_get(
    name = "grpc",
    exported_deps = [
        ":genproto_protobuf",
        ":protobuf",
        ":rpcstatus",
        ":x_net",
    ],
    get = "google.golang.org/grpc",
    install = [
        "",
        "attributes",
        "backoff",
        "balancer",
        "balancer/base",
        "balancer/grpclb/...",
        "balancer/roundrobin",
        "binarylog/...",
        "codes",
        "connectivity",
        "credentials/...",
        "encoding",
        "encoding/gzip",
        "encoding/proto",
        "grpclog",
        "health",
        "health/grpc_health_v1",
        "internal/...",
        "keepalive",
        "metadata",
        "naming",
        "peer",
        "resolver",
        "resolver/dns",
        "resolver/passthrough",
        "serviceconfig",
        "stats",
        "status",
        "tap",
    ],
    licences = ["apache-2.0"],
    repo = "github.com/grpc/grpc-go",
    revision = "v1.27.0",
    deps = [
        ":oauth2",
        ":rpcstatus",
        ":x_sys",
    ],
)

GENPROTO_VERSION = "a023cd5227bd25fd1f5c5633743ff3eacc93d169"

go_get(
    name = "rpcstatus",
    get = "google.golang.org/genproto/googleapis/rpc/...",
    licences = ["apache-2.0"],
    revision = GENPROTO_VERSION,
    deps = [":protobuf"],
)

go_get(
    name = "genproto_api",
    get = "google.golang.org/genproto/googleapis/api",
    install = [
        "annotations",
        "httpbody",
    ],
    licences = ["apache-2.0"],
    revision = GENPROTO_VERSION,
    deps = [":protobuf"],
)

go_get(
    name = "genproto_type",
    get = "google.golang.org/genproto/googleapis/type/expr",
    licences = ["apache-2.0"],
    revision = GENPROTO_VERSION,
    deps = [":protobuf"],
)

go_get(
    name = "iam",
    get = "google.golang.org/genproto/googleapis/iam/...",
    licences = ["apache-2.0"],
    revision = GENPROTO_VERSION,
    deps = [
        ":genproto_api",
        ":genproto_protobuf",
        ":genproto_type",
        ":grpc",
        ":protobuf",
    ],
)

go_get(
    name = "longrunning",
    get = "google.golang.org/genproto/googleapis/longrunning/...",
    licences = ["apache-2.0"],
    revision = GENPROTO_VERSION,
    deps = [
        ":genproto_api",
        ":genproto_protobuf",
        ":grpc",
        ":protobuf",
    ],
)

go_get(
    name = "genproto_devtools_containeranalysis",
    get = "google.golang.org/genproto/googleapis/devtools/containeranalysis/...",
    licences = ["apache-2.0"],
    revision = GENPROTO_VERSION,
    deps = [
        ":genproto_api",
        ":genproto_protobuf",
        ":grpc",
        ":iam",
        ":longrunning",
        ":protobuf",
    ],
)

go_get(
    name = "genproto_protobuf",
    get = "google.golang.org/genproto/protobuf/...",
    licences = ["apache-2.0"],
    revision = GENPROTO_VERSION,
    deps = [
        ":protobuf",
    ],
)

go_get(
    name = "genproto_grafeas",
    get = "google.golang.org/genproto/googleapis/grafeas/...",
    licences = ["apache-2.0"],
    revision = GENPROTO_VERSION,
    deps = [
        ":genproto_api",
        ":genproto_protobuf",
        ":grpc",
        ":rpcstatus",
    ],
)

go_get(
    name = "inf.v0",
    get = "gopkg.in/inf.v0",
    licences = ["bsd-3-clause"],
    revision = "3887ee99ecf07df5b447e9b00d9c0b2adaa9f3e4",
)

go_get(
    name = "jsonenums",
    binary = True,
    get = "github.com/campoy/jsonenums/...",
    licences = ["apache-2.0"],
    revision = "5bfdb73f884302bca3b3246048d27992ba1c08f3",
    deps = [
        ":x_tools",
    ],
)

go_get(
    name = "jwt-go",
    get = "github.com/dgrijalva/jwt-go",
    licences = ["mit"],
    revision = "06ea1031745cb8b3dab3f6a236daf2b0aa468b7e",
)

go_get(
    name = "jwt",
    binary = True,
    get = [],
    install = [
        "github.com/dgrijalva/jwt-go/cmd/jwt/...",
    ],
    licences = ["mit"],
    deps = [":jwt-go"],
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

go_get(
    name = "goimports",
    binary = True,
    get = [],
    install = ["golang.org/x/tools/cmd/goimports"],
    licences = ["bsd-3-clause"],
    deps = [":x_tools"],
)

go_get(
    name = "stringer",
    binary = True,
    get = [],
    install = ["golang.org/x/tools/cmd/stringer"],
    licences = ["bsd-3-clause"],
    deps = [":x_tools"],
)

go_get(
    name = "moduluschecking",
    get = "github.com/AntoineAugusti/moduluschecking/...",
    licences = ["mit"],
    revision = "cd83b28f27b5103fbdc97a7352d62a3c9f1d698d",
)

go_get(
    name = "mux",
    get = "github.com/gorilla/mux",
    licences = ["bsd-3-clause"],
    revision = "v1.6.2",
)

go_get(
    name = "x_net",
    get = "golang.org/x/net/...",
    licences = ["bsd-3-clause"],
    revision = "e086a090c8fdb9982880f0fb6e3db47af1856533",  # 21 Apr 2020
    strip = [
        "http2/h2demo",  # demo tries to compile some google storage rot
    ],
    deps = [
        ":x_crypto",
        ":x_text",
    ],
)

go_get(
    name = "perks",
    get = "github.com/beorn7/perks/...",
    licences = ["mit"],
    revision = "4c0e84591b9aa9e6dcfdf3e020114cd81f89d5f9",
)

go_get(
    name = "pgx",
    get = "github.com/jackc/pgx/...",
    licences = ["mit"],
    revision = "cba88a4b7d0b6a614b6a1de5f1f7477922961393",
    strip = [
        "log",
        "examples",
        "pgtype/testutil",
    ],
    deps = [
        ":decimal",
        ":errors",
        ":go.uuid",
    ],
)

go_get(
    name = "decimal",
    get = "github.com/shopspring/decimal",
    licences = ["mit"],
    revision = "v1.1.0",
)

go_get(
    name = "apd",
    get = "github.com/cockroachdb/apd",
    licences = ["apache-2.0"],
    revision = "v2.0.1",
    deps = [
        ":errors",
    ],
)

go_get(
    name = "pq",
    get = "github.com/lib/pq/...",
    licences = ["MIT"],
    revision = "88edab0803230a3898347e77b474f8c1820a1f20",
)

go_get(
    name = "procfs",
    get = "github.com/prometheus/procfs/...",
    licences = ["apache-2.0"],
    revision = "v0.0.3",
    deps = [
        ":x_sync",
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
    name = "protoc-gen-go",
    binary = True,
    get = [],
    install = ["github.com/golang/protobuf/protoc-gen-go"],
    licences = ["bsd-3-clause"],
    deps = [":protobuf"],
)

go_get(
    name = "protoc-gen-grpc-gateway",
    binary = True,
    get = [],
    install = ["github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"],
    licences = ["bsd-3-clause"],
    deps = [
        ":glog",
        ":grpc",
        ":grpc-gateway",
        ":protobuf",
        ":x_net",
    ],
)

go_get(
    name = "snappy",
    get = "github.com/golang/snappy",
    licences = ["bsd-3-clause"],
    revision = "553a641470496b2327abcac10b36396bd98e45c9",
)

go_get(
    name = "appc_spec",
    get = "github.com/appc/spec",
    install = [
        "...",
        "vendor/...",
    ],
    licences = ["apache-2.0"],
    revision = "20ef57ba534045e0aed7b76c6c40b1338b008fc4",
    deps = [
        ":errorutil",
        ":inf.v0",
        ":pflag",
        ":semver",
        ":x_net",
    ],
)

go_get(
    name = "errorutil",
    get = "go4.org/errorutil",
    licences = ["apache-2.0"],
    revision = "03efcb870d84809319ea509714dd6d19a1498483",
)

go_get(
    name = "pflag",
    get = "github.com/spf13/pflag",
    licences = ["bsd-3-clause"],
    revision = "298182f68c66c05229eb03ac171abe6e309ee79a",
)

go_get(
    name = "sqlx",
    get = "github.com/jmoiron/sqlx/...",
    licences = ["mit"],
    revision = "05cef0741ade10ca668982355b3f3f0bcf0ff0a8",
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
    name = "testify",
    get = "github.com/stretchr/testify",
    install = [
        "assert",
        "require",
        "mock",
        "suite",
        "vendor/...",
    ],
    licences = ["MIT"],
    revision = "v1.6.1",
    deps = [
        ":difflib",
        ":go-spew",
        ":objx",
        ":yaml.v3",
    ],
)

go_get(
    name = "difflib",
    get = "github.com/pmezard/go-difflib/...",
    licences = ["bsd-3-clause"],
    revision = "v1.0.0",
)

go_get(
    name = "objx",
    get = "github.com/stretchr/objx",
    licences = ["mit"],
    revision = "v0.1.0",
)

go_get(
    name = "x_text",
    get = "golang.org/x/text/...",
    licences = ["bsd-3-clause"],
    revision = "v0.3.0",
    strip = [
        "cmd",
        "message/pipeline",
    ],
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
    name = "x_mod",
    get = "golang.org/x/mod",
    install = [
        "semver",
    ],
    licences = ["bsd-3-clause"],
    revision = "e5e73c1b9c72835114eb6daab038373d39515006",
)

go_get(
    name = "uuid",
    get = "github.com/pborman/uuid",
    licences = ["bsd-3-clause"],
    revision = "cccd189d45f7ac3368a0d127efb7f4d08ae0b655",
)

go_get(
    name = "hashicorp_go-sockaddr",
    get = "github.com/hashicorp/go-sockaddr",
    licences = ["MPL-2.0"],
    revision = "6d291a969b86c4b633730bfc6b8b9d64c3aafed9",
)

go_get(
    name = "ryanuber_go-glob",
    get = "github.com/ryanuber/go-glob",
    licences = ["MIT"],
    revision = "256dc444b735e061061cf46c809487313d5b0065",
)

go_get(
    name = "sermoDigital_jose",
    get = "github.com/SermoDigital/jose",
    install = [
        ".",
        "crypto",
        "jws",
        "jwt",
    ],
    licences = ["MIT"],
    revision = "803625baeddc3526d01d321b5066029f53eafc81",
)

go_get(
    name = "armon_go-metrics",
    get = "github.com/armon/go-metrics",
    licences = ["MIT"],
    revision = "f0300d1749da6fa982027e449ec0c7a145510c3c",
    deps = [
        ":hashicorp_go-immutable-radix",
    ],
)

go_get(
    name = "hashicorp_go-immutable-radix",
    get = "github.com/hashicorp/go-immutable-radix",
    licences = ["MPL-2.0"],
    revision = "0146a9aba1948ded4ed290cfd3fded2c15313f63",
    deps = [
        ":lru",
    ],
)

go_get(
    name = "armon_go-radix",
    get = "github.com/armon/go-radix/...",
    licences = ["MIT"],
    revision = "1a2de0c21c94309923825da3df33a4381872c795",
    deps = [
    ],
)

go_get(
    name = "oklog_run",
    get = "github.com/oklog/run",
    licences = ["apache-2.0"],
    revision = "6934b124db28979da51d3470dadfa34d73d72652",
    deps = [
    ],
)

go_get(
    name = "mitchellh_go-testing-interface",
    get = "github.com/mitchellh/go-testing-interface",
    licences = ["MIT"],
    revision = "6d0b8010fcc857872e42fc6c931227569016843c",
)

go_get(
    name = "hashicorp_yamux",
    get = "github.com/hashicorp/yamux",
    licences = ["MPL-2.0"],
    revision = "df201c70410deac145ff09e1126f3f8834ead1ec",
)

go_get(
    name = "hashicorp_go-plugin",
    get = "github.com/hashicorp/go-plugin",
    install = [
        ".",
        "internal/plugin",
    ],
    licences = ["MPL-2.0"],
    revision = "809113480b559c989ea9cfcff62e9d387961f60b",
    deps = [
        ":grpc",
        ":hashicorp_yamux",
        ":hclog",
        ":mitchellh_go-testing-interface",
        ":oklog_run",
        ":protobuf",
        ":x_net",
    ],
)

go_get(
    name = "mgutz_ansi",
    get = "github.com/mgutz/ansi",
    licences = ["mit"],
    revision = "9520e82c474b0a04dd04f8a40959027271bab992",
    deps = [
        ":mattn_go-colorable",
    ],
)

go_get(
    name = "mattn_go-isatty",
    get = "github.com/mattn/go-isatty",
    licences = ["MIT"],
    revision = "2a2f0ea997f99becfd8423e4f337d171753fe685",
    deps = [
        ":x_sys",
    ],
)

go_get(
    name = "hashicorp_go-memdb",
    get = "github.com/hashicorp/go-memdb",
    licences = ["MPL-2.0"],
    revision = "e3965aaa499eb83f14b147b83866658476828b3f",
    deps = [
        ":hashicorp_go-immutable-radix",
    ],
)

go_get(
    name = "ini",
    get = "gopkg.in/ini.v1",
    licences = ["apache-2.0"],
    revision = "v1.57.0",
)

go_get(
    name = "elazarl_go-bindata-assetfs",
    get = "github.com/elazarl/go-bindata-assetfs",
    licences = ["BSD 2-Clause"],
    revision = "d0111fe6fb11d467f8aa090457c2edaa4a2d7bab",
)

go_get(
    name = "hashicorp_go-uuid",
    get = "github.com/hashicorp/go-uuid",
    licences = ["MPL-2.0"],
    revision = "f32f5fe8d6a86bfd2a4f58998002f7b0b91cf78f",
)

go_get(
    name = "jefferai_jsonx",
    get = "github.com/jefferai/jsonx",
    licences = ["MPL-2.0"],
    patch = "jefferai_jsonx.patch",
    revision = "a29fe22bc88c13040654d8381139d2433a5b76b6",
    deps = [
        ":jeffail_gabs",
    ],
)

go_get(
    name = "jeffail_gabs",
    get = "github.com/Jeffail/gabs",
    licences = ["MIT"],
    revision = "1f7d097f998e701866a053a116f3ff355e4214fd",
)

go_get(
    name = "keybase_go-crypto",
    get = "github.com/keybase/go-crypto/...",
    licences = ["BSD-3-Clause"],
    revision = "a054578053044161085030449d77f6ceee2cc31a",
    deps = [
        ":x_sys",
    ],
)

go_get(
    name = "mitchellh_copystructure",
    get = "github.com/mitchellh/copystructure",
    licences = ["MIT"],
    revision = "9a1b6f44e8da0e0e374624fb0a825a231b00c537",
    deps = [":mitchellh_reflectwalk"],
)

go_get(
    name = "mitchellh_reflectwalk",
    get = "github.com/mitchellh/reflectwalk",
    licences = ["MIT"],
    revision = "3e2c75dfad4fbf904b58782a80fd595c760ad185",
)

go_get(
    name = "patrickmn_go-cache",
    get = "github.com/patrickmn/go-cache",
    licences = ["MIT"],
    revision = "46f407853014144407b6c2ec7ccc76bf67958d93",
)

go_get(
    name = "duo",
    get = "github.com/duosecurity/duo_api_golang/...",
    licences = ["bsd-3-clause"],
    revision = "9b49acdeffee4049463e12a2249d0a05d54e599c",
)

go_get(
    name = "go-ldap",
    get = "github.com/go-ldap/ldap",
    licences = ["MIT"],
    revision = "v3.2.3",
    deps = [
        ":asn1-ber",
        ":ntlmssp",
    ],
)

go_get(
    name = "asn1-ber",
    get = "github.com/go-asn1-ber/asn1-ber",
    licences = ["MIT"],
    revision = "v1.5.1",
)

go_get(
    name = "ntlmssp",
    get = "github.com/Azure/go-ntlmssp",
    licences = ["MIT"],
    revision = "66371956d46c8e2133a2b72b3d320e435465011f",
    deps = [":x_crypto"],
)

go_get(
    name = "vault",
    get = "github.com/hashicorp/vault",
    install = [
        "api/...",
        "audit",
        "builtin/credential/ldap",
        "builtin/plugin",
        "builtin/logical/database/...",
        "builtin/logical/transit",
        "helper/base62",
        "helper/compressutil",
        "helper/consts",
        "helper/hclutil",
        "helper/identity",
        "helper/jsonutil",
        "helper/license",
        "helper/logging",
        "helper/forwarding",
        "helper/dbtxn",
        "helper/policyutil",
        "helper/errutil",
        "helper/certutil",
        "helper/mlock",
        "helper/cryptoutil",
        "helper/ldaputil",
        "helper/mfa/...",
        "helper/password",
        "helper/namespace",
        "helper/keysutil",
        "helper/kdf",
        "helper/locksutil",
        "helper/parseutil",
        "helper/pathmanager",
        "helper/pgpkeys",
        "helper/identity/...",
        "helper/pluginutil",
        "helper/reload",
        "helper/salt",
        "helper/strutil",
        "helper/tlsutil",
        "helper/storagepacker",
        "helper/wrapping",
        "helper/xor",
        "http",
        "logical",
        "logical/plugin",
        "logical/plugin/pb",
        "logical/framework",
        "physical",
        "physical/inmem",
        "plugins/database/postgresql",
        "plugins/helper/database/...",
        "plugins",
        "shamir",
        "vault",
        "vault/seal",
        "version",
    ],
    licences = ["MPL-2.0"],
    patch = "hashicorp_vault-testing.patch",
    revision = "85909e3373aa743c34a6a0ab59131f61fd9e8e43",
    strip = [
        "vendor",
        "plugins/database/mysql",
    ],
    deps = [
        ":armon_go-metrics",
        ":armon_go-radix",
        ":asn1-ber",
        ":cleanhttp",
        ":duo",
        ":elazarl_go-bindata-assetfs",
        ":go-ldap",
        ":go-multierror",
        ":go-rootcerts",
        ":grpc",
        ":gziphandler",
        ":hashicorp_go-immutable-radix",
        ":hashicorp_go-memdb",
        ":hashicorp_go-plugin",
        ":hashicorp_go-sockaddr",
        ":hashicorp_go-uuid",
        ":hcl",
        ":hclog",
        ":jefferai_jsonx",
        ":keybase_go-crypto",
        ":lz4",
        ":mapstructure",
        ":mitchellh_copystructure",
        ":patrickmn_go-cache",
        ":pester",
        ":pq",
        ":protobuf",
        ":retryablehttp",
        ":ryanuber_go-glob",
        ":sermoDigital_jose",
        ":snappy",
        ":structs",
        ":version",
        ":x_net",
        ":x_time_rate",
    ],
)

go_get(
    name = "hcl",
    get = "github.com/hashicorp/hcl/...",
    licences = ["mpl-2.0"],
    revision = "23c074d0eceb2b8a5bfdbb271ab780cde70f05a8",
)

go_get(
    name = "errwrap",
    get = "github.com/hashicorp/errwrap",
    licences = ["mpl-2.0"],
    revision = "v1.0.0",
)

go_get(
    name = "go-multierror",
    get = "github.com/hashicorp/go-multierror",
    licences = ["mpl-2.0"],
    revision = "v1.1.0",
    deps = [":errwrap"],
)

go_get(
    name = "go-homedir",
    get = "github.com/mitchellh/go-homedir",
    licences = ["mit"],
    revision = "b8bc1bf767474819792c23f32d8286a45736f1c6",
)

go_get(
    name = "go-rootcerts",
    get = "github.com/hashicorp/go-rootcerts",
    licences = ["mpl-2.0"],
    revision = "6bb64b370b90e7ef1fa532be9e591a81c3493e00",
    deps = [":go-homedir"],
)

go_get(
    name = "pester",
    get = "github.com/sethgrid/pester",
    licences = ["mit"],
    revision = "0af5bab1e1ea2860c5aef8e77427bab011d774d8",
)

go_get(
    name = "websocket",
    get = "github.com/gorilla/websocket",
    licences = ["bsd-2-clause"],
    revision = "v1.4.2",
)

go_get(
    name = "yaml.v3",
    get = "gopkg.in/yaml.v3",
    licences = ["apache-2.0"],
    revision = "v3",
)

go_get(
    name = "yaml.v2",
    get = "gopkg.in/yaml.v2",
    licences = ["apache-2.0"],
    revision = "v2.2.4",
)

go_get(
    name = "mapstructure",
    get = "github.com/mitchellh/mapstructure",
    licences = ["mit"],
    revision = "06020f85339e21b2478f756a78e295255ffa4d6a",
)

go_get(
    name = "structs",
    get = "github.com/fatih/structs",
    licences = ["mit"],
    revision = "f5faa72e73092639913f5833b75e1ac1d6bc7a63",
)

go_get(
    name = "sarama",
    get = "github.com/Shopify/sarama",
    licences = ["mit"],
    revision = "v1.19.0",
    # Only here temporarily until we migrate entirely to confluent-kafka.
    # TODO(simon): Remove once we are using only a single Go Kafka driver.
    # TODO(alessandro, dimitar): Remove payment order visibility when switching to confluent or
    # when a common kafka error wrapper is added.
    visibility = [
        "//common/go/kafka/...",
        "//vault/audit/audit_processor/...",
        "//vault/common/go/stream_api/...",
        "//vault/infrastructure/db_data_exporter/processors/...",
        "//vault/kernel/payment_order/common/errors/...",
        "//vault/kernel/postings_api/postings_processor/processor/workers/streamers/...",
    ],
    deps = [
        ":go-metrics",
        ":go-resiliency",
        ":go-spew",
        ":go-xerial-snappy",
        ":lz4",
        ":queue",
    ],
)

go_get(
    name = "go-xerial-snappy",
    get = "github.com/eapache/go-xerial-snappy",
    licences = ["mit"],
    revision = "bb955e01b9346ac19dc29eb16586c90ded99a98c",
    visibility = [],
    deps = [
        ":snappy",
    ],
)

go_get(
    name = "go-metrics",
    get = "github.com/rcrowley/go-metrics",
    licences = ["bsd-2-clause"],
    revision = "1f30fe9094a513ce4c700b9a54458bbb0c96996c",
    visibility = [],
    deps = [":stathat"],
)

go_get(
    name = "stathat",
    get = "github.com/stathat/go",
    licences = ["mit"],
    revision = "74669b9f388d9d788c97399a0824adbfee78400e",
    visibility = [],
)

go_get(
    name = "go-resiliency",
    get = "github.com/eapache/go-resiliency/...",
    licences = ["mit"],
    revision = "b1fe83b5b03f624450823b751b662259ffc6af70",
    visibility = [],
)

go_get(
    name = "lz4",
    get = "github.com/pierrec/lz4",
    licences = ["bsd-3-clause"],
    revision = "08c27939df1bd95e881e2c2367a749964ad1fceb",
    deps = [":xxHash"],
)

go_get(
    name = "xxHash",
    get = "github.com/pierrec/xxHash/...",
    licences = ["bsd-3-clause"],
    revision = "a0006b13c722f7f12368c00a3d3c2ae8a999a0c6",
    visibility = [],
)

go_get(
    name = "queue",
    get = "github.com/eapache/queue",
    licences = ["mit"],
    revision = "44cc805cf13205b55f69e14bcb69867d1ae92f98",
    visibility = [],
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
    name = "logfmt",
    get = "github.com/go-logfmt/logfmt",
    licences = ["mit"],
    revision = "390ab7935ee28ec6b286364bba9b4dd6410cb3d5",
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
    name = "thrift",
    get = "github.com/apache/thrift",
    install = ["lib/go/thrift"],
    licences = ["apache-2.0"],
    revision = "327ebb6c2b6df8bf075da02ef45a2a034e9b79ba",
)

go_get(
    name = "opentracing",
    get = "github.com/opentracing/opentracing-go/...",
    licences = ["apache-2.0"],
    revision = "1949ddbfd147afd4d964a9f00b24eb291e0e7c38",
    deps = [":x_net"],
)

go_get(
    name = "observer",
    get = "github.com/opentracing-contrib/go-observer",
    licences = ["apache-2.0"],
    revision = "a52f2342449246d5bcc273e65cbdcfa5f7d6c63c",
    deps = [":opentracing"],
)

go_get(
    name = "zipkin-opentracing",
    get = "github.com/openzipkin/zipkin-go-opentracing/...",
    licences = ["mit"],
    revision = "45e90b00710a4c34a1a7d8a78d90f9b010b0bd4d",
    deps = [
        ":gogo-protobuf",
        ":logfmt",
        ":observer",
        ":opentracing",
        ":sarama",
        ":thrift",
    ],
)

go_get(
    name = "ar",
    get = "github.com/blakesmith/ar",
    licences = ["mit"],
    revision = "8bd4349a67f2533b078dbc524689d15dba0f4659",
)

go_get(
    name = "lru",
    get = "github.com/hashicorp/golang-lru/...",
    licences = ["mpl-2.0"],
    revision = "0fb14efe8c47ae851c0034ed7a448854d3d34cf3",
)

go_get(
    name = "retryablehttp",
    get = "github.com/hashicorp/go-retryablehttp",
    licences = ["mpl-2.0"],
    revision = "v0.6.1",
    deps = [
        ":cleanhttp",
        ":hclog",
    ],
)

go_get(
    name = "hclog",
    get = "github.com/hashicorp/go-hclog",
    licences = ["mpl-2.0"],
    revision = "v0.9.2",
)

go_get(
    name = "logrus",
    get = "github.com/sirupsen/logrus",
    licences = ["mit"],
    revision = "v1.6.0",
    deps = [
        ":x_crypto",
        ":x_sys",
    ],
)

go_get(
    name = "gonduit",
    get = "github.com/thought-machine/gonduit/...",
    licences = ["mit"],
    revision = "33ee2e99d6bda153843b21ce242429211ae1a240",
    strip = ["test/server"],
    deps = [":typed"],
)

go_get(
    name = "typed",
    get = "github.com/karlseguin/typed",
    licences = ["mit"],
    revision = "2f542a6ff2d9962d9af4becdddf2e1a724a7d776",
)

go_get(
    name = "gomeme",
    get = "github.com/peterebden/gomeme",
    licences = ["mit"],
    revision = "f962c4b3e65c28d664f4ab86a766f2e38c752e9b",
    deps = [
        ":freetype",
        ":x_image",
    ],
)

go_get(
    name = "freetype",
    get = "github.com/golang/freetype/...",
    licences = ["freetype"],
    revision = "e2365dfdc4a05e4b8299a783240d4a7d5a65d4e4",
    deps = [
        ":x_image",
    ],
)

go_get(
    name = "x_image",
    get = "golang.org/x/image/...",
    licences = ["bsd-3-clause"],
    revision = "af66defab954cb421ca110193eed9477c8541e2a",
    deps = [":x_text"],
)

go_get(
    name = "survey",
    get = "gopkg.in/AlecAivazis/survey.v1",
    install = [
        "",
        "core",
        "terminal",
        "vendor/...",
    ],
    licences = ["mit"],
    repo = "github.com/AlecAivazis/survey",
    revision = "17861e192dc11fd2f5081df1932c94cce262fa1e",
)

go_get(
    name = "iso20022",
    get = "github.com/thought-machine/finance-messaging/...",
    licences = ["apache-2.0"],
    revision = "51a20b5a8bb10a65e75c3992b85f4d9f11314fbf",
    deps = [
        ":logging",
    ],
)

go_get(
    name = "blang-semver",
    get = "github.com/blang/semver/...",
    licences = ["mit"],
    revision = "2ee87856327ba09384cabd113bc6b5d174e9ec0f",
)

go_get(
    name = "postgres_exporter",
    binary = True,
    get = "github.com/wrouesnel/postgres_exporter/cmd/postgres_exporter",
    licences = ["apache-2.0"],
    revision = "v0.4.7",
    deps = [
        ":blang-semver",
        ":kingpin",
        ":pq",
        ":prometheus",
        ":prometheus_common",
        ":template",
        ":yaml.v2",
    ],
)

go_get(
    name = "mattn_go-colorable",
    get = "github.com/mattn/go-colorable",
    licences = ["MIT"],
    revision = "98ec13f34aabf44cc914c65a1cfb7b9bc815aef1",
    deps = [":mattn_go-isatty"],
)

go_get(
    name = "fatih_color",
    get = "github.com/fatih/color",
    licences = ["mit"],
    revision = "3f9d52f7176a6927daacff70a3e8d1dc2025c53e",
    strip = ["vendor"],
    deps = [
        ":mattn_go-colorable",
        ":mattn_go-isatty",
        ":x_sys",
    ],
)

go_get(
    name = "spf13_cobra",
    get = "github.com/spf13/cobra",
    licences = ["apache-2.0"],
    revision = "d2d81d9a96e23f0255397222bb0b4e3165e492dc",
    deps = [
        ":pflag",
    ],
)

go_get(
    name = "spf13_jwalterweatherman",
    get = "github.com/spf13/jwalterweatherman",
    licences = ["mit"],
    revision = "94f6ae3ed3bceceafa716478c5fbf8d29ca601a1",
)

go_get(
    name = "pelletier_go-toml",
    get = "github.com/pelletier/go-toml",
    licences = ["mit"],
    revision = "c01d1270ff3e442a8a57cddc1c92dc1138598194",
)

go_get(
    name = "fsnotify_fsnotify",
    get = "github.com/fsnotify/fsnotify",
    licences = ["bsd-3-clause"],
    revision = "c2828203cd70a50dcccfb2761f8b1f8ceef9a8e9",
    deps = [
        ":x_sys",
    ],
)

go_get(
    name = "magiconair_properties",
    get = "github.com/magiconair/properties",
    licences = ["bsd-2-clause"],
    revision = "c2353362d570a7bfa228149c62842019201cfb71",
)

go_get(
    name = "xeipuuv_gojsonpointer",
    get = "github.com/xeipuuv/gojsonpointer",
    licences = ["apache-2.0"],
    revision = "4e3ac2762d5f479393488629ee9370b50873b3a6",
)

go_get(
    name = "xeipuuv_gojsonreference",
    get = "github.com/xeipuuv/gojsonreference",
    licences = ["apache-2.0"],
    revision = "bd5ef7bd5415a7ac448318e64f11a24cd21e594b",
    deps = [
        ":xeipuuv_gojsonpointer",
    ],
)

go_get(
    name = "xeipuuv_gojsonschema",
    get = "github.com/xeipuuv/gojsonschema",
    licences = ["apache-2.0"],
    revision = "ac52e6811b56ee2b7730a88915e956afe2cc0d69",
    deps = [
        ":xeipuuv_gojsonreference",
    ],
)

go_get(
    name = "spf13_cast",
    get = "github.com/spf13/cast",
    licences = ["mit"],
    revision = "8c9545af88b134710ab1cd196795e7f2388358d7",
)

go_get(
    name = "spf13_afero",
    get = "github.com/spf13/afero/...",
    licences = ["apache-2.0"],
    revision = "d40851caa0d747393da1ffb28f7f9d8b4eeffebd",
    deps = [
        ":sftp",
        ":x_text",
    ],
)

go_get(
    name = "sftp",
    get = "github.com/pkg/sftp",
    licences = ["bsd-2-clause"],
    revision = "08de04f133f27844173471167014e1a753655ac8",
    deps = [
        ":errors",
        ":kr_fs",
        ":x_crypto",
    ],
)

go_get(
    name = "spf13_viper",
    get = "github.com/spf13/viper",
    licences = ["mit"],
    revision = "907c19d40d9a6c9bb55f040ff4ae45271a4754b9",
    deps = [
        ":fsnotify_fsnotify",
        ":hcl",
        ":magiconair_properties",
        ":mapstructure",
        ":pelletier_go-toml",
        ":pflag",
        ":spf13_afero",
        ":spf13_cast",
        ":spf13_jwalterweatherman",
        ":xeipuuv_gojsonschema",
        ":yaml.v2",
    ],
)

go_get(
    name = "kubeval",
    binary = True,
    get = "github.com/instrumenta/kubeval",
    licences = ["apache-2.0"],
    revision = "0.14.0",
    deps = [
        ":fatih_color",
        ":go-multierror",
        ":kubernetes_yaml",
        ":spf13_cobra",
        ":spf13_viper",
    ],
)

go_get(
    name = "kubernetes_yaml",
    get = "sigs.k8s.io/yaml",
    licences = ["mit"],
    revision = "v1.1.0",
    deps = [":yaml.v2"],
)

go_get(
    name = "yaml",
    get = "github.com/ghodss/yaml",
    licences = ["mit"],
    revision = "v1.0.0",
    deps = [":yaml.v2"],
)

go_get(
    name = "saml",
    get = "github.com/crewjam/saml",
    install = [
        "",
        "logger",
        "samlidp/...",
        "xmlenc/...",
    ],
    licences = ["bsd-2-clause"],
    revision = "ebc5f787b786ee76ee69bf49184fade38d2238af",
    deps = [
        ":check.v1",
        ":clockwork",
        ":goji",
        ":goxmldsig",
        ":jwt-go",
        ":uniuri",
        ":x_crypto",
    ],
)

go_get(
    name = "etree",
    get = "github.com/beevik/etree",
    licences = ["bsd-2-clause"],
    revision = "9d7e8feddccb4ed1b8afb54e368bd323d2ff652c",
)

go_get(
    name = "uniuri",
    get = "github.com/dchest/uniuri",
    licences = ["CC0 1.0 Universal"],
    revision = "8902c56451e9b58ff940bbe5fec35d5f9c04584a",
)

go_get(
    name = "clockwork",
    get = "github.com/jonboulle/clockwork",
    licences = ["apache-2.0"],
    revision = "2eee05ed794112d45db504eb05aa693efd2b8b09",
)

go_get(
    name = "pretty",
    get = "github.com/kr/pretty",
    licences = ["mit"],
    revision = "73f6ac0b30a98e433b289500d779f50c1a6f0712",
    deps = [
        ":kr_text",
    ],
)

go_get(
    name = "kr_text",
    get = "github.com/kr/text",
    licences = ["mit"],
    revision = "e2ffdb16a802fe2bb95e2e35ff34f0e53aeef34f",
)

go_get(
    name = "kr_fs",
    get = "github.com/kr/fs",
    licences = ["bsd-3-clause"],
    revision = "1455def202f6e05b95cc7bfc7e8ae67ae5141eba",
)

go_get(
    name = "goxmldsig",
    get = "github.com/russellhaering/goxmldsig",
    install = [
        "",
        "etreeutils",
        "types",
    ],
    licences = ["apache-2.0"],
    revision = "7acd5e4a6ef74fe1b082c20f119556adf70c3944",
    deps = [
        ":clockwork",
        ":etree",
    ],
)

go_get(
    name = "goji",
    get = "github.com/zenazn/goji",
    install = [
        "",
        "bind",
        "graceful",
        "graceful/listener",
        "web",
        "web/middleware",
        "web/mutil",
    ],
    licences = ["MIT"],
    revision = "64eb34159fe53473206c2b3e70fe396a639452f2",
)

go_get(
    name = "check.v1",
    get = "gopkg.in/check.v1",
    licences = ["bsd-2-clause"],
    revision = "788fd78401277ebd861206a03c884797c6ec5541",
    deps = [
        ":pretty",
    ],
)

go_get(
    name = "ldap",
    get = "gopkg.in/ldap.v3",
    licences = ["mit"],
    repo = "github.com/go-ldap/ldap",
    revision = "v3",
    deps = [":asn1ber"],
)

go_get(
    name = "asn1ber",
    get = "gopkg.in/asn1-ber.v1",
    licences = ["mit"],
    repo = "github.com/go-asn1-ber/asn1-ber",
    revision = "v1.3",
)

go_get(
    name = "copier",
    get = "github.com/jinzhu/copier",
    licences = ["mit"],
    revision = "7e38e58719c33e0d44d585c4ab477a30f8cb82dd",
)

go_get(
    name = "ojcm_copier",
    get = "github.com/ojcm/copier",
    licences = ["mit"],
    revision = "06067470b4cd2df6027c0bebcfd8f7c0762ee99e",
)

go_get(
    name = "deepcopy",
    get = "github.com/mohae/deepcopy",
    licences = ["mit"],
    revision = "c48cc78d482608239f6c4c92a4abd87eb8761c90",
)

go_get(
    name = "squirrel",
    get = "github.com/Masterminds/squirrel",
    licences = ["MIT"],
    revision = "12f2162c8d8d31528dfd432bafd3acaaf793768e",
    deps = [
        ":lann_builder",
    ],
)

go_get(
    name = "lann_builder",
    get = "github.com/lann/builder",
    licences = ["mit"],
    revision = "47ae307949d02aa1f1069fdafc00ca08e1dbabac",
    deps = [
        ":lann_ps",
    ],
)

go_get(
    name = "lann_ps",
    get = "github.com/lann/ps",
    licences = ["mit"],
    revision = "62de8c46ede02a7675c4c79c84883eb164cb71e3",
)

go_get(
    name = "mockery",
    binary = True,
    get = "github.com/vektra/mockery/...",
    licences = ["bsd-3-clause"],
    revision = "e78b021dcbb558a8e7ac1fc5bc757ad7c277bb81",
    deps = [
        ":testify",
        ":x_tools",
    ],
)

go_get(
    name = "libxml2",
    get = "github.com/lestrrat-go/libxml2/...",
    licences = ["mit"],
    revision = "8ad9bf187c39812bd6083ccc1261a713f41e0c17",
    deps = [
        ":errors",
    ],
)

CLOUD_GOOGLE_GO_VERSION = "v0.46.3"

go_get(
    name = "cloud-storage",
    get = "cloud.google.com/go/storage",
    licences = ["apache-2.0"],
    revision = CLOUD_GOOGLE_GO_VERSION,
    deps = [
        ":cloud-iam",
        ":cloud-internal",
    ],
)

go_get(
    name = "cloud-iam",
    get = "cloud.google.com/go/iam",
    licences = ["apache-2.0"],
    revision = CLOUD_GOOGLE_GO_VERSION,
    deps = [
        ":api",
        ":cloud-metadata",
        ":gax-go",
        ":grpc",
        ":iam",
    ],
)

go_get(
    name = "cloud-internal",
    get = "cloud.google.com/go/internal",
    install = [
        "",
        "optional",
        "trace",
        "version",
    ],
    licences = ["apache-2.0"],
    revision = CLOUD_GOOGLE_GO_VERSION,
    deps = [
        ":api",
        ":gax-go",
        ":go.opencensus.io",
    ],
)

go_get(
    name = "cloud-metadata",
    get = "cloud.google.com/go/compute/metadata",
    licences = ["apache-2.0"],
    revision = CLOUD_GOOGLE_GO_VERSION,
)

go_get(
    name = "cloud-containeranalysis",
    get = "cloud.google.com/go/containeranalysis/apiv1",
    licences = ["apache-2.0"],
    revision = CLOUD_GOOGLE_GO_VERSION,
    deps = [
        ":api",
        ":cloud-grafeas",
        ":gax-go",
        ":genproto_devtools_containeranalysis",
        ":grpc",
        ":iam",
    ],
)

go_get(
    name = "cloud-grafeas",
    get = "cloud.google.com/go/grafeas/apiv1",
    licences = ["apache-2.0"],
    revision = CLOUD_GOOGLE_GO_VERSION,
    deps = [
        ":api",
        ":gax-go",
        ":genproto_devtools_containeranalysis",
        ":genproto_grafeas",
    ],
)

go_get(
    name = "api",
    get = "google.golang.org/api",
    install = [
        "calendar/v3",
        "googleapi",
        "googleapi/transport",
        "internal",
        "internal/gensupport",
        "internal/third_party/uritemplates",
        "iterator",
        "option",
        "storage/v1",
        "sheets/v4",
        "support/bundler",
        "transport",
        "transport/grpc",
        "transport/http",
        "transport/http/internal/propagation",
        "dataflow/...",
        "compute/v1",
        "container/v1",
        "drive/v3",
    ],
    licences = ["bsd-3-clause"],
    repo = "https://github.com/googleapis/google-api-go-client",
    revision = "v0.17.0",
    deps = [
        ":gax-go",
        ":go.opencensus.io",
        ":grpc",
        ":oauth2",
        ":x_sync",
    ],
)

go_get(
    name = "gax-go",
    get = "github.com/googleapis/gax-go",
    install = [
        "",
        "v2",
    ],
    licences = ["bsd-3-clause"],
    revision = "4f559a1e5fd0a1cac6a74c886f1c87917e23ac94",
    deps = [
        ":grpc",
    ],
)

go_get(
    name = "go.opencensus.io",
    get = "go.opencensus.io/...",
    licences = ["apache-2.0"],
    repo = "github.com/census-instrumentation/opencensus-go",
    revision = "v0.22.1",
    deps = [
        ":cmp",
        ":groupcache",
        ":grpc",
        ":protobuf",
        ":x_net",
    ],
)

go_get(
    name = "oauth2",
    get = "golang.org/x/oauth2",
    install = [
        "",
        "google",
        "internal",
        "jws",
        "jwt",
    ],
    licences = ["bsd-3-clause"],
    revision = "d668ce993890a79bda886613ee587a69dd5da7a6",
    deps = [
        ":cloud-metadata",
        ":x_net",
    ],
)

go_get(
    name = "humanize",
    get = "github.com/dustin/go-humanize",
    licences = ["MIT"],
    revision = "v1.0.0",
)

go_get(
    name = "gocloud",
    get = "gocloud.dev",
    install = [
        # This is an awkward lot but necessary so we don't have to pull in transitive
        # dependencies for the providers for aws, azure etc.
        "aws",
        "blob",
        "blob/driver",
        "blob/fileblob",
        "blob/gcsblob",
        "blob/memblob",
        "blob/s3blob",
        "gcerrors",
        "gcp",
        "internal/escape",
        "internal/gcerr",
        "internal/oc",
        "internal/openurl",
        "internal/retry",
        "internal/useragent",
    ],
    licences = ["apache-2.0"],
    repo = "github.com/google/go-cloud",
    revision = "v0.17.0",
    deps = [
        ":aws-sdk-go",
        ":cloud-storage",
        ":gax-go",
        ":go.opencensus.io",
        ":grpc",
        ":oauth2",
        ":wire",
        ":x_errors",
    ],
)

go_get(
    name = "x_errors",
    get = "golang.org/x/xerrors/...",
    licences = ["bsd-3-clause"],
    revision = "a5947ffaace3e882f334c1750858b4a6a7e52422",
    deps = [":x_sys"],
)

go_get(
    name = "wire",
    get = "github.com/google/wire",
    licences = ["apache-2.0"],
    revision = "v0.3.0",
)

go_get(
    name = "go-junit",
    get = "github.com/goddenrich/go-junit",
    licences = ["mit"],
    revision = "918c0668713dc2f5b9294092beede2ccbad15a33",
)

# TODO(mike): Replace the ones in third_party/go/kubernetes with these
# MIT Licence
go_get(
    name = "ginkgo",
    get = "github.com/onsi/ginkgo/...",
    licences = ["mit"],
    revision = "v1.8.0",
    deps = [
        ":hpcloud-tail",
    ],
)

# MIT Licence
go_get(
    name = "gomega",
    get = "github.com/onsi/gomega/...",
    licences = ["mit"],
    revision = "v1.5.0",
    deps = [
        ":protobuf",
        ":x_net",
        ":yaml.v2",
    ],
)

go_get(
    name = "hpcloud-tail",
    get = "github.com/hpcloud/tail/...",
    install = [
        "",
        "ratelimiter/...",
        "util/...",
        "vendor/...",
        "watch/...",
    ],
    licences = ["MIT"],
    revision = "a1dbeea552b7c8df4b542c66073e393de198a800",
)

go_get(
    name = "go-elasticsearch",
    get = "github.com/elastic/go-elasticsearch/...",
    install = [
        "",
        "esapi/...",
        "estransport/...",
        "internal/version",
    ],
    licences = ["apache-2.0"],
    revision = "6eca2ffff2c65f85c21f57ca570c1525a4b90f24",
    deps = [
        ":chroma",
        ":spf13_cobra",
        ":x_tools",
        ":yaml.v2",
    ],
)

go_get(
    name = "chroma",
    get = "github.com/alecthomas/chroma",
    install = [
        "",
        "formatters/...",
        "lexers/...",
        "styles/...",
    ],
    licences = ["mit"],
    revision = "9618112850d4f63a51b24dd8b3e2c6b0b819e6db",
    deps = [":regexp2"],
)

go_get(
    name = "regexp2",
    get = "github.com/dlclark/regexp2/...",
    licences = ["mit"],
    revision = "7632a260cbaf5e7594fc1544a503456ecd0827f1",
    deps = [":fnmatch"],
)

go_get(
    name = "fnmatch",
    get = "github.com/danwakefield/fnmatch",
    licences = ["bsd-2-clause"],
    revision = "cbb64ac3d964b81592e64f957ad53df015803288",
)

go_get(
    name = "sanitize",
    get = "github.com/kennygrant/sanitize",
    licences = ["bsd-3-clause"],
    revision = "06ec0d0dbcd497d01e5189b16f5263f712e61a8e",
    deps = [":x_net"],
)

go_get(
    name = "gucci",
    binary = True,
    get = "github.com/noqcks/gucci",
    licences = ["mit"],
    revision = "1.2.1",
    deps = [
        ":errors",
        ":sprig",
        ":yaml.v2",
    ],
)

go_get(
    name = "sprig",
    get = "github.com/Masterminds/sprig",
    licences = ["MIT"],
    revision = "2.19.0",
    deps = [
        ":google-uuid",
        ":goutils",
        ":mergo",
        ":sprig-semver",
        ":urfave-cli",
        ":x_crypto",
        ":xstrings",
    ],
)

go_get(
    name = "google-uuid",
    get = "github.com/google/uuid",
    licences = ["bsd-3-clause"],
    revision = "v1.1.1",
)

go_get(
    name = "goutils",
    get = "github.com/Masterminds/goutils",
    licences = ["apache-2.0"],
    revision = "v1.1.0",
)

go_get(
    name = "sprig-semver",
    get = "github.com/Masterminds/semver",
    licences = ["mit"],
    revision = "v1.4.2",
)

go_get(
    name = "xstrings",
    get = "github.com/huandu/xstrings",
    licences = ["mit"],
    revision = "v1.2.0",
)

go_get(
    name = "mergo",
    get = "github.com/imdario/mergo",
    licences = ["bsd-3-clause"],
    revision = "v0.3.7",
)

go_get(
    name = "urfave-cli",
    get = "github.com/urfave/cli",
    licences = ["mit"],
    revision = "v1.19.1",
)

go_get(
    name = "golang-set",
    get = "github.com/deckarep/golang-set",
    licences = ["mit"],
    revision = "v1.7.1",
)

go_get(
    name = "strcase",
    get = "github.com/iancoleman/strcase",
    licences = ["MIT"],
    revision = "e506e3ef73653e84c592ba44aab577a46678f68c",
)

go_get(
    name = "gomemcache",
    get = "github.com/bradfitz/gomemcache",
    hashes = [
        "4494425a45378843d26f3eb6715bad9f86866fcc",
    ],
    install = [
        "memcache",
    ],
    licences = ["apache-2.0"],
    revision = "1952afaa557dc08e8e0d89eafab110fb501c1a2b",
)

go_get(
    name = "go-junit-report",
    binary = True,
    get = "github.com/jstemmer/go-junit-report/...",
    hashes = [
        "b0957697366b1133c98da47489ec400a5b621c5f",
    ],
    licences = ["mit"],
    revision = "af01ea7f8024089b458d804d5cdf190f962a9a0c",
)

go_get(
    name = "x_time_rate",
    get = "golang.org/x/time/rate",
    licences = ["bsd-3-clause"],
    revision = "f51c12702a4d776e4c1fa9b0fabab841babae631",
    deps = [
        ":x_net",
    ],
)

go_get(
    name = "stripe",
    get = "github.com/stripe/stripe-go/...",
    licences = ["MIT"],
    revision = "v61.10.0",
    deps = [
        ":x_net",
    ],
)

go_get(
    name = "stripe-mock",
    binary = True,
    get = "github.com/stripe/stripe-mock/...",
    licences = ["MIT"],
    deps = [
        ":x_net",
    ],
)

go_get(
    name = "json-iterator",
    get = "github.com/json-iterator/go/...",
    licences = ["MIT"],
    revision = "v1.1.6",
    deps = [
        ":concurrent",
        ":reflect2",
    ],
)

go_get(
    name = "concurrent",
    get = "github.com/modern-go/concurrent/...",
    licences = ["apache-2.0"],
    revision = "1.0.3",
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
    name = "gomega_unmarshalled_matchers",
    get = "github.com/benjamintf1/unmarshalledmatchers/...",
    licences = ["mit"],
    revision = "f49d1795d752929500b8ed2289f1e21f10414667",
    deps = [
        ":gomega",
    ],
)

go_get(
    name = "containerregistry",
    get = "github.com/google/go-containerregistry",
    hashes = [
        "e5a123ec6aada59a3f69c3cad165171d1b563db8",
    ],
    install = [
        "pkg/...",
    ],
    licences = ["apache-2.0"],
    patch = "containerregistry_no_azure.patch",
    revision = "65ccf86b63450797094ceee6b3d1ddd79253cd89",
    strip = [
        "pkg/crane",
        "pkg/gcrane",
        "pkg/v1/daemon",
        "vendor",
        "cmd",
    ],
    deps = [
        ":cmp",
        ":errors",
        ":oauth2",
        ":spf13_cobra",
        ":x_sync",
        "//third_party/go/kubernetes",
        "//third_party/go/kubernetes:client-go",
    ],
)

go_get(
    name = "cmp",
    get = "github.com/google/go-cmp/cmp/...",
    hashes = [
        "135c3aa24dca4861b78521c267b3dc4d06087d93",
    ],
    licences = ["bsd-2-clause"],
    revision = "1b316004397f1f336546ca058ddb5b95c41a8772",
)

go_get(
    name = "docker",
    get = "github.com/docker/docker",
    hashes = [
        "a324f5512b8617c705b61bfd319497d86f52b4ef",
    ],
    install = [
        "api/types/...",
        "client",
    ],
    licences = ["apache-2.0"],
    repo = "github.com/moby/moby",
    revision = "v17.03.2-ce",
    strip = [
        "vendor",
    ],
    deps = [
        ":docker-distribution",
        ":docker-go-connections",
        ":docker-go-units",
        ":x_time_rate",
    ],
)

go_get(
    name = "docker-go-connections",
    get = "github.com/docker/go-connections",
    hashes = [
        "7fc5a6953b0d713230e681ba59bf0d64b5c0a134",
    ],
    licences = ["apache-2.0"],
    revision = "v0.4.0",
    deps = [
        ":errors",
    ],
)

go_get(
    name = "docker-go-units",
    get = "github.com/docker/go-units",
    hashes = [
        "5f3f5bf6076fe1eb23e175304c9e289a8abf102e",
    ],
    licences = ["apache-2.0"],
    revision = "v0.4.0",
)

go_get(
    name = "docker-distribution",
    get = "github.com/docker/distribution",
    hashes = [
        "0685444c9d846d1a3b6a040fd4431e29d3cadf86",
    ],
    install = [
        "",
        "reference",
        "digestset",
    ],
    licences = ["apache-2.0"],
    revision = "v2.7.1",
    strip = [
        "vendor",
    ],
    deps = [
        ":opencontainers-godigest",
        ":opencontainers-imagespec",
    ],
)

go_get(
    name = "opencontainers-godigest",
    get = "github.com/opencontainers/go-digest",
    hashes = [
        "55c3804f4c359d8be9658fdea91906ff56ac515a",
    ],
    install = [
        "",
    ],
    licences = ["apache-2.0"],
    revision = "ac19fd6e7483ff933754af248d80be865e543d22",
)

go_get(
    name = "opencontainers-imagespec",
    get = "github.com/opencontainers/image-spec",
    hashes = [
        "6c467a33d0b6bc840186ab182f013fbda705308e",
    ],
    install = [
        "specs-go",
        "specs-go/v1",
    ],
    licences = ["apache-2.0"],
    revision = "9f4348abedbe4415e6db1f08689fa7588045d982",
    deps = [
        ":opencontainers-godigest",
    ],
)

go_get(
    name = "tcpinfo",
    get = "github.com/brucespang/go-tcpinfo",
    licences = ["mit"],
    revision = "e6cc7410d081080ec38065c00bb3b72f3b144913",
)

go_get(
    name = "olekukonko-tablewriter",
    get = "github.com/olekukonko/tablewriter",
    licences = ["mit"],
    revision = "cc27d85e17cec9768d2ac401ea5d619a9628f16d",
    deps = [
        ":mattn-go-runewidth",
    ],
)

go_get(
    name = "mattn-go-runewidth",
    get = "github.com/mattn/go-runewidth",
    licences = ["mit"],
    revision = "703b5e6b11ae25aeb2af9ebb5d5fdf8fa2575211",
)

go_get(
    name = "gosec",
    binary = True,
    get = "github.com/securego/gosec",
    install = ["github.com/securego/gosec/cmd/..."],
    licences = ["apache-2.0"],
    revision = "29cec138dcc94f347e6d41550a5223571dc7d6cf",
    deps = [
        ":tls-observatory-constants",
        ":x_tools",
        ":yaml.v2",
        ":zxcvbn-go",
    ],
)

go_get(
    name = "zxcvbn-go",
    get = "github.com/nbutton23/zxcvbn-go",
    licences = ["mit"],
    revision = "eafdab6b0663b4b528c35975c8b0e78be6e25261",
)

go_get(
    name = "tls-observatory-constants",
    get = "github.com/mozilla/tls-observatory/constants",
    licences = ["MPL-2.0"],
    revision = "1.3.12",
)

go_get(
    name = "slack",
    get = "github.com/slack-go/slack/...",
    install = [
        "",
        "internal/...",
        "slackevents",
        "slackutilsx",
        "vendor/...",
    ],
    licences = ["bsd-2-clause"],
    revision = "v0.6.5",
)

go_get(
    name = "contrib-go-opencensus-io-exporter-prometheus",
    get = "contrib.go.opencensus.io/exporter/prometheus",
    licences = ["apache-2.0"],
    repo = "github.com/census-ecosystem/opencensus-go-exporter-prometheus",
    revision = "1d6bc2a7914046e6146ede481471d2801ba34121",
    deps = [
        ":go.opencensus.io",
        ":prometheus",
    ],
)

go_get(
    name = "groupcache",
    get = "github.com/golang/groupcache",
    install = [
        "lru",
    ],
    licences = ["apache-2.0"],
    revision = "404acd9df4cc9859d64fb9eed42e5c026187287a",
)

go_get(
    name = "gofpdf",
    get = "github.com/jung-kurt/gofpdf/...",
    licences = ["MIT"],
    revision = "v12.13.5",
)

go_get(
    name = "tz",
    get = "4d63.com/tz",
    licences = ["MIT"],
    revision = "v1.1.0",
)

go_get(
    name = "roaring",
    get = "github.com/RoaringBitmap/roaring",
    licences = ["apache-2.0"],
    revision = "ef1f54b597d2447c90b6801915da89e47ba8093d",
    deps = [
        ":go-unsnap-stream",
        ":msgp",
    ],
)

go_get(
    name = "bleve",
    get = "github.com/blevesearch/bleve",
    install = [
        "",  # Not sure why this is needed...
        "analysis",
        "analysis/analyzer/standard",
        "analysis/datetime/flexible",
        "analysis/datetime/optional",
        "analysis/lang/en",
        "analysis/token/lowercase",
        "analysis/token/porter",
        "analysis/token/stop",
        "analysis/tokenizer/unicode",
        "document",
        "geo",
        "index",
        "index/scorch",
        "index/scorch/mergeplan",
        "index/scorch/segment",
        "index/scorch/segment/zap",
        "index/store",
        "index/store/boltdb",
        "index/store/gtreap",
        "index/upsidedown",
        "mapping",
        "numeric",
        "registry",
        "search",
        "search/collector",
        "search/facet",
        "search/highlight",
        "search/highlight/format/html",
        "search/highlight/fragmenter/simple",
        "search/highlight/highlighter/html",
        "search/highlight/highlighter/simple",
        "search/query",
        "search/scorer",
        "search/searcher",
        "size",
    ],
    licences = ["apache-2.0"],
    revision = "e19c8bdb428bbabda3c3c74aaabb2d44dbce74ff",
    deps = [
        ":bbolt",
        ":go-porterstemmer",
        ":gtreap",
        ":mmap-go",
        ":protobuf",
        ":roaring",
        ":segment",
        ":snappy",
        ":vellum",
    ],
)

go_get(
    name = "go-porterstemmer",
    get = "github.com/blevesearch/go-porterstemmer",
    licences = ["MIT"],
    revision = "23a2c8e5cf1f380f27722c6d2ae8896431dc7d0e",
)

go_get(
    name = "segment",
    get = "github.com/blevesearch/segment",
    licences = ["apache-2.0"],
    revision = "762005e7a34fd909a84586299f1dd457371d36ee",
)

go_get(
    name = "vellum",
    get = "github.com/couchbase/vellum",
    install = [
        "",
        "levenshtein",
        "regexp",
        "utf8",
        "vendor/github.com/edsrzf/mmap-go",
        "vendor/github.com/willf/bitset",
    ],
    licences = ["apache-2.0"],
    revision = "ef2e028c01fdb60c46da4067d2e83745b8d54120",
)

go_get(
    name = "mmap-go",
    get = "github.com/edsrzf/mmap-go",
    licences = ["bsd-3-clause"],
    revision = "904c4ced31cdffe19e971afa0b3d319ff06d9c72",
    deps = [
        ":x_sys",
    ],
)

go_get(
    name = "bbolt",
    get = "github.com/etcd-io/bbolt",
    licences = ["MIT"],
    revision = "6e135e5d7e3d06257109f1458f8336a05ada6718",
)

go_get(
    name = "go-unsnap-stream",
    get = "github.com/glycerine/go-unsnap-stream",
    licences = ["MIT"],
    revision = "81cf024a9e0a906651886e23cb984a852a30b622",
    deps = [
        ":snappy",
    ],
)

go_get(
    name = "fwd",
    get = "github.com/philhofer/fwd",
    licences = ["MIT"],
    revision = "bb6d471dc95d4fe11e432687f8b70ff496cf3136",
)

go_get(
    name = "gtreap",
    get = "github.com/steveyen/gtreap",
    licences = ["MIT"],
    revision = "0abe01ef9be25c4aedc174758ec2d917314d6d70",
)

# The license on this is non-standard. It includes a provision to include
# the copyright notice.
# https://github.com/tinylib/msgp/blob/master/LICENSE
# I think it is fine to include this for now, but we probably want to get
# a review of this license (and to ensure we can be complaint with it).
go_get(
    name = "msgp",
    get = "github.com/tinylib/msgp/msgp",
    licences = ["MIT"],
    revision = "490d90d0e69181d9cfedd28343d876cd361adcc3",
    deps = [
        ":fwd",
    ],
)

go_get(
    name = "goquery",
    get = "github.com/PuerkitoBio/goquery",
    licences = ["bsd-3-clause"],
    revision = "v1.5.0",
    deps = [
        ":cascadia",
    ],
)

go_get(
    name = "cascadia",
    get = "github.com/andybalholm/cascadia",
    licences = ["bsd-2-clause"],
    revision = "v1.0.0",
    deps = [
        ":x_net",
    ],
)

go_get(
    name = "opa",
    get = "github.com/open-policy-agent/opa/...",
    licences = ["apache-2.0"],
    revision = "62bb63dd3b9ef98412dd2348c64b21382682fa2b",
    strip = [
        "vendor/github.com/ghodss/yaml",
        "vendor/github.com/pkg/errors",
        "vendor/github.com/OneOfOne/xxhash",
        "vendor/github.com/rcrowley/go-metrics",
        "vendor/github.com/gobwas/glob",
        "vendor/github.com/yashtewari/glob-intersection",
    ],
    deps = [
        ":errors",
        ":glob-intersection",
        ":go-metrics",
        ":gobwas-glob",
        ":oneOfOne-xxhash",
        ":yaml",
    ],
)

go_get(
    name = "oneOfOne-xxhash",
    get = "github.com/OneOfOne/xxhash",
    licences = ["apache-2.0"],
    revision = "74ace4fe5525ef62ce28d5093d6b0faaa6a575f3",
)

go_get(
    name = "gobwas-glob",
    get = "github.com/gobwas/glob/...",
    licences = ["MIT"],
    revision = "5ccd90ef52e1e632236f7326478d4faa74f99438",
)

go_get(
    name = "glob-intersection",
    get = "github.com/yashtewari/glob-intersection",
    licences = ["apache-2.0"],
    revision = "5c77d914dd0ba7bedca923f97232d37137e038f3",
    deps = [
        ":errors",
    ],
)

go_get(
    name = "godbus",
    get = "github.com/godbus/dbus",
    licences = ["BSD"],
    revision = "v5.0.3",
)

go_get(
    name = "libsecret",
    get = "github.com/gsterjov/go-libsecret",
    licences = ["MIT"],
    revision = "a6f4afe4910cad8688db3e0e9b9ac92ad22d54e1",
    deps = [":godbus"],
)

go_get(
    name = "protoreflect",
    get = "github.com/jhump/protoreflect/...",
    install = [
        "/desc",
        "/desc/protoparse",
        "/internal/",
        "/desc/internal",
        "/dynamic",
        "/codec",
    ],
    licences = ["apache-2.0"],
    revision = "43f69034e50eb4a370349b6b7194ef0c7b33d5ed",
    deps = [
        ":genproto_protobuf",
    ],
)

go_get(
    name = "go-set",
    get = "github.com/scylladb/go-set/...",
    licences = ["apache-2.0"],
    revision = "e560bb8f49bb7f34d4f59b7e771f6e1307c329da",
)

go_get(
    name = "h2conn",
    get = "github.com/posener/h2conn",
    revision = "13e7df33ed1548e5902d1b32bd6b5906cce194f5",
    deps = [":x_net"],
)

go_get(
    name = "github",
    get = "github.com/google/go-github/github",
    licences = [
        "BSD-3-Clause",
    ],
    revision = "v29.0.3",
    deps = [
        ":go-querystring",
        ":oauth2",
        ":x_crypto",
    ],
)

go_get(
    name = "go-querystring",
    get = "github.com/google/go-querystring/query",
    licences = [
        "BSD-3-Clause",
    ],
    revision = "v1.0.0",
)

go_get(
    name = "tgo",
    get = "github.com/trivago/tgo",
    install = [
        "tcontainer",
        "treflect",
    ],
    licences = ["BSD 3-Clause"],
    revision = "v1.0.0",
)

go_get(
    name = "go-jira",
    get = "github.com/andygrunwald/go-jira",
    licences = ["MIT"],
    revision = "v1.12.0",
    deps = [
        ":errors",
        ":go-querystring",
        ":jwt-go",
        ":structs",
        ":tgo",
    ],
)

go_get(
    name = "bender",
    get = "github.com/pinterest/bender",
    install = [
        "",
        "hist",
        "http",
    ],
    licences = ["apache-2.0"],
    revision = "b6161870dde9daebf81ad04aa585bf0fe398811c",
)

go_get(
    name = "uilive",
    get = "github.com/gosuri/uilive",
    licences = ["MIT"],
    revision = "v0.0.4",
)

go_get(
    name = "uiprogress",
    get = "github.com/gosuri/uiprogress",
    install = [
        "util/strutil",
        "",
    ],
    licences = ["MIT"],
    revision = "v0.0.1",
    deps = [
        ":uilive",
    ],
)

go_get(
    name = "go-spin",
    get = "github.com/tj/go-spin",
    licences = ["MIT"],
    revision = "v1.1.0",
)

go_get(
    name = "fbender",
    get = "github.com/facebookincubator/fbender",
    install = [
        "cmd/core",
        "cmd/core/runner",
        "cmd/core/options",
        "flags",
        "log",
        "metric",
        "tester",
        "tester/http",
        "tester/run",
        "recorders",
        "utils",
    ],
    licences = ["BSD"],
    revision = "3881cdb104d3f77c2d8397678d192b6fe1a3aecd",
    deps = [
        ":bender",
        ":logrus",
        ":spf13_cobra",
        ":pflag",
        # Could not manage to get rid of these even with
        # just tester in install
        ":uiprogress",
        ":go-spin",
    ],
)

go_get(
    name = "terraform",
    get = "github.com/hashicorp/terraform",
    hashes = ["d51f517738fd97356e8396980566bcfeb1dc3111"],
    install = [
        "addrs",
        "backend",
        "communicator",
        "config",
        "configs/...",
        "dag",
        "digraph",
        "experiments",
        "flatmap",
        "helper/...",
        "httpclient",
        "internal/...",
        "lang/...",
        "moduledeps",
        "plans/...",
        "plugin/discovery",
        "providers",
        "provisioners",
        "registry/...",
        "repl",
        "state",
        "states/...",
        "terraform",
        "tfdiags",
        "vendor/cloud.google.com/...",
        "vendor/github.com/...",
        "vendor/go.opencensus.io/...",
        "vendor/google.golang.org/...",
        "vendor/gopkg.in/...",

        # we'll not take x/net
        "vendor/golang.org/x/oauth2/...",
        "vendor/golang.org/x/text/...",
        "vendor/golang.org/x/crypto/...",
        "vendor/golang.org/x/mod/...",
        "vendor/golang.org/x/sys/...",
        "vendor/golang.org/x/time/...",
        "vendor/golang.org/x/tools/...",
        "version",
    ],
    licences = [
        "Mozilla Public License 2.0 (MPL 2.0)",
        "MIT",
        "Apache2",
    ],
    revision = "33ff1c03bb960b332be3af2e333462dde88b279e",  # v0.12.24 March 19th 2020
    strip = [
        "vendor/golang.org/x/net",
    ],
    deps = [
        # don't try to use vendor/golang.org/x/net, Cthulhu lives here, turn back now
        ":x_net",
    ],
)

go_get(
    name = "starlark-go",
    get = "go.starlark.net/...",
    install = [
        "starlark",
        "internal/compile",
        "internal/spell",
        "resolve",
        "starlarkstruct",
        "starlarktest",
        "syntax",
    ],
    licences = ["bsd-3-clause"],
    revision = "be5394c419b6941c14d194a98925a71512f8f979",
)

go_get(
    name = "promtool",
    binary = True,
    get = "github.com/prometheus/prometheus",
    install = ["cmd/..."],
    licences = ["apache-2.0"],
    revision = "v2.19.1",
    deps = [":prometheus"],
)

go_get(
    name = "alertmanager",
    get = "github.com/prometheus/alertmanager/...",
    install = [
        "vendor/...",
        "...",
    ],
    licences = ["apache-2.0"],
    revision = "v0.21.0",
)

go_get(
    name = "amtool",
    binary = True,
    get = [],
    install = ["github.com/prometheus/alertmanager/cmd/amtool"],
    licences = ["apache-2.0"],
    deps = [
        ":alertmanager",
    ],
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
    name="backoff",
    get = "github.com/cenkalti/backoff",
    licences = ["mit"],
    deps = [
        ":go-spew",
        ":difflib",
        ":x_sys",
        ":testify",
    ],
)
