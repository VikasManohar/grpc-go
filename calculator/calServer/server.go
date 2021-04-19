package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/vikasmanohar/grpc-go/calculator/calPb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) PrimeNumberDecomposition(request *calPb.PrimeNumbeRequest, stream calPb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Println("PrimeNumberDecomposition server")
	n := int(request.Int1)

	k := 2
	for n > 1 {
		if n%k == 0 {
			res := &calPb.PrimeNumberResponse{
				Res: int32(k),
			}
			err := stream.Send(res)

			if err != nil {
				log.Fatalln("error sending response from PrimeNumberDecomposition RPC", err)
			}
			time.Sleep(1000 * time.Millisecond)
			n /= k
		} else {
			k++
		}
	}
	return nil
}

func (s *server) Sum(ctx context.Context, req *calPb.Input) (*calPb.Result, error) {
	fmt.Println("Request received", req)
	input1, input2 := req.In1, req.In2
	res := &calPb.Result{
		Res: input1 + input2,
	}
	return res, nil
}

func (s *server) ComputeAverage(stream calPb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("Inside ComputeAverage")
	requests := make([]int32, 0)
	for {
		input, err := stream.Recv()
		if err == io.EOF {
			res := 0.0
			sum := 0
			for _, v := range requests {
				sum += int(v)
			}
			res = float64(sum) / float64(len(requests))
			return stream.SendAndClose(&calPb.ComputeAverageResponse{
				Res: float32(res),
			})
		}
		if err != nil {
			log.Fatalf("error while receiving request from Client %v\n", err)
		}
		requests = append(requests, input.GetInput())
	}
}

func (s *server) FindMaximum(stream calPb.CalculatorService_FindMaximumServer) error {
	maxSoFar := -1
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln("error while receiving request from client ", err)
			return err
		}
		temp := req.GetInput()
		if temp > int32(maxSoFar) {
			maxSoFar = int(temp)
			err := stream.Send(&calPb.FindMaximumResponse{
				Res: temp,
			})
			if err != nil {
				log.Fatalln("error sending respone to the client ", err)
			}
		}
	}
	return nil

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
