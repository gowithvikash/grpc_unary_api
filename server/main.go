package main

import (
	pb "github.com/gowithvikash/grpc_with_go/grpc_unary_api/proto"
)

type Server struct {
	pb.GreetServiceServer
}

var (
	network = "tcp"
	address = "0.0.0.0:50051"
)

func main() {

}
