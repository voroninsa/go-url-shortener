package main

import (
	"log"
	"net/http"

	"github.com/voroninsa/go-url-shortener/configs"
	"github.com/voroninsa/go-url-shortener/internal/api"
)

func main() {
	configs.InitConfig()
	// if err := configs.InitConfig(); err != nil {
	// 	log.Fatal(err)
	// }
	// common.InitConfig(storgageType)
	// common.InitConfig(*storgageType)
	// if err := storage.PutUrl(common.EncodeString("vk.com"), "vk.com"); err != nil {
	// 	panic(err)
	// }
	// storage.PutUrl(common.EncodeString("yahoo.com"), "yahoo.com")

	server := api.NewServer()
	// server.Router.GET("/:shortenUrl", getHandler)
	// router.POST("/shorten-url", postHandler)

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatal(err)
	}
}
