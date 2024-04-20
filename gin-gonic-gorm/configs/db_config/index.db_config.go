package db_config

import "os"

var DB_DRIVER = "mysql"
var DB_HOST = "127.0.0.1" //localhost
var DB_PORT = "3306"
var DB_NAME = "gin_gonic_gorm"
var DB_USER = "root"
var DB_PASSWORD = ""

func InitDatabaseConfig() {
	driverEnv := os.Getenv("DB_DRIVER")

	if driverEnv != "" {
		DB_DRIVER = driverEnv
	}

	hostEnv := os.Getenv("DB_HOST")

	if hostEnv != "" {
		DB_HOST = hostEnv
	}
	portEnv := os.Getenv("DB_PORT")

	if portEnv != "" {
		DB_PORT = portEnv
	}
	nameEnv := os.Getenv("DB_NAME")

	if nameEnv != "" {
		DB_NAME = nameEnv
	}
	usernameEnv := os.Getenv("DB_USER")

	if usernameEnv != "" {
		DB_USER = usernameEnv
	}
	passwordEnv := os.Getenv("DB_PASSWORD")

	if passwordEnv != "" {
		DB_PASSWORD = passwordEnv
	}
}
