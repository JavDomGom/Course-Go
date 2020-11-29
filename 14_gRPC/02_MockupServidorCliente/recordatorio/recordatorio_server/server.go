package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/JavierDominguezGomez/Course-Go/14_gRPC/02_MockupServidorCliente/recordatorio/recordatoriopb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Recordatorio(ctx context.Context, req *recordatoriopb.RecordatorioRequest) (*recordatoriopb.RecordatorioResponse, error) {
	fmt.Printf("Recordatorio function was invoked with %v\n", req)

	id := string(req.GetRecordatorio().GetId())
	mensaje := req.GetRecordatorio().GetMensaje()
	fechaLimite := req.GetRecordatorio().GetFechaLimite()

	customRecordatorio := "Recordatorio con id " + id + ", mensaje: " + mensaje + ", fecha_limite: " + fechaLimite

	res := &recordatoriopb.RecordatorioResponse{
		CustomRecordatorio: customRecordatorio,
	}
	return res, nil
}

func main() {
	fmt.Println("Recordatorio, Go server is running!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	recordatoriopb.RegisterRecordatorioServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
