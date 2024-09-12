package tracing

import (
    "go.opentelemetry.io/otel/sdk/trace"
)

// NewTracerProvider creates a new TracerProvider
func NewTracerProvider() *trace.TracerProvider {
    // Customize tracer provider if needed
    return trace.NewTracerProvider()
}
