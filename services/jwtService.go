package services

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtServices struct{
	secretKey string
}

func JWTAuthService() JWTService{
	return &jwtServices{
		secretKey: getEnv(),
	}
}

func getEnv() string{
	secret := os.Getenv("SECRET")
	if secret == ""{
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}