package main

import (
	"context"
	"fmt"
	"log"
	calcualtorpb "swiggy/grpc/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("coul not connet %f", err)
	}
	defer cc.Close()
	c := calcualtorpb.NewCalculatorServiceClient(cc)
	doUnary(c)
}

func doUnary(c calcualtorpb.CalculatorServiceClient) {
	req := &calcualtorpb.SumRequest{
		FirstNumber:  15,
		SecondNumber: 15,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling greet RPC %v", err)
	}
	log.Printf("Sum: %v", res.SumResult)
}
