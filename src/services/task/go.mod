module github.com/sleepyfox97/otel-grpc-example/src/services/task

go 1.20

replace github.com/sleepyfox97/otel-grpc-example/src/services/uid => ../uid

require (
	github.com/sleepyfox97/otel-grpc-example/src/services/uid v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.60.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231212172506-995d672761c0 // indirect
)
