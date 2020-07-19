package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grcp-setup/calculator/calculatorpb"
	"log"
)

func main() {
	fmt.Println("Calculator client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	doUnary(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient){
	req := &calculatorpb.SumRequest{
		FirstNumber:  100,
		SecondNumber: 200,
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("Response sum", res.SumResult)
}
