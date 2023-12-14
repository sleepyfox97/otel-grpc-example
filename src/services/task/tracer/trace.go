package tracer

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"log"
)

func TraceSetting(service string) (trace.Tracer, func(context.Context) error) {
	tr := otel.GetTracerProvider().Tracer(service)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	var shutdown func(context.Context) error
	var err error
	shutdown, err = zipkinTracer(service)
	if err != nil {
		log.Fatal(err)
	}
	return tr, shutdown
}
