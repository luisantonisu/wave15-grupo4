package main

import (
	"log"

	"github.com/luisantonisu/wave15-grupo4/cmd/server"
	"github.com/luisantonisu/wave15-grupo4/internal/config"
)

func main() {
	// env
	// ...
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("error loading config: %v", err)
		return
	}
	// app
	log.Println("Starting server on " + cfg.ServerAddress)
	// - config
	app := server.NewServerChi(cfg)
	// - run
	if err := app.Run(*cfg); err != nil {
		log.Println(err)
		return
	}
}
