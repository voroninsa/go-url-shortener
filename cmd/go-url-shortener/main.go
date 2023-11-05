package main

import (
	"fmt"
	"net/http"

	"github.com/voroninsa/go-url-shortener/internal/api"
	"github.com/voroninsa/go-url-shortener/internal/common"
)

func main() {
	if err := common.InitConfig(); err != nil {
		fmt.Println(err)
		return
	}
	// common.InitConfig(storgageType)
	// common.InitConfig(*storgageType)
	// if err := storage.PutUrl(common.EncodeString("vk.com"), "vk.com"); err != nil {
	// 	panic(err)
	// }
	// storage.PutUrl(common.EncodeString("yahoo.com"), "yahoo.com")

	server := api.NewServer()
	// server.Router.GET("/:shortenUrl", getHandler)
	// router.POST("/shorten-url", postHandler)

	http.ListenAndServe(":8080", server)
}
