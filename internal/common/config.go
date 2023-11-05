package common

import (
	"flag"
	"fmt"
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
		return fmt.Errorf("invalid storage type: %s", storgageType)
	}

	CFG.StorageType = storgageType
	return nil
}
