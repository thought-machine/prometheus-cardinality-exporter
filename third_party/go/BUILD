package(default_visibility = ["PUBLIC"])

subinclude("///third_party/subrepos/pleasings//remote")

#########################
# Packages referenced by go files in repo
#########################

go_get(
    name = "go-flags",
    get = "github.com/jessevdk/go-flags",
    licences = ["bsd-3-clause"],
    revision = "v1.4.0",
)

go_get(
    name = "prometheus",
    get = "github.com/prometheus/client_golang/...",
    licences = ["apache-2.0"],
    revision = "v1.7.1",
    deps = [
        ":client_model",
        ":json-iterator",
        ":perks",
        ":procfs",
        ":prometheus_common",
        ":protobuf",
        ":x_net",
        ":xxhash",
        ":testify",
        ":x_sys",
        ":pretty",
        ":check.v1",
        ":yaml.v2"
    ],
)

go_get(
    name="backoff",
    get = "github.com/cenkalti/backoff",
    licences = ["MIT"],
    revision="31cc31bb63269a3c813d1f26e0ab058452b7d803",
    deps = [
        ":go-spew",
        ":difflib",
        ":x_sys",
        ":testify",
    ],
)

go_get(
    name = "logrus",
    get = "github.com/sirupsen/logrus",
    licences = ["MIT"],
    revision = "v1.6.0",
    deps = [
        ":x_crypto",
        ":x_sys",
    ],
 )

go_get(
    name = "testify",
    get = "github.com/stretchr/testify/...",
    licences = ["MIT"],
    deps = [
        ":difflib",
        ":go-spew",
        ":objx",
        ":yaml.v3",
        ":yaml.v2",
    ],
    revision = "v1.6.1",
)

go_get(
    name = "mock",
    get = "github.com/golang/mock/...",
    licences = ["apache-2.0"],
    revision = "v1.4.4",
    deps = [
        ":x_net",
        ":x_tools",
    ],
)

#########################
# Package dependencies
#########################

go_get(
    name = "goldmark",
    get = "github.com/yuin/goldmark",
    licences = ["MIT"],
    revision = "v1.2.1",
    install = [
        "ast",
        "renderer/html",
        "text",
    ]
)

go_get(
    name = "x_mod",
    get = "golang.org/x/mod",
    install = [
        "semver",
    ],
    licences = ["bsd-3-clause"],
    revision = "ce943fd02449f621243c9ea6e64098e84752b92b",
)

go_get(
    name = "x_errors",
    get = "golang.org/x/xerrors/...",
    licences = ["bsd-3-clause"],
    revision = "5ec99f83aff198f5fbd629d6c8d8eb38a04218ca",
    deps = [":x_sys"],
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
    revision = "gopls/v0.4.4",
    deps = [
        ":x_crypto",
        ":x_errors",
        ":x_mod",
        ":x_net",
        ":x_sync",
        ":goldmark",
    ],
)

go_get(
    name = "check.v1",
    get = "gopkg.in/check.v1",
    licences = ["bsd-2-clause"],
    revision = "41f04d3bba152ddec2103e299fed053415705330",
    deps = [
        ":pretty",
    ],
)

go_get(
    name = "pretty",
    get = "github.com/kr/pretty",
    licences = ["MIT"],
    revision = "v0.2.1",
    deps = [
        ":kr_text",
    ],
)

go_get(
    name = "kr_text",
    get = "github.com/kr/text",
    licences = ["MIT"],
    revision = "v0.2.0",
)

go_get(
    name = "xxhash",
    get = "github.com/cespare/xxhash",
    licences = ["MIT"],
    revision = "v2.1.1",
)

go_get(
    name = "errors",
    get = "github.com/pkg/errors",
    licences = ["bsd-2-clause"],
    revision = "v0.9.1"
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
    revision = "bf48bf16ab8d622ce64ec6ce98d2c98f916b6303",
    deps = [
        ":cloud-metadata",
        ":x_net",
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
    name = "x_net",
    get = "golang.org/x/net/...",
    revision = "c89045814202410a2d67ec20ecf177ec77ceae7f",
    licences = ["bsd-3-clause"],
    strip = [
        "http2/h2demo",
    ],
    deps = [
        ":x_crypto",
        ":x_text",
        ":x_sys",
    ],
)

go_get(
    name = "x_sys",
    get = "golang.org/x/sys",
    install = [
        "unix",
        "cpu",
    ],
    licences = ["bsd-3-clause"],
    revision = "95b1ffbd15a57cc5abb3f04402b9e8ec0016a52c",
)

go_get(
    name = "cmp",
    get = "github.com/google/go-cmp/cmp",
    licences = ["bsd-2-clause"],
    revision = "v0.5.2",
)


go_get(
    name = "x_crypto",
    get = "golang.org/x/crypto/...",
    revision = "e84da0312774c21d64ee2317962ef669b27ffb41",
    licences = ["bsd-3-clause"],
    deps = [":x_sys"],
)

go_get(
    name = "json-iterator",
    get = "github.com/json-iterator/go",
    licences = ["MIT"],
    deps = [
        ":concurrent",
        ":reflect2",
    ],
    revision = "v1.1.10",
)

go_get(
    name = "procfs",
    get = "github.com/prometheus/procfs/...",
    licences = ["apache-2.0"],
    revision = "v0.1.3",
    deps = [
        ":x_sync",
    ],
)

go_get(
    name = "difflib",
    get = "github.com/pmezard/go-difflib/...",
    licences = ["bsd-3-clause"],
    revision = "v1.0.0",
)

go_get(
    name = "go-spew",
    get = "github.com/davecgh/go-spew/spew",
    licences = ["isc"],
    revision = "v1.1.1",
)

go_get(
    name = "objx",
    get = "github.com/stretchr/objx",
    licences = ["MIT"],
    revision = "v0.3.0",
)

go_get(
    name = "yaml.v3",
    get = "gopkg.in/yaml.v3",
    licences = ["apache-2.0"],
    revision = "v2.3.0",
)

go_get(
    name = "glog",
    get = "github.com/golang/glog",
    licences = ["apache-2.0"],
    revision = "23def4e6c14b4da8ac2ed8007337bc5eb5007998",
)

go_get(
    name = "gogo-protobuf",
    get = "github.com/gogo/protobuf",
    install = [
        "proto",
        "sortkeys",
    ],
    licences = ["bsd-3-clause"],
    revision = "v1.3.1",
    strip = ["proto/proto3_proto"],
)

go_get(
    name = "inf.v0",
    get = "gopkg.in/inf.v0",
    licences = ["bsd-3-clause"],
    revision = "v0.9.1",
)

go_get(
    name = "lru",
    get = "github.com/hashicorp/golang-lru",
    licences = ["mpl-2.0"],
    revision = "v0.5.4",

)

go_get(
    name = "pflag",
    get = "github.com/spf13/pflag",
    licences = ["bsd-3-clause"],
    revision = "v1.0.5",
)


go_get(
    name = "x_time_rate",
    get = "golang.org/x/time/rate",
    licences = ["bsd-3-clause"],
    deps = [":x_net"],
    revision = "3af7569d3a1e776fc2a3c1cec133b43105ea9c2e",
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
    name = "websocket",
    get = "github.com/gorilla/websocket",
    licences = ["bsd-2-clause"],
    revision = "v1.4.2",
)

go_get(
    name = "yaml.v2",
    get = "gopkg.in/yaml.v2",
    revision = "v2.2.5",
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
    revision = "v0.65.0",
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
    revision = "a41fca850d0b6f392931a78cbae438803ea0b886",
)

go_get(
    name = "x_sync",
    get = "golang.org/x/sync/...",
    licences = ["bsd-3-clause"],
    revision = "6e8e738ad208923de99951fe0b48239bfd864f28",
    deps = [
        ":x_net",
    ],
)

go_get(
    name = "golang_protobuf_extensions",
    get = "github.com/matttproud/golang_protobuf_extensions/...",
    licences = ["apache-2.0"],
    revision = "v1.0.1",
    deps = [
        ":protobuf",
    ],
)

go_get(
    name = "kingpin",
    get = "gopkg.in/alecthomas/kingpin.v2",
    licences = ["MIT"],
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
    revision = "fb15b899a75114aa79cc930e33c46b577cc664b1",
)

go_get(
    name = "units",
    get = "github.com/alecthomas/units",
    licences = ["MIT"],
    revision = "f65c72e2690dc4b403c8bd637baf4611cd4c069b",
)

go_get(
    name = "snappy",
    get = "github.com/golang/snappy",
    licences = ["bsd-3-clause"],
    revision = "v0.0.1",
)

go_get(
    name = "uuid",
    get = "github.com/pborman/uuid",
    licences = ["bsd-3-clause"],
    revision = "v1.2.1",
    deps = [
        ":google_uuid",
    ],
)

go_get(
    name = "google_uuid",
    get = "github.com/google/uuid",
    licences = ["bsd-3-clause"],
    revision = "v1.1.2",
)

go_get(
    name = "perks",
    get = "github.com/beorn7/perks/...",
    licences = ["MIT"],
    revision = "v1.0.1",
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
    revision = "v0.2.0",
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
    revision = "v0.10.0",
    deps = [
        ":client_model",
        ":golang_protobuf_extensions",
        ":kingpin",
        ":logrus",
        ":protobuf",
    ],
)

go_get(
    name = "docker",
    get = "github.com/docker/docker",
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
    licences = ["apache-2.0"],
    revision = "v0.4.0",
    deps = [
        ":errors",
    ],
)

go_get(
    name = "docker-go-units",
    get = "github.com/docker/go-units",
    licences = ["apache-2.0"],
    revision = "v0.4.0",
)

go_get(
    name = "docker-distribution",
    get = "github.com/docker/distribution",
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
    install = [
        "",
    ],
    licences = ["apache-2.0"],
    revision = "v1.0.0",
)

go_get(
    name = "opencontainers-imagespec",
    get = "github.com/opencontainers/image-spec",
    install = [
        "specs-go",
        "specs-go/v1",
    ],
    licences = ["apache-2.0"],
    revision = "v1.0.1",
    deps = [
        ":opencontainers-godigest",
    ],
)
