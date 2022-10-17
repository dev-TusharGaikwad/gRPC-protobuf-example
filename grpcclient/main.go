package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "proto-example/proto"
)

const (
	address = "localhost:50051"
)

var obj pb.Person

func main() {
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetClient(conn)

	go func() {
		for {
			var s string
			fmt.Println("\n\nPress Any key to send request to server")
			fmt.Scanln(&s)
			runGreetClient(client)
			time.Sleep(time.Second)
		}
	}()

	select {}

}

func runGreetClient(call pb.GreetClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	req := &pb.EmptyMsg{}
	res, err := call.Call(ctx, req)
	if err != nil {
		log.Fatalf("%v.Call(_) = _, %v", call, err)
	}
}
