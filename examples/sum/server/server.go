package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/go-grpc/examples/sum/pb"
	"google.golang.org/grpc"
)

// server to implement sum service
type server struct {
	pb.UnimplementedSumServiceServer
}

func (*server) Sum(ctx context.Context, req *pb.SumReq) (*pb.SumResp, error) {
	num1 := req.GetNums().GetNum1()
	num2 := req.GetNums().GetNum2()

	result := num1 + num2

	return &pb.SumResp{Result: result}, nil
}

func main() {
	fmt.Println("Starting server")

	// Listening to tcp network on address localhost with port 60001
	lis, err := net.Listen("tcp", "0.0.0.0:60001")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// create new grpc service
	s := grpc.NewServer()
	// register sum service with grpc
	pb.RegisterSumServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to listen to the server %v", err)
	}

	log.Println("Connected to the server")

}
