module github.com/shiqinfeng1/goMono

go 1.22

require (
	cloud.google.com/go/compute/metadata v0.2.3
	github.com/docker/docker v24.0.7+incompatible
	github.com/fluent/fluent-logger-golang v1.9.0
	github.com/go-kratos/gateway v0.0.0-20231213100859-abf713bc415f
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20231113102135-421dbc7dae0f
	github.com/go-kratos/kratos/contrib/log/zerolog/v2 v2.0.0-20231116090954-1e4e37ad8735
	github.com/go-kratos/kratos/contrib/metrics/prometheus/v2 v2.0.0-20231122041305-e84cddeabddb
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20231107125212-e38a364340c5
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20231122041305-e84cddeabddb
	github.com/go-kratos/kratos/v2 v2.7.3
	github.com/go-kratos/swagger-api v1.0.1
	github.com/go-sql-driver/mysql v1.7.1
	github.com/gogf/gf/v2 v2.6.1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/golang/protobuf v1.5.3
	github.com/google/uuid v1.4.0
	github.com/google/wire v0.5.0
	github.com/gorilla/handlers v1.5.2
	github.com/jmoiron/sqlx v1.3.5
	github.com/nacos-group/nacos-sdk-go v1.1.4
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.17.0
	github.com/rs/zerolog v1.31.0
	github.com/stretchr/testify v1.8.4
	go.etcd.io/etcd/client/v3 v3.5.10
	go.opentelemetry.io/otel v1.17.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.4.1
	go.opentelemetry.io/otel/sdk v1.17.0
	go.uber.org/automaxprocs v1.5.1
	go.uber.org/multierr v1.11.0
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1
	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	cloud.google.com/go/compute v1.23.0 // indirect
	github.com/BurntSushi/toml v1.2.0 // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.18 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/distribution/reference v0.5.0 // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-errors/errors v1.4.2 // indirect
	github.com/go-kratos/aegis v0.2.1-0.20230616030432-99110a3f05f4 // indirect
	github.com/go-kratos/feature v0.0.0-20230724160043-79ea0633def6 // indirect
	github.com/go-kratos/grpc-gateway/v2 v2.5.1-0.20210811062259-c92d36e434b1 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v1.1.2 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/lufia/plan9stats v0.0.0-20230326075908-cb1d2100619a // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/prometheus/client_model v0.4.1-0.20230718164431-9a2bf3000d16 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/shirou/gopsutil/v3 v3.23.6 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/tinylib/msgp v1.1.6 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	go.etcd.io/etcd/api/v3 v3.5.10 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.10 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.4.1 // indirect
	go.opentelemetry.io/otel/metric v1.17.0 // indirect
	go.opentelemetry.io/otel/trace v1.17.0 // indirect
	go.opentelemetry.io/proto/otlp v0.12.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
	google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.5.1 // indirect
)
