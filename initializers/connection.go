package initializers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"todo-app-go/models"
)

var DB *gorm.DB

func ConnectToDb(){
	dsn := os.Getenv("DB")
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Todo{})

	DB = db
}