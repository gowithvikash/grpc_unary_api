package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/gowithvikash/grpc_with_go/grpc_unary_api/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var (
	network = "tcp"
	address = "0.0.0.0:50051"
)

func main() {
	lis, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
func (s *Server) Simple_Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Println("___ Simple_Greet Function Was Invoked At Client___")
	return &pb.GreetResponse{Result: fmt.Sprintf("Hello And Welcome %s !\n", in.Name)}, nil
}
