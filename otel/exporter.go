package otel

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
)

func newExporter(ctx context.Context) (metric.Exporter, error) {
	return otlpmetricgrpc.New(ctx)
}
