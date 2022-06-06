package main

import (
	"fmt"
	"log"

	pb "github.com/gowithvikash/grpc_with_go/grpc_unary_api/proto"
	"google.golang.org/grpc"
)

var (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	do_Simple_Greet(c)

}

func do_Simple_Greet(c pb.GreetServiceClient) {
	fmt.Println("do_Simple_Greet Function Was Invoked At Client")

}
