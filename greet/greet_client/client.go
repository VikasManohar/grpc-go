package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/vikasmanohar/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("I am client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server %v", err)
	}
	defer func() {
		err := cc.Close()
		if err != nil {
			log.Fatalln("error closing client connection to the server", err)
		}
	}()
	c := greetpb.NewGreetServiceClient(cc)
	// doUnary(c)

	// doServerStreaming(c)

	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("========================")
	fmt.Println("unary rpc")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vikas",
			LastName:  "M",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling greet %v\n", err)
	}
	fmt.Printf("Response from Unary RPC: %v\n", res.GetResult())
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("========================")
	fmt.Println("Inside ServerStreaming client")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Vikas",
			LastName:  "M",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalln("Error while calling GreetManyTiems RPC", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//reached the end of this stream
			break
		}
		if err != nil {
			log.Fatalln("error while reading stream", err)
		}
		fmt.Println("Response from GreetManyTimes/ServerStreaming RPC", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("========================")
	fmt.Println("Inside ClientStreaming client")
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jon",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Bob",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Tim",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Pony",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error calling LongGreet %v\n", err)
	}
	for _, req := range requests {
		fmt.Println("sending request ", req)
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("error calling LongGreet %v\n", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("error receiving response from LongGreet ", err)
	}
	fmt.Println("resopnse from LongGreet ", res)
}
