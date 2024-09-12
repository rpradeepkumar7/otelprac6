module emptyexporter

go 1.22

require (
    go.opentelemetry.io/collector v0.108.0
    go.opentelemetry.io/collector/component v0.108.0
    go.opentelemetry.io/collector/exporter/exporterhelper v0.108.0
    go.opentelemetry.io/collector/pdata v0.14.0
    github.com/gogo/protobuf v1.3.2
    github.com/json-iterator/go v1.1.12
    github.com/stretchr/testify v1.9.0
    go.uber.org/goleak v1.3.0
    google.golang.org/grpc v1.65.0
    google.golang.org/protobuf v1.34.2
)

replace go.opentelemetry.io/collector/pdata => ../pdata
