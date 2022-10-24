package main

import (
	"apigolang/controllers"
	"apigolang/database"
	"apigolang/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB("root:root@tcp(localhost:3306)/task_management?parseTime=true")
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateJWTToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}