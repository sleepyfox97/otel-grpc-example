package handler

import (
	"context"
	"github.com/sleepyfox97/otel-grpc-example/src/services/task/api"
	"github.com/sleepyfox97/otel-grpc-example/src/services/task/client"
	"go.opentelemetry.io/otel"
	"reflect"
)

type taskService struct {
	api.UnimplementedTaskServiceServer
}

func NewService() (api.TaskServiceServer, error) {
	return &taskService{}, nil
}

func (s *taskService) Create(ctx context.Context, _ *api.CreateRequest) (*api.CreateResponse, error) {
	_, span := otel.GetTracerProvider().Tracer(reflect.TypeOf(s).PkgPath()).Start(ctx, "task create")
	defer span.End()

	id, err := client.UidClient(ctx)
	if err != nil {
		return nil, err
	}
	return &api.CreateResponse{
		Id: *id,
	}, nil
}
