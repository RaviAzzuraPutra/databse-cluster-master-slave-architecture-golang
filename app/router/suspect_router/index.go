package suspect_router

import (
	"databse-cluster-master-slave-architecture-golang/app/controller/suspect_controller"

	"github.com/gin-gonic/gin"
)

func SuspectRouter(app *gin.Engine, SuspectController *suspect_controller.Suspect_Controller) {

	suspect := app.Group("/api/suspect")

	suspect.POST("/create/:id_case", SuspectController.Create)
	suspect.GET("/get-all/:id_case", SuspectController.GetAll)
	suspect.GET("/get-id/:id/:id_case", SuspectController.GetById)
	suspect.PUT("/update/:id/:id_case", SuspectController.Update)
	suspect.DELETE("/delete/:id/:id_case", SuspectController.Delete)

}
