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

