package main

import (
	"log"
	"net/http"

	"task-manager-bff/handlers" 

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Não foi possível encontrar o arquivo .env.")
	}

	router := chi.NewRouter()

	router.Get("/api/projetos", handlers.GetProjetosHandler)

	port := ":8080"
	log.Printf("Servidor BFF iniciado na porta %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}