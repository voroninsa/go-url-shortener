package configs

import (
	"flag"
	"fmt"
	"strings"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	StorageType string `envconfig:"default=inmemory"`
	DBName      string `envconfig:"default=short_urls,postgres_db"`
	DBPort      string `envconfig:"default=5432,database_port"`
	DBHost      string `envconfig:"default=postgres_container,database_host"`
	DBUserName  string `envconfig:"default=root,postgres_user"`
	DBPassword  string `envconfig:"postgres_password"`
}

var CFG Config

func InitConfig() error {
	flagString := flag.String("s", "InMemory", "Storage type to use ('SQL' | 'InMemory')")
	flag.Parse()
	storgageType := strings.ToLower(*flagString)

	if storgageType != "sql" && storgageType != "inmemory" {
		return fmt.Errorf("invalid storage type: %s", storgageType)
	}

	envconfig.Init(&CFG)

	CFG.StorageType = storgageType

	fmt.Println("Initialized config: ", CFG)

	return nil
}
