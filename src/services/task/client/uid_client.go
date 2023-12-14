package client

import (
	"context"
	uid "github.com/sleepyfox97/otel-grpc-example/src/services/uid/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client uid.UIDServiceClient
)

func UidClient(ctx context.Context) (*string, error) {
	address := "uid.default.svc.cluster.local:5051"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client = uid.NewUIDServiceClient(conn)

	var res *uid.GenerateResponse
	res, err = client.Generate(ctx, &uid.GenerateRequest{})
	if err != nil {
		return nil, err
	}
	return &res.ID, nil
}
