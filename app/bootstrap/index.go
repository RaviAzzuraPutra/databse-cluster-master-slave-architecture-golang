package bootstrap

import (
	"databse-cluster-master-slave-architecture-golang/app/config"
	"databse-cluster-master-slave-architecture-golang/app/config/app_config"
	"databse-cluster-master-slave-architecture-golang/app/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitAPP() {
	_ = godotenv.Load()

	config.Config()
	database.Connect()

	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"Message": "The application is running well. 💮",
		})
	})

	app.Run(app_config.PORT)
}
