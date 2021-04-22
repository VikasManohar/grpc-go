package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/vikasmanohar/grpc-go/calculator/calPb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Inside CalculatorService Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Error connecting to the server", err)
	}
	defer func() {
		err := cc.Close()
		if err != nil {
			log.Fatalln("error closing client connection to the server", err)
		}
	}()
	c := calPb.NewCalculatorServiceClient(cc)
	doUnary(c)

	doServerStreaming(c)

	doClientStreaming(c)

	doBiDiStreaming(c)
}

func doUnary(c calPb.CalculatorServiceClient) {
	fmt.Println("Inside unary Client")
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

func doServerStreaming(c calPb.CalculatorServiceClient) {
	fmt.Println("Inside Server Streaming Client:")
	n := 99
	req := &calPb.PrimeNumbeRequest{
		Int1: int32(n),
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalln("Error calling PrimeNumberDecomposition")
	}
	fmt.Print("prime number decomposition of ", n, ": ")
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("Error reading stream")
		} else {
			fmt.Print(msg.Res, ",")
		}
	}
}

func doClientStreaming(c calPb.CalculatorServiceClient) {
	fmt.Println("Inside Client Streaming Client")

	requests := []*calPb.ComputeAverageRequest{
		{
			Input: 1,
		},
		{
			Input: 2,
		},
		{
			Input: 3,
		},
		{
			Input: 4,
		},
	}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalln("Error calling PrimeNumberDecomposition")
	}
	fmt.Println("Average of numbers ")

	for _, req := range requests {
		fmt.Print(req.GetInput(), " ")
		err := stream.Send(req)
		if err != nil {
			log.Fatalln("error calling ComputeAverage", err)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error receiving response from ComputeAverage", err)
	}
	fmt.Println("is ", res.GetRes())
}

func doBiDiStreaming(c calPb.CalculatorServiceClient) {
	fmt.Println("Inside Bi Directional Streaming Client")

	requests := []*calPb.FindMaximumRequest{
		{
			Input: 1,
		},
		{
			Input: 5,
		},
		{
			Input: 3,
		},
		{
			Input: 6,
		},
		{
			Input: 2,
		},
		{
			Input: 20,
		},
	}
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalln("error while trying to create client stream ", err)
	}
	waitChan := make(chan struct{})
	fmt.Println("The numbers are ")
	go func() {
		for _, req := range requests {
			fmt.Println(req.GetInput(), " ")
			err := stream.Send(&calPb.FindMaximumRequest{
				Input: req.GetInput(),
			})
			if err != nil {
				log.Fatalln("Error sending request to the server ", err)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("Error getting response from server ", err)
			}
			fmt.Println("max number so far : ", res.GetRes())
		}
		close(waitChan)
	}()
	<-waitChan
}
