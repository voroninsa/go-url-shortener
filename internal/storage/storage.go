package storage

import (
	"github.com/voroninsa/go-url-shortener/internal/storage/base"
	"github.com/voroninsa/go-url-shortener/internal/storage/inmemory"
	"github.com/voroninsa/go-url-shortener/internal/storage/postgres"
)

func NewStorage(storageType string) base.IStorage {
	switch storageType {
	case "SQL":
		return postgres.NewSqlStorage()
	case "InMemory":
		return inmemory.NewInMemoryStorage()
	}
	return nil
}
