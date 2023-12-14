package handler

import (
	"context"
	"github.com/sleepyfox97/otel-grpc-example/src/services/task/api"
	"github.com/sleepyfox97/otel-grpc-example/src/services/task/client"
)

type taskService struct {
	api.UnimplementedTaskServiceServer
}

func NewService() (api.TaskServiceServer, error) {
	return &taskService{}, nil
}

func (s *taskService) Create(ctx context.Context, _ *api.CreateRequest) (*api.CreateResponse, error) {

	//以下2行を加えることで，分散トレーシングにGenerate関数の範囲を含めることができる．
	//_, span := otel.Tracer("uid").Start(ctx, "handler")
	//defer span.End()

	id, err := client.UidClient(ctx)
	if err != nil {
		return nil, err
	}
	return &api.CreateResponse{
		Id: *id,
	}, nil
}
