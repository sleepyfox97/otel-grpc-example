package handler

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sleepyfox97/otel-grpc-example/src/services/uid/api"
)

type uidService struct {
	api.UnimplementedUIDServiceServer
}

func NewService() (api.UIDServiceServer, error) {
	return &uidService{}, nil
}

func (s *uidService) Generate(ctx context.Context, _ *api.GenerateRequest) (*api.GenerateResponse, error) {

	//以下2行を加えることで，分散トレーシングにGenerate関数の範囲を含めることができる．
	//_, span := otel.Tracer("uid").Start(ctx, "handler")
	//defer span.End()

	id := uuid.New().String()
	fmt.Sprintf("call uid generate: %s\n", id)
	return &api.GenerateResponse{ID: id}, nil
}
