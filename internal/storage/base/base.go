package base

type IStorage interface {
	GetUrl(shortUrl string) (string, error)
	PutUrl(shortUrl, url string) error
}
