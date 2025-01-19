package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "rjoshi101/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (S *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := fmt.Sprintf("Hello, %s!", req.GetName())
	return &pb.HelloResponse{Message: message}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen error %err", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("grpc server is running in port 1234")

	if err := s.Serve(lis); err != nil {
		log.Fatalf(" failed to serve: %v", err)
	}
}
