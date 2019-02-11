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
	Config.DBHost = os.Getenv("SECRET_DB_HOST")
	Config.DBUser = os.Getenv("SECRET_DB_USER")
	Config.DBPassword = os.Getenv("SECRET_DB_PASSWORD")
	Config.DBPort = os.Getenv("SECRET_DB_PORT")
	Config.DBName = os.Getenv("SECRET_DB_NAME")
}
