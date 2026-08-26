package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"1_3_11/pb"
)

func main() {
	miMensaje := &pb.Humano{
		Nombre: "Benjamin Rojas",
		Edad:   21,
		Iq:     100,
	}

	fmt.Println(miMensaje)

	conexion, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer conexion.Close()

	cliente := pb.NewMaquinaHumanaClient(conexion)

	respuesta, err := cliente.EvaluarHumano(context.Background(), miMensaje)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("es un humano valido?", respuesta.Valido)
}
