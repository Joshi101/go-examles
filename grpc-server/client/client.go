package main

import (
	"context"
	"log"
	"time"

	pb "rjoshi101/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Ritesh"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf(" client received %s", response.GetMessage())
}
