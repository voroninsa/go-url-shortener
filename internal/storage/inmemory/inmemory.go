package inmemory

import "github.com/voroninsa/go-url-shortener/internal/storage/base"

type InMemoryStorage struct {
	StorageURL map[string]string
}

func (i *InMemoryStorage) GetUrl(shortUrl string) (string, error) {
	return i.StorageURL[shortUrl], nil
}

func (i *InMemoryStorage) PutUrl(url string, short_url string) bool {
	i.StorageURL[short_url] = url
	return true
}

func NewInMemoryStorage() base.IStorage {
	return &InMemoryStorage{
		StorageURL: make(map[string]string),
	}
}
