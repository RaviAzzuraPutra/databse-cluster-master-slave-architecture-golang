package suspect_controller

import (
	"databse-cluster-master-slave-architecture-golang/app/helper"
	"databse-cluster-master-slave-architecture-golang/app/interface/service/suspect_service_interface"
	"databse-cluster-master-slave-architecture-golang/app/request/suspects_request"

	"github.com/gin-gonic/gin"
)

type Suspect_Controller struct {
	service suspect_service_interface.Suspect_Service_Interface
}

func NewSuspectControllerRegistry(suspect_service suspect_service_interface.Suspect_Service_Interface) *Suspect_Controller {
	return &Suspect_Controller{
		service: suspect_service,
	}
}

func (c *Suspect_Controller) Create(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	request := new(suspects_request.Suspects_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	input := &suspects_request.Suspects_Dto{
		Case_ID:        &ID_Case,
		ID_Card_Number: request.ID_Card_Number,
		Full_Name:      request.Full_Name,
		Address:        request.Address,
		Alibi:          request.Alibi,
	}

	suspect, errCreate := c.service.Create(ID_Case, input)

	if errCreate != nil {
		if appError, ok := errCreate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errCreate.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errCreate.Error(),
		})
	}

	ctx.JSON(201, gin.H{
		"Messaage": "Success Create Suspect",
		"Data":     suspect,
	})

}

func (c *Suspect_Controller) GetAll(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	suspect, errGet := c.service.GetAll(ID_Case)

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errGet.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errGet.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Get Suspect Data",
		"Data":    suspect,
	})

}

func (c *Suspect_Controller) GetById(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	ID := ctx.Param("id")

	suspect, errGet := c.service.GetById(ID, ID_Case)

	if errGet != nil {
		if appError, ok := errGet.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errGet.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errGet.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Get Suspect Data",
		"Data":    suspect,
	})

}

func (c *Suspect_Controller) Update(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	ID := ctx.Param("id")

	request := new(suspects_request.Suspects_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Message": "Bad Request",
			"Error":   errRequest.Error(),
		})
	}

	input := &suspects_request.Suspects_Dto{
		Case_ID:        &ID_Case,
		ID_Card_Number: request.ID_Card_Number,
		Full_Name:      request.Full_Name,
		Address:        request.Address,
		Alibi:          request.Alibi,
	}

	suspect, errUpdate := c.service.Update(ID, ID_Case, input)

	if errUpdate != nil {
		if appError, ok := errUpdate.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errUpdate.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errUpdate.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Update Suspect Data",
		"Data":    suspect,
	})

}

func (c *Suspect_Controller) Delete(ctx *gin.Context) {

	ID_Case := ctx.Param("id_case")

	ID := ctx.Param("id")

	errDelete := c.service.Delete(ID, ID_Case)

	if errDelete != nil {
		if appError, ok := errDelete.(*helper.AppError); ok {
			ctx.JSON(appError.Code, gin.H{
				"Message": appError.Message,
				"Error":   errDelete.Error(),
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Message": "Internal Server Error",
			"Error":   errDelete.Error(),
		})
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Delete Suspect Data",
	})

}
