package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grcp-setup/greet/greetpb"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct {

}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Functie invoked met %v\n", req)
	voornaam := req.GetGreeting().GetFirstName()
	result := voornaam + " is afgestudeerd"
	res := &greetpb.GreetResponse{Result:result,}
	return res, nil
}

func (*server) 	GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("Greet many timesFunctie invoked met %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++{
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main()  {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed connecting %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
