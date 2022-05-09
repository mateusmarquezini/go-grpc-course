package main

import (
	"context"
	"github.com/mateusmarquezini/grpc-go-course/calculator/proto"
	"log"
)

func doAvg(c proto.CalculatorServiceClient) {

	stream, err := c.CalculateAvg(context.Background())

	if err != nil {
		log.Fatalf("error while opening the stream: %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 50}

	for _, number := range numbers {
		stream.Send(&proto.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}
