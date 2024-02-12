package routes

import (
	"github.com/Saadyyyy/Unit-Testing/api/handler"
	"github.com/Saadyyyy/Unit-Testing/api/repository"
	"github.com/Saadyyyy/Unit-Testing/api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Api(r *gin.Engine, db *gorm.DB, ctx *gin.Context) {
	repository := repository.NewMakananRepository(db)
	service := service.NewMakananService(repository)
	controller := handler.NewMakananController(service, ctx)

	makanan := r.Group("makan")
	{
		makanan.POST("/", controller.Create)
	}
}
