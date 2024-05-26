package main

import (
	"github.com/Guesstrain/website/controller"
	"github.com/Guesstrain/website/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	databaseService database.DatabaseService = database.NewDatabaseService()

	overviewController controller.OverviewController = controller.New(databaseService)
)

func main() {
	server := gin.New()

	server.Use(cors.Default())

	server.Use(gin.Recovery(), gin.Logger())

	apiRoutes := server.Group("api")
	{
		apiRoutes.GET("/overview", func(ctx *gin.Context) {
			ctx.JSON(200, overviewController.FindAll())
		})
	}

	server.Run(":8080")
}
