package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grcp-setup/greet/greetpb"
	"io"
	"log"
)

func main() {
	fmt.Println("Hello im a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	//doUnary(c)
	doServerStreaming(c)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting:&greetpb.Greeting{
			FirstName:  "Djamaile",
			SecondName: "Rahamat",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("server streaming error, %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// reached end
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream, %v", err)
		}
		log.Printf("Response from stream %v", msg.GetResult())
	}
}

func doUnary(c greetpb.GreetServiceClient){
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting {
			FirstName:"Djamaile",
			SecondName:"Rahamat",
		},
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("Response afstuderen", res.Result)
}
