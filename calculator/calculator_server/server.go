package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grcp-setup/calculator/calculatorpb"
	"log"
	"net"
	"context"
)

type server struct {

}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	sum := firstNumber + secondNumber

	res := &calculatorpb.SumResponse{SumResult:sum}

	return res, nil
}

func main()  {
	fmt.Println("Calculator service")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed connecting %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
