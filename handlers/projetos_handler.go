package handlers

import (
	"task-manager-bff/client"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetProjetosHandler(w http.ResponseWriter, r *http.Request) {
	body, statusCode, err := clients.GetProjetos()
	if err != nil {
		http.Error(w, "Erro", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}

func CreateProjetoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	respBody, statusCode, err := clients.CreateProjeto(body)
	if err != nil {
		http.Error(w, "Erro", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(respBody)
}

func GetProjetoByIDHandler(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")
	body, statusCode, err := clients.GetProjetoByID(projectID)
	if err != nil {
		http.Error(w, "Erro", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}

func UpdateProjetoHandler(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	respBody, statusCode, err := clients.UpdateProjeto(projectID, body)
	if err != nil {
		http.Error(w, "Erro", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(respBody)
}

func DeleteProjetoHandler(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "projectID")
	statusCode, err := clients.DeleteProjeto(projectID)
	if err != nil {
		log.Printf("Erro ao deletar projeto: %v", err)
		http.Error(w, "Erro", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
}