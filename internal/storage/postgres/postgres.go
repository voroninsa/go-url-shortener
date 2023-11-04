package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/voroninsa/go-url-shortener/internal/storage/base"
)

type SQLStorage struct {
	StorageURL *sql.DB
}

func (s *SQLStorage) GetUrl(shortUrl string) (string, error) {
	var url string

	if err := s.StorageURL.QueryRow(`SELECT url FROM urls WHERE short_url = $1`, shortUrl).Scan(&url); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("No such record: ", shortUrl)
		}
		return "", fmt.Errorf("Error DataBase: ", err)
	}

	return url, nil
}

func (s *SQLStorage) PutUrl(url string, short_url string) bool {
	// s.StorageMap[short_url] = url
	return true
}

func NewSqlStorage() base.IStorage {
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlconn := "host=localhost port=5432 user=root password=qwerty dbname=short_urls sslmode=disable"
	db, _ := sql.Open("postgres", psqlconn)

	return &SQLStorage{
		StorageURL: db,
	}
}
