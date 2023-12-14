package handler

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sleepyfox97/otel-grpc-example/src/services/uid/api"
	"go.opentelemetry.io/otel"
	"reflect"
)

type uidService struct {
	api.UnimplementedUIDServiceServer
}

func NewService() (api.UIDServiceServer, error) {
	return &uidService{}, nil
}

func (s *uidService) Generate(ctx context.Context, _ *api.GenerateRequest) (*api.GenerateResponse, error) {
	_, span := otel.GetTracerProvider().Tracer(reflect.TypeOf(s).PkgPath()).Start(ctx, "uid generate")
	defer span.End()

	id := uuid.New().String()
	fmt.Sprintf("call uid generate: %s\n", id)
	return &api.GenerateResponse{ID: id}, nil
}
