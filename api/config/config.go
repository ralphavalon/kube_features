package config

import "os"

var Config = struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBPort     string
	DBName     string
}{}

func init() {
	Config.DBHost = os.Getenv("DB_HOST")
	Config.DBUser = os.Getenv("DB_USER")
	Config.DBPassword = os.Getenv("DB_PASSWORD")
	Config.DBPort = os.Getenv("DB_PORT")
	Config.DBName = os.Getenv("DB_NAME")
}
