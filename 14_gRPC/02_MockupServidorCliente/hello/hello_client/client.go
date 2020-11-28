package main

import (
	"fmt"

	"../hellopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client is running!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.fatalf("Failed to connect %v", err)
	}

	defer cc.Close()

	c := hellopb.NewHelloServiceClient(cc)
}
