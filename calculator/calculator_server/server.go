package main

import (
	"context"
	"fmt"
	"log"
	"net"
	calcualtorpb "swiggy/grpc/calculator/calculatorpb"

	grpc "google.golang.org/grpc"
)

type server struct {
	calcualtorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calcualtorpb.SumRequest) (*calcualtorpb.SumResponse, error) {
	fmt.Printf("Greet Function was invoked with %v", req)
	firstNum := req.FirstNumber
	secondNum := req.SecondNumber
	result := firstNum + secondNum
	res := calcualtorpb.SumResponse{
		SumResult: result,
	}
	return &res, nil
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()
	calcualtorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
