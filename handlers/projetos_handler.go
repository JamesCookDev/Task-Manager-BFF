package handlers

import (
	"log"
	"net/http"

	"task-manager-bff/clients"  
)

func GetProjetosHandler(w http.ResponseWriter, r *http.Request) {
	body, statusCode, err := clients.GetProjetos()
	if err != nil {
		log.Printf("Erro ao chamar o cliente da API Django: %v", err)
		http.Error(w, "Erro ao processar sua requisição", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(body)
}