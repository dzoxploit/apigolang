package main

import (
	"apigolang/controllers"
	"apigolang/database"
	"apigolang/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB("root:@tcp(localhost:3306)/task_management?parseTime=true")
	database.Migrate()

	router := initRouter()
	router.Run(":8000")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateJWTToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/dashboard", controllers.Ping)
			secured.GET("/tasks", controllers.FindTasks)
			secured.POST("/tasks", controllers.CreateTask)
			secured.GET("/tasks/one", controllers.FindTask) 
			secured.PUT("/tasks/update", controllers.UpdateTask)
			secured.POST("/tasks/delete", controllers.DeleteTask)  

			secured.GET("/users", controllers.GetAllUsers)
			secured.POST("/users", controllers.CreateUser)
			secured.GET("/users/one", controllers.GetUserByID) 
			secured.PUT("/users/update", controllers.UpdateUser)
			secured.POST("/users/delete", controllers.DeleteUser)  
			
		}
	}
	return router
}
