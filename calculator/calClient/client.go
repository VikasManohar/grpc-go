package main

import (
	"context"
	"fmt"
	"github.com/vikasmanohar/grpc-go/calculator/calPb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Inside CalculatorService Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Error connecting to the server", err)
	}
	defer cc.Close()
	c := calPb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c calPb.CalculatorServiceClient) {
	req := &calPb.Input{
		In1: 33,
		In2: 99,
	}
	result, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln("Error while sending the request to the client", err)
	}
	fmt.Println("Sum of the numbers", req.In1, req.In2, "are", result.Res)
}