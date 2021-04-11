package main

import (
	"context"
	"fmt"
	"github.com/vikasmanohar/grpc-go/greet2/greetPb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server2 struct{}

func (*server2) Greet(ctx context.Context, req *greetPb.GreetRequest) (*greetPb.GreetResponse, error) {
	fmt.Println("Request to server2 : ", req)
	firstName := req.GetGreeting().GetFirstName()
	greetingString := "Hello " + firstName
	res := &greetPb.GreetResponse{
		Result: greetingString,
	}
	return res, nil
}
func main() {
	fmt.Println("Hello from inside Server2")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Error opening port", err)
	}
	s := grpc.NewServer()

	greetPb.RegisterGreetServiceServer(s, &server2{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
