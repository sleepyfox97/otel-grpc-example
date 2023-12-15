package tracer

import (
	"context"
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"log"
)

func googleTracer(service string) (func(context.Context) error, error) {
	ctx := context.Background()
	projectID := "GCP_PROJECT_ID"
	exporter, err := texporter.New(texporter.WithProjectID(projectID))
	if err != nil {
		log.Fatalf("texporter.New: %v", err)
	}

	// Identify your application using resource detection
	res, err := resource.New(ctx,
		// Use the GCP resource detector to detect information about the GCP platform
		resource.WithDetectors(gcp.NewDetector()),
		// Keep the default detectors
		resource.WithTelemetrySDK(),
		// Add your own custom attributes to identify your application
		resource.WithAttributes(
			semconv.ServiceNameKey.String(service),
		),
	)
	if err != nil {
		return nil, err
	}
	traceRatio := 1.0
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(traceRatio)),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	//defer tp.ForceFlush(ctx) // flushes any pending spans
	otel.SetTracerProvider(tp)
	return tp.Shutdown, nil
}
