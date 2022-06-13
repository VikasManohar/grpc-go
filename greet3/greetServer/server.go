package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vikasmanohar/grpc-go/greet3/greetPb"
	"google.golang.org/grpc"
)

type server2 struct{}

func (*server2) Greet(ctx context.Context, req *greetPb.GreetRequest) (*greetPb.GreetResponse, error) {
	fmt.Println("Request to server2 : ", req)
	// firstName := req.GetGreeting().GetFirstName()
	// greetingString := "Hello " + firstName
	// res := &greetPb.GreetResponse{
	// 	SecretMsg: nil,
	// 	Result:    greetingString,
	// }
	return nil, nil
}
func main() {
	fmt.Println("Hello from inside Server3")
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
