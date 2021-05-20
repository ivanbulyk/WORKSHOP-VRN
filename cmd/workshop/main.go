package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ivanbulyk/WORKSHOP-VRN/internal/config"
	"github.com/ivanbulyk/WORKSHOP-VRN/internal/handler"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	path := cfg.Host + ":" + cfg.Port

	log.Printf("Starting server at %s", path)
	err = http.ListenAndServe(path, r)
	log.Fatal(err)

	log.Print("Shutting server down...")
}
