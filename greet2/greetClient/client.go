package main

import (
	"context"
	"fmt"
	"github.com/vikasmanohar/grpc-go/greet2/greetPb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello from client2")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Error connecting to port", err)
	}
	defer cc.Close()
	c := greetPb.NewGreetServiceClient(cc)
	UnaryGreet(c)

}

func UnaryGreet(c greetPb.GreetServiceClient) {
	req := &greetPb.GreetRequest{
		Greeting: &greetPb.Greeting{
			FirstName: "Vikas",
			LastName:  "Manohar",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln("error receiving response from the server", err)
	}
	fmt.Println("Response from server", res)
}
