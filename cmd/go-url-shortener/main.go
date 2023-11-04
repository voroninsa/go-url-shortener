package main

import (
	"fmt"

	"github.com/voroninsa/go-url-shortener/internal/storage"
)

func main() {
	// storage := newInMemoryStorage()
	storage := storage.NewStorage("SQL")
	// storage.PutUrl("Yutube.com", "weugnk")

	fmt.Println(storage.GetUrl("ddfassdetu"))

}
