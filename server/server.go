package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/yash-arora-swiggy/gRPC/calculator"
	"google.golang.org/grpc"
)

type server struct {
	calculator.UnimplementedCalculatorServiceServer
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

func (*server) Prime(req *calculator.PrimeNumbersRequest, resp calculator.CalculatorService_PrimeServer) error {

	fmt.Println("Prime Numbers function invoked for server side streaming")
	for i := 2; i <= int(req.Number); i++ {
		if isPrime(i) {
			res := calculator.PrimeNumbersResponse{
				PrimeNum: int64(i),
			}
			time.Sleep(1000 * time.Millisecond)
			resp.Send(&res)
		}
	}
	return nil

}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
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
