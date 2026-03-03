package bootstrap

import (
	"databse-cluster-master-slave-architecture-golang/app/config"
	"databse-cluster-master-slave-architecture-golang/app/config/app_config"
	"databse-cluster-master-slave-architecture-golang/app/database"
	"databse-cluster-master-slave-architecture-golang/app/registry/cases_registry"
	"databse-cluster-master-slave-architecture-golang/app/registry/suspect_registry"
	"databse-cluster-master-slave-architecture-golang/app/router/cases_router"
	"databse-cluster-master-slave-architecture-golang/app/router/suspect_router"

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

	CasesModule := cases_registry.Case_Registry()
	SuspectModule := suspect_registry.Suspect_Registry()

	cases_router.CasesRouter(app, CasesModule.Cases_Controller)
	suspect_router.SuspectRouter(app, SuspectModule.Suspect_Controller)

	app.Run(app_config.PORT)
}
