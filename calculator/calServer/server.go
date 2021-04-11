package main

import (
	"context"
	"fmt"
	"github.com/vikasmanohar/grpc-go/calculator/calPb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *calPb.Input) (*calPb.Result, error) {
	fmt.Println("Request received", req)
	input1, input2 := req.In1, req.In2
	res := &calPb.Result{
		Res: input1 + input2,
	}
	return res, nil
}

func main() {
	fmt.Println("Inside server code of CalculatorService")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalln("Error setting up listener on port", err)
	}
	s := grpc.NewServer()

	calPb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln("Error serving CalculatorService", err)
	}
}
