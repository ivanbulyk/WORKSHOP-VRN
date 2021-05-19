package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ivanbulyk/WORKSHOP-VRN/internal/handler"
)

func main() {
	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Print("Starting server at :8080")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)

	log.Print("Shutting server down...")
}
