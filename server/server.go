package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/yash-arora-swiggy/gRPC/calculator"
	"google.golang.org/grpc"
)

type server struct {
	calculator.CalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculator.SumRequest) (resp *calculator.SumResponse, err error) {

	num1 := req.GetX()
	num2 := req.GetY()

	res := num1 + num2

	resp = &calculator.SumResponse{
		Sum: res,
	}
	return resp, nil
}

func main() {
	fmt.Println("Starting Server")

	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(s, &server{})

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
