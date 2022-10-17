package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "proto-example/proto"
)

const (
	port = ":50051"
)

// UnimplementedGreetServer can be embedded to have forward compatible implementations.
type myServer struct {
	pb.UnimplementedGreetServer
}

var obj pb.Person

func main() {

	initServer()

	lis, err := net.Listen("tcp", port) // opening a port to listen from client
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() // creating new gRPC server
	defer s.Stop()

	pb.RegisterGreetServer(s, &myServer{}) //(calls RegisterService): RegisterService registers a service and its implementation to the gRPC server.

	log.Printf("Listening at %v , server : %T ", lis.Addr(), s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func initServer() {
	log.Println()
	fmt.Println("------------------------initServer Called-------------------")

	obj.Name = "Tushar"
	obj.Age = 22
	obj.Text = "this is sample text"
}

func (c *myServer) Call(ctx context.Context, in *pb.EmptyMsg) (*pb.Resp, error) {
	res := pb.Resp{}
}
