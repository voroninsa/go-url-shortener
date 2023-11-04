package main

import (
	"fmt"

	"github.com/voroninsa/go-url-shortener/internal/storage"
)

func main() {
	// storage := newInMemoryStorage()
	storage := storage.NewStorage("SQL")
	// storage.PutUrl("Yutube.com", "weugnk")

	if err := storage.PutUrl("dsafdsgfdh", "vk.com"); err != nil {
		fmt.Println("ERROR: ", err)
	}

	// fmt.Println(storage.GetUrl("ddfassdetu"))

}
