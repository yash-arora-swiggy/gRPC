package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yash-arora-swiggy/gRPC/calculator"
	"google.golang.org/grpc"
)

func main() {

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculator.NewCalculatorServiceClient(cc)
	Calculate(c)

}

func Calculate(c calculator.CalculatorServiceClient) {

	fmt.Println("Making gRPC call for sum")

	req := calculator.SumRequest{
		X: 5,
		Y: 3,
	}

	resp, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling greet grpc unary call: %v", err)
	}

	log.Printf("Sum of the numbers is : %v", resp.Sum)

}
