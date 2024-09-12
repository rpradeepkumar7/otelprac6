package metrics

import (
    "go.opentelemetry.io/otel/sdk/metric"
)

// NewMeterProvider creates a new MeterProvider
func NewMeterProvider() *metric.MeterProvider {
    return metric.NewMeterProvider()
}
