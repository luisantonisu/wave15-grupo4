package main

import (
	"fmt"

	"github.com/luisantonisu/wave15-grupo4/cmd/server"
)

func main() {
	// env
	// ...

	// app
	// - config
	cfg := &server.ConfigServerChi{
		ServerAddress:  ":8080",
		LoaderFilePath: "infrastructure/json/mock.json", // todo
	}
	app := server.NewServerChi(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
