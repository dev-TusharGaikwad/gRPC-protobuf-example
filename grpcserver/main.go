package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "grpc-api-service/proto"
)

const (
	port = ":50051"
)

// UnimplementedGreetServer can be embedded to have forward compatible implementations.
type myServer struct {
	pb.UnimplementedGreetServer
}

var obj1 pb.Resp

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
	obj := &pb.Person{}
	obj.Name = "Ojas"
	obj.Age = 22
	obj.Text = "this is sample text"

	obj1.Reply = obj.Name + " " + strconv.Itoa(int(obj.Age)) + " " + obj.Text
}

func (c *myServer) Call(ctx context.Context, in *pb.EmptyMsg) (*pb.Resp, error) {
	return &pb.Resp{Reply: obj1.Reply}, nil
}
