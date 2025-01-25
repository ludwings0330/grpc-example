# grpc-example

# grpc-example

# grpc-example

1. Install Go gRPC and protocol Buffers plugins

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. Compile proto

```shell
cd proto
protoc --go_out=. --go-grpc_out=. model.proto
```

3. run server

```shell
go run server.go
```

4. run client

```shell
go run client.go
```
