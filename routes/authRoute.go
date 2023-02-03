package routes

import (
	"todo-app-go/controllers"
	"todo-app-go/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoute(incomingRoutes *gin.Engine) {
	user := incomingRoutes.Group("/user")
	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)
	user.GET("/", middlewares.CheckCredentials, controllers.Profile)
	user.PUT("/", middlewares.CheckCredentials, controllers.ChangePass)
}