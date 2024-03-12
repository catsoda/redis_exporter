package otel

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func initProvider() (func(context.Context) error, error) {

	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceName("test-service"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// 创建指标导出器的选项，指定导出的目标地址为 10.1.0.6:4317
	opts := []otlpmetricgrpc.Option{
		otlpmetricgrpc.WithEndpoint("10.1.0.6:4317"),
	}

	exp, err := otlpmetricgrpc.New(ctx, opts...)
	if err != nil {
		panic(err)
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(exp)),
		metric.WithResource(res))
	defer func() {
		if err := meterProvider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
	otel.SetMeterProvider(meterProvider)

	// From here, the meterProvider can be used by instrumentation to collect
	// telemetry.
	return meterProvider.Shutdown, nil
}
