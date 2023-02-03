package middlewares

import (
	"net/http"
	"time"
	"todo-app-go/initializers"
	"todo-app-go/models"
	"todo-app-go/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckCredentials(c *gin.Context){
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, _ := services.JWTAuthService().ValidateToken(tokenString)

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		//Check the Exp 
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
		}

		// Find the user with token sub
		var user models.User

		initializers.DB.First(&user, claims["sub"])

		if user.Id == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
		}

		//Attach to req
		c.Set("user", user)

		// Continue
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	}
}