package main

import (
	"log"
	"net/http"
	"io"

	"github.com/go-chi/chi/v5"
)

const djangoAPI = "http://localhost:8000/api/v1/projects/"

const apiToken = ""


func getProjetosHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", djangoAPI, nil)
	if err != nil {
		http.Error(w, "Erro ao criar requisição para a API interna", http.StatusInternalServerError)
		log.Println("Erro NewRequest:", err)
		return
	}

	req.Header.Add("Authorization", "Bearer " + apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Erro ao fazer requisição para a API interna", http.StatusInternalServerError)
		log.Println("Erro client.Do:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Erro ao ler a resposta da API interna", http.StatusInternalServerError)
		log.Println("Erro io.ReadAll:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func main() {
	router := chi.NewRouter()
	router.Get("/api/projetos", getProjetosHandler)

	port := ":8080"
	log.Printf("Servidor BFF iniciado na porta %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}


