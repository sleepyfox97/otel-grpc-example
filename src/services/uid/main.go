package main

import (
	"context"
	_ "context"
	"fmt"
	"github.com/sleepyfox97/otel-grpc-example/src/services/uid/api"
	"github.com/sleepyfox97/otel-grpc-example/src/services/uid/handler"
	"github.com/sleepyfox97/otel-grpc-example/src/services/uid/tracer"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	_ "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	_ "google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	port := 5051
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	//tracerを呼び出す
	tr, shutdown := tracer.TraceSetting("uid")
	defer func() {
		if err = shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()
	_, span := tr.Start(ctx, "uid", trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	h, err := handler.NewService()
	api.RegisterUIDServiceServer(s, h)

	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(s, healthSrv)
	healthSrv.SetServingStatus("uid", healthpb.HealthCheckResponse_SERVING)

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		err = s.Serve(listener)
		if err != nil {
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
