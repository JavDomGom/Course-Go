package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JavierDominguezGomez/Course-Go/14_gRPC/02_MockupServidorCliente/hello/hellopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client is running!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer cc.Close()

	c := hellopb.NewHelloServiceClient(cc)

	helloUnary(c)
}

func helloUnary(c hellopb.HelloServiceClient) {
	fmt.Println("Starting unary RPC Hello.")

	req := &hellopb.HelloRequest{
		Hello: &hellopb.Hello{
			FirstName: "Javier",
			Prefix:    "Mr.",
		},
	}

	res, err := c.Hello(context.Background(), req)

	if err != nil {
		log.Fatalf("Error, calling Hello RPC: \n%v", err)
	}

	log.Printf("Response Hello: %v", res.CustomHello)
}
