package controllers

import (
	"net/http"
	"os"
	"time"
	"todo-app-go/initializers"
	"todo-app-go/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context){
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Membaca Inputan",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Membuat Token",
		})
		return
	}
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal Membuat User",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Registrasi Akun",
	})
}

func Login(c *gin.Context){
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Membaca Inputan",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.Id == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email atau Password salah",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email atau Password salah",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed Create Token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Profile(c *gin.Context){
	user, _ := c.Get("user")

	email := user.(models.User).Email

	c.JSON(http.StatusOK, gin.H{
		"email": email,
	})
}

func ChangePass(c *gin.Context){
	var body struct{
		Password string 
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var account models.User
	user, _ := c.Get("user")

	email := user.(models.User).Email

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	if initializers.DB.Model(&account).Where("email = ?", email).UpdateColumn("password", hash).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Data Tidak Ditemukkan", 
		})
		return
	}
	

	c.JSON(http.StatusOK, gin.H{
		"message": "Password Berhasil Di Ganti",
	})
}