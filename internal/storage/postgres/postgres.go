package postgres

import (
	"database/sql"
	"errors"
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
			return "", errors.New(fmt.Sprint("No such record: ", shortUrl))
		}
		return "", errors.New(fmt.Sprint("Error DataBase: ", err))
	}

	return url, nil
}

func (s *SQLStorage) PutUrl(shortUrl string, url string) error {
	db_url, err := s.GetUrl(shortUrl)
	// Check if record exists. Don't check errors in Get
	if db_url != "" {
		return nil
	}

	_, err = s.StorageURL.Exec(`INSERT INTO urls (short_url, url) VALUES ($1, $2)`, shortUrl, url)
	if err != nil {
		return errors.New(fmt.Sprint("Adding data error: ", err))
	}

	return nil
}

func NewSqlStorage() base.IStorage {
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlconn := "host=localhost port=5432 user=root password=qwerty dbname=short_urls sslmode=disable"
	db, _ := sql.Open("postgres", psqlconn)

	return &SQLStorage{
		StorageURL: db,
	}
}
