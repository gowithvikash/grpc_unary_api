package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/gowithvikash/grpc_with_go/grpc_unary_api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	do_Simple_Greet(c)

}

func do_Simple_Greet(c pb.GreetServiceClient) {
	fmt.Println("\n_______________ Action Number : 01 _______________")
	fmt.Println("_____  do_Simple_Greet Function Was Invoked At Client  ____")
	res, err := c.Simple_Greet(context.Background(), &pb.GreetRequest{Name: "Vikash Parashar"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("___ do_Simple_Greet_Result: %v\n", res.Result)

}
