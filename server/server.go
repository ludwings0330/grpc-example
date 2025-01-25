package main

import (
	"context"
	"log"
	"net"

	pb "proto/modelservice"

	"google.golang.org/grpc"
)

// 서버 구현
type server struct {
	pb.UnimplementedModelServiceServer
}

func (s *server) GetModelInfo(ctx context.Context, req *pb.ModelRequest) (*pb.ModelResponse, error) {
	log.Printf("[Server] GetModelInfo")
	log.Printf("Received request for model: %s, version: %s", req.ModelName, req.Version)
	return &pb.ModelResponse{
		Status:  "Success",
		Message: "Model is available",
	}, nil
}

func (s *server) RunInference(ctx context.Context, req *pb.InferenceRequest) (*pb.InferenceResponse, error) {
	log.Printf("[Server] RunInference")
	log.Printf("Received request for model: %s, features: %+v", req.ModelName, req.InputFeatures)
	return &pb.InferenceResponse{
		Prediction: 7.7,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterModelServiceServer(grpcServer, &server{})

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
