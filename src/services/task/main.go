package main

import (
	_ "context"
	"fmt"
	"github.com/sleepyfox97/otel-grpc-example/src/services/task/api"
	"github.com/sleepyfox97/otel-grpc-example/src/services/task/handler"
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
	port := 5050
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	h, err := handler.NewService()
	api.RegisterTaskServiceServer(s, h)

	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(s, healthSrv)
	healthSrv.SetServingStatus("task", healthpb.HealthCheckResponse_SERVING)

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		err := s.Serve(listener)
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
