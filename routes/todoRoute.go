package routes

import (
	"todo-app-go/controllers"
	"todo-app-go/middlewares"

	"github.com/gin-gonic/gin"
)

func TodoRoute(incomingRoutes *gin.Engine) {
	todo := incomingRoutes.Group("/todo")
	todo.POST("/", middlewares.CheckCredentials, controllers.Create)
	todo.GET("/", middlewares.CheckCredentials, controllers.ShowTodo)
	todo.PUT("/:id", middlewares.CheckCredentials, controllers.UpdateTodo)
	todo.DELETE("/", middlewares.CheckCredentials, controllers.DeleteTodo)
}