package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JavierDominguezGomez/Course-Go/14_gRPC/02_MockupServidorCliente/recordatorio/recordatoriopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Go client is running!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	defer cc.Close()

	c := recordatoriopb.NewRecordatorioServiceClient(cc)

	recordatorioUnary(c)
}

func recordatorioUnary(c recordatoriopb.RecordatorioServiceClient) {
	fmt.Println("Starting unary RPC Recordatorio.")

	req := &recordatoriopb.RecordatorioRequest{
		Recordatorio: &recordatoriopb.Recordatorio{
			Id:          1,
			Mensaje:     "Te recuerdo que has quedado en llamar por teléfono",
			FechaLimite: "2020-12-01",
		},
	}

	res, err := c.Recordatorio(context.Background(), req)

	if err != nil {
		log.Fatalf("Error, calling Recordatorio RPC: \n%v", err)
	}

	log.Printf("Response Recordatorio: %v", res.CustomRecordatorio)
}
