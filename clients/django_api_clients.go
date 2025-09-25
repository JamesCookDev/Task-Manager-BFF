package clients

import (
	"context"
	"log"
	"os"
	"time"

	pb "task-manager-bff/proto" 

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GetProjetos agora retorna uma mensagem Protobuf em vez de bytes
func GetProjetos() (*pb.ListarProjetosResponse, error) {
	// Pega o endereço do servidor gRPC do ambiente
	djangoGrpcUrl := os.Getenv("DJANGO_GRPC_URL")
	if djangoGrpcUrl == "" {
		log.Fatal("Variável de ambiente DJANGO_GRPC_URL não definida")
	}

	// Estabelece uma conexão com o servidor gRPC.
	// `insecure` é usado para desenvolvimento; em produção, usaríamos TLS.
	conn, err := grpc.Dial(djangoGrpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Cria um "stub" de cliente a partir da conexão
	c := pb.NewProjetoServiceClient(conn)

	// Define um timeout para a chamada
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Faz a chamada RPC!
	res, err := c.ListarProjetos(ctx, &pb.ListarProjetosRequest{})
	if err != nil {
		return nil, err
	}

	return res, nil
}