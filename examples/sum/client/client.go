package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-grpc/examples/sum/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Client")

	conn, err := grpc.Dial("0.0.0.0:60001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to dial client %v", err)
	}

	c := pb.NewSumServiceClient(conn)

	doUnary(c)
}

func doUnary(c pb.SumServiceClient) {

	req := pb.SumReq{
		Nums: &pb.Number{
			Num1: 23,
			Num2: 49,
		},
	}

	res, err := c.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("Unable to do sum due to %v", err)
	}

	log.Println("Result of addition is ", res.Result)
}
