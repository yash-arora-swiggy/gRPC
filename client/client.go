package main

import (
	"context"
	"fmt"
	"io"
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
	fmt.Print("Enter the method you want to try: ")
	var method string
	fmt.Scan(&method)

	switch method {
	case "Sum":
		var x, y int
		fmt.Print("Enter the 2 numbers you want to add: ")
		fmt.Scan(&x)
		fmt.Scan(&y)
		Add(c, x, y)
	case "Prime":
		var num int
		fmt.Print("Enter the prime number: ")
		fmt.Scan(&num)
		Prime(c, num)
	}

}

func Add(c calculator.CalculatorServiceClient, x int, y int) {

	fmt.Println("Making gRPC call for sum")

	req := calculator.SumRequest{
		X: int64(x),
		Y: int64(y),
	}

	resp, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling greet grpc unary call: %v", err)
	}

	fmt.Println("Sum of the numbers is : ", resp.Sum)
}

func Prime(c calculator.CalculatorServiceClient, num int) {

	fmt.Println("Making gRPC call for prime")

	req := calculator.PrimeNumbersRequest{
		Number: int64(num),
	}

	resp, err := c.Prime(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling greet grpc unary call: %v", err)
	}

	for {
		msg, err := resp.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while receving server stream : %v", err)
		}

		log.Println("Response From Server, Prime Number : ", msg.GetPrimeNum())
	}
}
