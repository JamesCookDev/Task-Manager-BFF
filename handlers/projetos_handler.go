package handlers

import (
	"task-manager-bff/clients"
	"log"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson" 
)

func GetProjetosHandler(w http.ResponseWriter, r *http.Request) {
	// Chama nosso novo cliente gRPC
	grpcResponse, err := clients.GetProjetos()
	if err != nil {
		log.Printf("Erro ao chamar o cliente gRPC: %v", err)
		http.Error(w, "Erro ao processar sua requisição", http.StatusInternalServerError)
		return
	}

	// Converte a resposta Protobuf para JSON
	jsonBytes, err := protojson.Marshal(grpcResponse)
	if err != nil {
		log.Printf("Erro ao converter resposta para JSON: %v", err)
		http.Error(w, "Erro ao processar resposta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}