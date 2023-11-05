package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/voroninsa/go-url-shortener/internal/api"
	"github.com/voroninsa/go-url-shortener/internal/common"
)

func main() {
	var storgageType string
	flag.StringVar(&storgageType, "s", "InMemory", "Storage type to use ('SQL' | 'InMemory')")

	fmt.Println(storgageType)
	// common.InitConfig(storgageType)
	common.InitConfig("InMemory")
	// if err := storage.PutUrl(common.EncodeString("vk.com"), "vk.com"); err != nil {
	// 	panic(err)
	// }
	// storage.PutUrl(common.EncodeString("yahoo.com"), "yahoo.com")

	server := api.NewServer()
	// server.Router.GET("/:shortenUrl", getHandler)
	// router.POST("/shorten-url", postHandler)

	http.ListenAndServe(":8080", server)
}
