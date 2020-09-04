package(default_visibility = ["PUBLIC"] )

subinclude("///third_party/subrepos/pleasings//remote")

#########################
# Packages referenced by go files in repo
#########################

go_get(
    name = "go-flags",
    get = "github.com/jessevdk/go-flags",
    licences = ["bsd-3-clause"],
    revision = "v1.4.0",
    hashes = [
        "031e87226e1f3f1660aa1e859f9f0417e5d42464fb64969e685811f9a723ce6b",
    ],
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
    hashes = [
        "a154bd548148a411378f97cc3f1a54f323a93ba9d611207825e0e7c7e5b38dc6",
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
    hashes = [
        "e4b75eb61ca5837e61d035b69dfd37b09f2b4ff487b7190f3a7ed44675755856",
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
     hashes = [
        "ad261ee7ef01dc4800466c7e8dc5d31028f30cd1705f751e63910fd475449fee",
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
    hashes = [
        "e15efaab7fa5ea657cd985f03bb9e0bd535a8bd01d0f6c37e80a698e767fa83f",
    ],
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
    hashes = [
        "62c71cc09d45f5fb2b2a64a877a99b333295af4425918a655c2df8c1eac7879b",
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
    ],
    hashes = [
        "d7dc2d7863ed6960a466f7814bd22d484b87b84cb1b6a544ee5a7e1bc5727fe9",
    ],
)

go_get(
    name = "x_mod",
    get = "golang.org/x/mod",
    install = [
        "semver",
    ],
    licences = ["bsd-3-clause"],
    revision = "ce943fd02449f621243c9ea6e64098e84752b92b",
    hashes = [
        "7c47cd30783ed9fd7f86dcbb751978c9c700f88e54b9d95f34a5c3bb5635c35a",
    ],
)

go_get(
    name = "x_errors",
    get = "golang.org/x/xerrors/...",
    licences = ["bsd-3-clause"],
    revision = "5ec99f83aff198f5fbd629d6c8d8eb38a04218ca",
    deps = [":x_sys"],
    hashes = [
        "3bebb60d4b80827db801f179f9dbc95e20860a173fbc4d0704850020b339e290",
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
    revision = "gopls/v0.4.4",
    deps = [
        ":x_crypto",
        ":x_errors",
        ":x_mod",
        ":x_net",
        ":x_sync",
        ":goldmark",
    ],
    hashes = [
        "aa000721e26292d58281d47626e67c0408073054f8aeb3ab619ddd03dbd9ce9b",
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
    hashes = [
        "37620c3963f914ed108fc2941f9fadec02a49ee32813ce8a2b9e25dd685377e3",
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
    hashes = [
        "9016b8420a5508c610999447f77c6107806949f48ec096c5a1d6511031054c34",
    ],
)

go_get(
    name = "kr_text",
    get = "github.com/kr/text",
    licences = ["MIT"],
    revision = "v0.2.0",
    hashes = [
        "0c17393843bc8c9bad3530a43741843f9050013b992e8ac57f3c2ba7e6c41834",
    ],
)

go_get(
    name = "xxhash",
    get = "github.com/cespare/xxhash",
    licences = ["MIT"],
    revision = "v2.1.1",
    hashes = [
        "22e2cbb9e4ede0891fb15dd1fbf96ac7669e8e24695380a5b019722a2ceefcc3",
    ],
)

go_get(
    name = "errors",
    get = "github.com/pkg/errors",
    licences = ["bsd-2-clause"],
    revision = "v0.9.1",
    hashes = [
        "f2efacf221626cee92274baf18713452b47908117c8aafe0f1c8bde455670f56",
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
    revision = "bf48bf16ab8d622ce64ec6ce98d2c98f916b6303",
    deps = [
        ":cloud-metadata",
        ":x_net",
    ],
    hashes = [
        "1e30fed14c6edabde86ad5f42399728ff8f0c0d648bf6bc75455ba6197337da3",
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
    hashes = [
        "8ea8625363e0d378b9474383fbb30630f163852bf614b4f4476e5190d5bbbc99",
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
    hashes = [
        "b057e07e945bc2957121d9cd26b1448b4c838bb4e202f06811ae4086f46474e0",
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
    hashes = [
        "024d1015f868da12ee23239c4674ba57f485a7196c92227190b7c7c61be22202",
    ],
)

go_get(
    name = "cmp",
    get = "github.com/google/go-cmp/cmp",
    licences = ["bsd-2-clause"],
    revision = "v0.5.2",
    hashes = [
        "2638b9dacdf9b82a7e7c0b184512ff93cd644d49d85c4db745fed4c56ebee184",
    ],
)


go_get(
    name = "x_crypto",
    get = "golang.org/x/crypto/...",
    revision = "e84da0312774c21d64ee2317962ef669b27ffb41",
    licences = ["bsd-3-clause"],
    deps = [":x_sys"],
    hashes = [
        "280a85a2ab2625c90b262659ad530b251ef76eca976b9c0493340eb192d67db6",
    ],
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
    hashes = [
        "18fd7d82bd6cb62c2adb6d39fe9199e58a72904d21032210ffd22aed07d0d6f3",
    ],
)

go_get(
    name = "procfs",
    get = "github.com/prometheus/procfs/...",
    licences = ["apache-2.0"],
    revision = "v0.1.3",
    deps = [
        ":x_sync",
    ],
    hashes = [
        "05a13e3a105c071f8ce0ae9f1de7b7bfbd446e63dd785c84cdd7f4a303376367",
    ],
)

go_get(
    name = "difflib",
    get = "github.com/pmezard/go-difflib/...",
    licences = ["bsd-3-clause"],
    revision = "v1.0.0",
    hashes = [
        "3d6ad64cae5e8c10ce29f343f296f36e79a352e7408d8bbb580df11d3bb86c44",
    ],
)

go_get(
    name = "go-spew",
    get = "github.com/davecgh/go-spew/spew",
    licences = ["isc"],
    revision = "v1.1.1",
    hashes = [
        "090995aa7f9f16ceb8f9d3c6ab37024be764caac6023f922329e3827501ee179",
    ],
)

go_get(
    name = "objx",
    get = "github.com/stretchr/objx",
    licences = ["MIT"],
    revision = "v0.3.0",
    hashes = [
        "9cdc33c4a5fc62924574e47bbedab314ffcfc313f234fc33cdab471a7609a0a3",
    ],
)

go_get(
    name = "yaml.v3",
    get = "gopkg.in/yaml.v3",
    licences = ["apache-2.0"],
    revision = "v2.3.0",
    hashes = [
        "9254bbb32117a4be5c233f61731c622ca88008309d4d86be6f1b7da2c5d86ce9",
    ],
)

go_get(
    name = "glog",
    get = "github.com/golang/glog",
    licences = ["apache-2.0"],
    revision = "23def4e6c14b4da8ac2ed8007337bc5eb5007998",
    hashes = [
        "1c0213d77145f5678d842ef1245e25bfe2831360b14a4ee4fa244e02b09e364c",
    ],
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
    hashes = [
        "2ba7a3244df95461a553f5dd8162d8d4528aca095373a6eb1369cfbf8578fdab",
    ],
)

go_get(
    name = "inf.v0",
    get = "gopkg.in/inf.v0",
    licences = ["bsd-3-clause"],
    revision = "v0.9.1",
    hashes = [
        "8a1234d6fa309c8b3b5ea5353adac45797e19b8836a0ba21718d068f871a618c",
    ],
)

go_get(
    name = "lru",
    get = "github.com/hashicorp/golang-lru",
    licences = ["mpl-2.0"],
    revision = "v0.5.4",

    hashes = [
        "67edd988a80a5f18bbd0363cfc74b9c6782a49d277a1f27ea1892ee32041f61b",
    ],
)

go_get(
    name = "pflag",
    get = "github.com/spf13/pflag",
    licences = ["bsd-3-clause"],
    revision = "v1.0.5",
    hashes = [
        "72d85b140b2866b869de748383faaaa6a7a8d9e30730d0eec8bd70e081e2a5ac",
    ],
)


go_get(
    name = "x_time_rate",
    get = "golang.org/x/time/rate",
    licences = ["bsd-3-clause"],
    deps = [":x_net"],
    revision = "3af7569d3a1e776fc2a3c1cec133b43105ea9c2e",
    hashes = [
        "a594f7c19df6a65c127f9e71d07a07fca9b4d145e16e008450445dfe70cbb3bf",
    ],
)

go_get(
    name = "x_text",
    get = "golang.org/x/text/...",
    revision = "v0.3.0",
    strip = [
        "cmd",
        "message/pipeline",
    ],
    hashes = [
        "72b1315e4f1e401a055feaf9722bb69e95edc36c2526ce6b3bb62c2b8fb39ce2",
    ],
)

go_get(
    name = "websocket",
    get = "github.com/gorilla/websocket",
    licences = ["bsd-2-clause"],
    revision = "v1.4.2",
    hashes = [
        "289da6a8cfbff95cdb23eceb59ef91b2a8d4e4fbca7c8fa352c48dfd903d5091",
    ],
)

go_get(
    name = "yaml.v2",
    get = "gopkg.in/yaml.v2",
    revision = "v2.2.5",
    hashes = [
        "a8f3dda3391cda769609d4159d440c05e857a8dce57fe2b1a03ae5060421f6ea",
    ],
)

go_get(
    name = "concurrent",
    get = "github.com/modern-go/concurrent/...",
    revision = "1.0.3",
    hashes = [
        "810908664f8b0073fcc52d8771586936d53bf4f5bff99e5760a2b7d46b2423e0",
    ],
)

go_get(
    name = "cloud-metadata",
    get = "cloud.google.com/go/compute/metadata",
    licences = ["apache-2.0"],
    revision = "v0.65.0",
    hashes = [
        "5583d818debe00ff9a7cb084d0f1f37571bd66d6b202096df81d663ad3d607af",
    ],
)

go_get(
    name = "reflect2",
    get = "github.com/modern-go/reflect2/...",
    licences = ["apache-2.0"],
    revision = "v1.0.1",
    deps = [
        ":concurrent",
    ],
    hashes = [
        "9f0a3cc245f3314d4d58e9d707ede245836f776830d1cd2667f563b37ad89d46",
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
    hashes = [
        "1ffe2aef867d2609e0de0f23f817fa790ff5d5690e8ba430e71e1e508b765631",
    ],
)

go_get(
    name = "x_sync",
    get = "golang.org/x/sync/...",
    licences = ["bsd-3-clause"],
    revision = "6e8e738ad208923de99951fe0b48239bfd864f28",
    deps = [
        ":x_net",
    ],
    hashes = [
        "46fd8f19b7925822de8d3aa528a20d82353fe7bc8fb85986fdd08296e0725f4b",
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
    hashes = [
        "c075fd2185452e396855715425cd32a8b0b4da34ccc1c2cde3b831ef4e6566de",
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
    hashes = [
        "9f4a88b2256d7fcbc7ac7c06c0a6c6a72ec45cffcd257c1ae220dcf148d8d9ed",
    ],
)

go_get(
    name = "template",
    get = "github.com/alecthomas/template",
    licences = ["bsd-3-clause"],
    revision = "fb15b899a75114aa79cc930e33c46b577cc664b1",
    hashes = [
        "5cbeb4e9a2eaae299a0ce4d1f265221919234e89efd42b8c379bcdd3d6815238",
    ],
)

go_get(
    name = "units",
    get = "github.com/alecthomas/units",
    licences = ["MIT"],
    revision = "f65c72e2690dc4b403c8bd637baf4611cd4c069b",
    hashes = [
        "f906a4e8b50e59eeef9f3b638c00211ed58558b9ef28c1d2e926670e1f874818",
    ],
)

go_get(
    name = "snappy",
    get = "github.com/golang/snappy",
    licences = ["bsd-3-clause"],
    revision = "v0.0.1",
    hashes = [
        "ddaeb18a5d8a616bb21453359f2f6987488944fa4df59845ed6f3fcb5675c653",
    ],
)

go_get(
    name = "uuid",
    get = "github.com/pborman/uuid",
    licences = ["bsd-3-clause"],
    revision = "v1.2.1",
    deps = [
        ":google_uuid",
    ],
    hashes = [
        "1208c227f4035d8957d641ccaf6ecbeee515be1ac1eaccc9cb6383afc3af19f2",
    ],
)

go_get(
    name = "google_uuid",
    get = "github.com/google/uuid",
    licences = ["bsd-3-clause"],
    revision = "v1.1.2",
    hashes = [
        "f14398d9c01543a46f9f35ff08d082abc940cfed185654fe58fe34a76d80720f",
    ],
)

go_get(
    name = "perks",
    get = "github.com/beorn7/perks/...",
    licences = ["MIT"],
    revision = "v1.0.1",
    hashes = [
        "25df30c2e6bada572d4443ad5317037e79c3372ab5fa79cc0c57e41091abc5ab",
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
    hashes = [
        "62be2ea9be0ba860f0032e43e32004a5c71a4e182ff949dbc53338fafc956935",
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
    hashes = [
        "e94cfec5b0d520d75a418e357b27c6d04350da7639489a677d1e60094184622e",
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
    hashes = [
        "8d3df5d142121396ccab6d8414aad48ed187cd280d00713c925c9dc0c402e6bb",
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
    hashes = [
        "21f4dbed0a24250b2b6a5fdddf54f7f1f58c2bf857395314ed76cde0b8154169",
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
    hashes = [
        "920d82dc00a2ef1f4de2b827e336962e0f396c9608c4c2bdd953da16180892ff",
    ],
)

go_get(
    name = "docker-go-units",
    get = "github.com/docker/go-units",
    licences = ["apache-2.0"],
    revision = "v0.4.0",
    hashes = [
        "ae04b6b932f0f1562b2deb932f7d5300a17c9e32a97424ca51a3c0281f05cdef",
    ],
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
    hashes = [
        "602cd454b021f39b73ccbb51a2b66ab11154f9086751a7fea3c7a2d881e9299e",
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
    hashes = [
        "3e4a6193395c946dfa682c5f582ec6ad3710633e972206602796e0cac8aa564b",
    ],
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
    hashes = [
        "1fc5497a6bbfeaac4c186026faf071e2a413c58c90d468e890109403934e561f",
    ],
)
