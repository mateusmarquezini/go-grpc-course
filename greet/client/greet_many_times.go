package main

import (
	"context"
	pb "github.com/mateusmarquezini/grpc-go-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {

	req := &pb.GreetRequest{FirstName: "Mateus"}

	stream, err := c.GreetStream(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
