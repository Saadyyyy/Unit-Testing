package main

import (
	"github.com/Saadyyyy/Unit-Testing/apps/config"
	"github.com/Saadyyyy/Unit-Testing/apps/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	db := config.InitializeDatabase()
	var ctx *gin.Context

	routes.Api(r, db, ctx)
	// routes.ImagesRouter(r, db)
	r.Run()
}
