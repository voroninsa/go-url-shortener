package main

import (
	"log"
	"net/http"

	"github.com/voroninsa/go-url-shortener/configs"
	"github.com/voroninsa/go-url-shortener/internal/api"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		log.Fatal(err)
	}

	server := api.NewServer()

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err)
	}
}
