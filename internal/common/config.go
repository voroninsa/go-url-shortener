package common

import (
	"errors"
	"flag"
	"strings"
)

type Config struct {
	StorageType string
	DBName      string
	DBPort      int16
	DBHost      string
	DBUserName  string
	DBPassword  string
}

var CFG Config

func InitConfig() error {
	flagString := flag.String("s", "InMemory", "Storage type to use ('SQL' | 'InMemory')")
	flag.Parse()
	storgageType := strings.ToLower(*flagString)

	if storgageType != "sql" && storgageType != "inmemory" {
		return errors.New("invalid storage type")
	}

	CFG.StorageType = storgageType
	return nil
}
