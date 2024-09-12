package emptyexporter

import (
    "context"

    "go.opentelemetry.io/collector/component"
    "go.opentelemetry.io/collector/exporter"
    "go.opentelemetry.io/collector/exporter/exporterhelper"
    "go.opentelemetry.io/collector/pdata/plog"
    "go.opentelemetry.io/collector/pdata/pmetrics"
    "go.opentelemetry.io/collector/pdata/ptrace"
)

const (
    typeStr = "emptyexporter"
)

// NewFactory creates a new factory for the exporter
func NewFactory() exporter.Factory {
    return exporter.NewFactory(
        component.NewID(typeStr),
        createDefaultConfig,
        exporter.WithTraces(createTracesExporter, component.StabilityLevelDevelopment),
        exporter.WithMetrics(createMetricsExporter, component.StabilityLevelDevelopment),
        exporter.WithLogs(createLogsExporter, component.StabilityLevelDevelopment),
    )
}

func createTracesExporter(
    ctx context.Context,
    set exporter.CreateSettings,
    config component.Config) (exporter.Traces, error) {
    s := NewEmptyexporter()
    return exporterhelper.NewTracesExporter(ctx, set, config, s.pushTraces)
}

func createMetricsExporter(
    ctx context.Context,
    set exporter.CreateSettings,
    config component.Config) (exporter.Metrics, error) {
    s := NewEmptyexporter()
    return exporterhelper.NewMetricsExporter(ctx, set, config, s.pushMetrics)
}

func createLogsExporter(
    ctx context.Context,
    set exporter.CreateSettings,
    config component.Config) (exporter.Logs, error) {
    s := NewEmptyexporter()
    return exporterhelper.NewLogsExporter(ctx, set, config, s.pushLogs)
}

// Config holds the configuration for the exporter
type Config struct{}

// createDefaultConfig creates the default configuration for the exporter
func createDefaultConfig() component.Config {
    return &Config{}
}

// Emptyexporter implements the logic for the exporter
type Emptyexporter struct{}

// NewEmptyexporter creates a new instance of the exporter
func NewEmptyexporter() *Emptyexporter {
    return &Emptyexporter{}
}

// pushTraces handles exporting traces
func (e *Emptyexporter) pushTraces(ctx context.Context, traces ptrace.Traces) error {
    // Add trace exporting logic here
    return nil
}

// pushMetrics handles exporting metrics
func (e *Emptyexporter) pushMetrics(ctx context.Context, metrics pmetrics.Metrics) error {
    // Add metric exporting logic here
    return nil
}

// pushLogs handles exporting logs
func (e *Emptyexporter) pushLogs(ctx context.Context, logs plog.Logs) error {
    // Add log exporting logic here
    return nil
}
