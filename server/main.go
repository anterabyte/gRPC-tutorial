package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/anterabyte/gRPC-tutorial/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	name := req.GetName()
	message := fmt.Sprintf("Hello, %s!", name)
	return &pb.HelloResponse{Message: message}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {

		log.Fatalf("Failed to listen, %v",err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf(`Server listening at %v`, lis.Addr())
	if err := s.Serve(lis)  ; err != nil {

		log.Fatalf("Failed to serve, %v",err)
	}
}
