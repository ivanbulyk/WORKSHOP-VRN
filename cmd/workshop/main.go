package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ivanbulyk/WORKSHOP-VRN/internal/api/jokes"
	"github.com/ivanbulyk/WORKSHOP-VRN/internal/config"
	"github.com/ivanbulyk/WORKSHOP-VRN/internal/handler"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	h := handler.NewHandler(apiClient)

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	path := cfg.Host + ":" + cfg.Port

	srv := &http.Server{
		Addr:    path,
		Handler: r,
	}

	// handle shutdown gracefully
	quit := make(chan os.Signal, 1)
	done := make(chan error, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		err := srv.Shutdown(ctx)
		//....
		done <- err
	}()

	log.Printf("Starting server at %s", path)
	_ = srv.ListenAndServe()

	err = <-done
	log.Printf("Shutting server down with %v", err)
}
