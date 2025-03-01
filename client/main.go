package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/anterabyte/gRPC-tutorial/proto"
)

func init() {

	fmt.Println("Intiating the Connection")

}

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect to the port, %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	name := "world"

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {

		log.Fatalln("Could not say hello", err)

	}

	log.Printf("Getting response from server, %v", r.GetMessage())
}
