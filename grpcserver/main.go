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

type myServer struct {
	pb.UnimplementedGreetServer
}

var obj pb.Person

func main() {

	initServer()

	lis, err := net.Listen("tcp", port) // opening a port to listen from client
