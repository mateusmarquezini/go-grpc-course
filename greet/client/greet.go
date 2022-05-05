package main

import (
	"context"
	pb "github.com/mateusmarquezini/grpc-go-course/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Mateus",
	})

	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}

	log.Printf("greeting: %s\n", res.Result)
}
