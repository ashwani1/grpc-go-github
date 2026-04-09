package main

import (
	"context"
	"log"
	"net"

	// pb "grpc-go-github/proto"
	pb "github.com/ashwani1/grpc-go-github/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello " + req.Name,
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Println("Server running on :50051")
	s.Serve(lis)
}
