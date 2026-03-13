package main

import (
	"context"
	pb "grpc-test-project/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}
	log.Printf("Response: %s", res.Message)
}
