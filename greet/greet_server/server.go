package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/vikasmanohar/grpc-go/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet funtion invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("GreetManyTimes function invoked with request", req)
	firstName := req.GetGreeting().FirstName
	for i := 0; i < 10; i++ {
		res := &greetpb.GreetManyTimesResponse{
			Result: "Hello " + firstName + " " + strconv.Itoa(i),
		}
		err := stream.Send(res)

		if err != nil {
			log.Fatalln("error sending response from GreetManyTimes RPC", err)
		}
		//adding a delay of one second per response so that each response can be seen clearly
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (s *server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("LongGreet server function invoked with streaming request")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//client has sent all the requests
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error with client stream %v\n", err)
		}
		result += "Hello " + req.GetGreeting().GetFirstName() + "! "
	}
}

func main() {
	fmt.Println("I am server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("error creating tcp connection %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}

}
