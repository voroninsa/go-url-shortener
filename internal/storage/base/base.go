package base

type IStorage interface {
	GetUrl(shortUrl string) (string, error)
	PutUrl(url, shortUrl string) bool
}
