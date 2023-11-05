package common

type Config struct {
	StorageType string
	DBName      string
	DBPort      int16
	DBHost      string
	DBUserName  string
	DBPassword  string
}

var CFG Config

func InitConfig(s string) {
	CFG.StorageType = s
}
