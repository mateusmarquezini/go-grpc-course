package main

import (
	"context"
	pb "github.com/mateusmarquezini/grpc-go-course/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})

	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}
