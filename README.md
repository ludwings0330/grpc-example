# grpc-example

## How to use

1. Install Go gRPC and protocol Buffers plugins

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. Compile proto

```shell
cd proto
# mkdir modelservice
protoc --go_out=./modelservice --go-grpc_out=./modelservice --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative model.proto
```

3. run server

```shell
go run server.go
```

4. run client

```shell
go run client.go
```

## Summary

### 1. Generating Go Structs from proto file

- The .proto file is used to define the data `structures` and `services`, and Go code is generated using the `protoc` compiler along with the `protoc-gen-go` and `protoc-gen-go-grpc` plugins
- Example command to generate Go code from a proto file

```shell
protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative model.proto
```

---

### 2. Proto Structure: message vs service

- The `.proto` file contains two primary elements:
  1. `message`: Defines the structure of data exchanged between the client and server. These are similar to structs in Go.
  2. `service`: Defines how the message are used by specifying the gRPC methods that can be invoked remotely.

---

### 3. Service Definition and Implementation in Go

- When `protoc` generates Go code from the proto file, it creates an `interface` based on the service definition
- The server embed this interface to implement the reuiqred functionality

```go
type server struct {
    modelservice.UnimplementedModelServiceServer
}

func (s *server) GetModelInfo(ctx context.Context, req *modelservice.ModelRequest) (*modelservice.ModelResponse, error) {
    return &modelservice.ModelResponse{
        Status:  "Success",
        Message: "Model is available",
    }, nil
}
```

```go
conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
if err != nil {
    log.Fatalf("Failed to connect: %v", err)
}
defer conn.Close()

client := modelservice.NewModelServiceClient(conn)
response, err := client.GetModelInfo(context.Background(), &modelservice.ModelRequest{ModelName: "example", Version: "1.0"})
if err != nil {
    log.Fatalf("Error calling GetModelInfo: %v", err)
}
log.Println("Response:", response)
```

---

### 4. gRPC Workflow Overview

1. Define the `proto` file with messages and services.
2. Generate Go code using `protoc` with the appropriate plugins.
3. Implement the gRPC service on the server-side
   - _Server MUST implement the interface by embedding the `Unimplemented<ServierName>Server`._
4. Call the gRPC service from the client-side
   - _In the client, embedding or implementing the service interface is NOT required._
     - The client simply invokes methods by creating an instance of the gRPC client.
   - gRPC generates a client that contains all the remote methods, and we directly call them.

---

### 5. Running Multiple gRPC Services:

1. Running Multiple Services on a Single Port

- You can register multiple gRPC services on a single `grpc.Server` instance and run them on the same port

```go
listener, err := net.Listen("tcp", ":50051")
grpcServer := grpc.NewServer()
modelservice.RegisterModelServiceServer(grpcServer, &modelServer{})
userservice.RegisterUserServiceServer(grpcServer, &userServer{})
grpcServer.Serve(listener)
```
