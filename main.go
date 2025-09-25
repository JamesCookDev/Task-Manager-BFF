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

	r := chi.NewRouter()

	r.Route("/api/projetos", func(r chi.Router) {
		r.Get("/", handlers.GetProjetosHandler)        // Listar
		r.Post("/", handlers.CreateProjetoHandler)       // Criar
		r.Get("/{projectID}", handlers.GetProjetoByIDHandler) // Obter por ID
		r.Put("/{projectID}", handlers.UpdateProjetoHandler)  // Atualizar
		r.Delete("/{projectID}", handlers.DeleteProjetoHandler) // Deletar
	})

	port := ":8080"
	log.Printf("Servidor BFF iniciado na porta %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}