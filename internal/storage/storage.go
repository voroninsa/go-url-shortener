package storage

import (
	"github.com/voroninsa/go-url-shortener/internal/storage/base"
	"github.com/voroninsa/go-url-shortener/internal/storage/inmemory"
	"github.com/voroninsa/go-url-shortener/internal/storage/postgres"
)

var AppStorage base.IStorage

func NewStorage(storageType string) base.IStorage {
	if AppStorage != nil {
		return AppStorage
	}
	switch storageType {
	case "sql":
		AppStorage = postgres.NewSqlStorage()
		return AppStorage
	case "imnemory":
		AppStorage = inmemory.NewInMemoryStorage()
		return AppStorage
	}
	return nil
}
