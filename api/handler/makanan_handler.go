package handler

import (
	"net/http"

	"github.com/Saadyyyy/Unit-Testing/api/service"
	"github.com/Saadyyyy/Unit-Testing/models"
	"github.com/gin-gonic/gin"
)

type MakananController struct {
	serv service.MakananService
}

func NewMakananController(serv service.MakananService, ctx *gin.Context) *MakananController {
	return &MakananController{serv: serv}
}

func (uc *MakananController) Create(ctx *gin.Context) {
	makan := models.Makanan{}
	data, err := uc.serv.Create(makan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Massage": err,
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"massage": "Succes",
		"data":    data,
	})
}
