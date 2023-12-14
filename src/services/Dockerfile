FROM golang:latest

ARG SERVICE_NAME=${SERVICE_NAME}

ENV TZ=Asia/Tokyo

WORKDIR /src

COPY . .

RUN cd ./uid && go mod download
RUN cd ./task && go mod download

RUN cd ./${SERVICE_NAME} && go build -mod=readonly -v -o /app ./main.go
ENTRYPOINT ["/app"]