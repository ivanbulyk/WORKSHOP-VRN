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

	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
