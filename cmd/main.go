package main

import (
	"fmt"
	"log"

	"github.com/luisantonisu/wave15-grupo4/cmd/server"
	"github.com/luisantonisu/wave15-grupo4/internal/config"
)

func main() {
	// env
	// ...
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Errorf("error loading config: %v", err)

	}
	// app
	log.Printf("config: %v", cfg)
	log.Println("Starting server on :" + cfg.ServerAddress)
	// - config
	app := server.NewServerChi(cfg)
	// - run
	if err := app.Run(*cfg); err != nil {
		fmt.Println(err)
		return
	}
}
