package bootstrap

import (
	"gin-gonic-gorm/configs"
	"gin-gonic-gorm/configs/app_config"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	//load .env file
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	// INIT CONFIGS
	configs.InitConfig()

	//DATABASE CONN
	database.ConnectDatabase()

	//INIT GIN ENGINE
	app := gin.Default()

	//INIT ROUTE
	routes.InitRoute(app)

	//RUN APP
	app.Run(app_config.Port)
}
