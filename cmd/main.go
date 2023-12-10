package main

import (
	"github.com/caarlos0/env/v6"
	"learn"
	"learn/internal/config"
	"learn/internal/email"
	"learn/internal/handler"
	"learn/internal/service"
	"log"
)

func main() {
	cfg := &config.Config{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("error due parse config %v", err)
	}
	client := email.NewClient(cfg)

	s := service.NewSerivce(client)

	h := handler.NewHandler(s)

	srv := new(learn.Server)

	log.Println("Running server")

	if err := srv.Run("8080", h.Init()); err != nil {
		log.Println(err)
	}

}
