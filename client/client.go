package main

import (
	"context"
	"log"
	"time"

	pb "proto/modelservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewModelServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ModelRequest{
		ModelName: "test_model",
		Version:   "v1.0",
		Features:  []string{"feature1", "feature2"},
	}
	log.Printf("[Client] GetModelInfo")
	resp, err := client.GetModelInfo(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetModelInfo: %v", err)
	}

	log.Printf("Response: %s", resp.Message)

	log.Printf("[Client] runInference")
	InferReq := &pb.InferenceRequest {
		ModelName: "test_model",
		InputFeatures: map[string]float32{"A": 10.3, "B": 12.5},
	}
	InferResp, err := client.RunInference(ctx, InferReq)
	if err != nil {
		log.Fatal("Error calling RunInference: %v", err)
	}

	log.Printf("RunInference Response: %v", InferResp)
}
