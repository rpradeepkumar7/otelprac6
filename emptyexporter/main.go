package main

import (
    "log"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/prometheus"
    "go.opentelemetry.io/otel/sdk/trace"
    "go.opentelemetry.io/otel/sdk/metric"
    "go.opentelemetry.io/otel/metric/global"
)

func main() {
    // Initialize exporter (Prometheus or other as needed)
    promExporter, err := prometheus.New(prometheus.Config{})
    if err != nil {
        log.Fatalf("failed to initialize prometheus exporter: %v", err)
    }

    // Register the exporter
    global.SetMeterProvider(metric.NewMeterProvider(promExporter))

    // Initialize tracing provider (example for using Prometheus)
    tracerProvider := trace.NewTracerProvider()
    otel.SetTracerProvider(tracerProvider)

    log.Println("Exporter running...")
}
