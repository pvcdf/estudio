package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"1_3_11/pb"
)

type servidorHumano struct {
	pb.UnimplementedMaquinaHumanaServer
}

func (s *servidorHumano) EvaluarHumano(ctx context.Context, req *pb.Humano) (*pb.ResultadoHumano, error) {
	log.Printf("se ha recibido a: ", req.Nombre, req.Edad, req.Iq)

	valido := false
	mensaje := "no se puede vro"

	if req.Edad >= 18 && req.Iq > 90 {
		valido = true
		mensaje = "humano listo pa la guerra"
	}
	if req.Edad < 18 && req.Iq > 90 {
		valido = false
		mensaje = "tas muy chiquito, rechazado"
	}
	if req.Edad >= 18 && req.Iq < 90 {
		valido = true
		mensaje = "tas muy tonto, rechazado"
	}
	return &pb.ResultadoHumano{
		Valido:  valido,
		Mensaje: mensaje,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("error puerto: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterMaquinaHumanaServer(grpcServer, &servidorHumano{})

	log.Println("iniciado")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("error sv: %v", err)
	}
}
