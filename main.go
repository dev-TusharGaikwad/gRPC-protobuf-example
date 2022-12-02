package main

import (
	"fmt"

	pb "github.com/golang/protobuf/proto"

	proto "grpc-api-service/proto"
)

func main() {
	dataObj := &proto.Person{}
	dataObj.Age = 22
	dataObj.Name = "xyz"

	obj, _ := pb.Marshal(dataObj)

	fmt.Println(obj)
}
