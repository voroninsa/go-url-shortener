package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/voroninsa/go-url-shortener/configs"
	"github.com/voroninsa/go-url-shortener/internal/common"
	"github.com/voroninsa/go-url-shortener/internal/storage"
	"github.com/voroninsa/go-url-shortener/internal/storage/base"
)

type Url struct {
	Url string //`json:"url"`
}

type IUrlController interface {
	GetUrl() (string, error)
	CreateShortUrl() (string, error)
}

type UrlController struct {
	storageDriver base.IStorage
	request       *http.Request
}

// Return URL saved in storage
func (u *UrlController) GetUrl() (string, error) {
	short_url := mux.Vars(u.request)["short_url"]
	url, err := u.storageDriver.GetUrl(short_url)

	if err != nil {
		return "", err
	}
	return url, nil
}

// Encode URL, save in storage and return encoded short URL
func (u *UrlController) CreateShortUrl() (string, error) {
	var decoded_url Url
	decoder := json.NewDecoder(u.request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&decoded_url)
	if err != nil {
		return "", fmt.Errorf("JSON Decode error: %s", err)
	}

	short_url := common.EncodeString(decoded_url.Url)

	err = u.storageDriver.PutUrl(short_url, decoded_url.Url)
	if err != nil {
		return "", fmt.Errorf("creating Short URL error: %s", err)
	}
	return short_url, nil
}

func InstantiateController(r *http.Request) IUrlController {
	fmt.Println("config: ", configs.CFG)
	st := configs.CFG.StorageType
	storage := storage.NewStorage(st)

	return &UrlController{
		storageDriver: storage,
		request:       r,
	}
}
