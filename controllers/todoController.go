package controllers

import (
	"encoding/json"
	"net/http"
	"todo-app-go/initializers"
	"todo-app-go/models"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context){
	var body struct{
		Title string
		Desc string
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Membaca Inputan",
		})
		return
	}

	user, _ := c.Get("user")

	id := user.(models.User).Id

	todo := models.Todo{Title: body.Title, Desc: body.Desc, IsComplete: false, UserId: id}
	result := initializers.DB.Create(&todo)
	if result.Error != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal Membuat Todo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil Membuat Todo",
	})
}

func ShowTodo(c *gin.Context){
	var todo []models.Todo

	user, _ := c.Get("user")

	id := user.(models.User).Id

	initializers.DB.Where("user_id = ?", id).Preload("User").Find(&todo)

	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
}

func UpdateTodo(c *gin.Context){
	var body struct{
		Title string
		Desc string
	}
	var todo models.Todo
	id := c.Param("id")
	user, _ := c.Get("user")
	idUser := user.(models.User).Id

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Gagal Membaca Inputan",
		})
		return
	}

	if initializers.DB.Model(&todo).Where("id = ?", id).Where("user_id = ?", idUser).Updates(map[string]interface{}{"title": body.Title, "desc": body.Desc}).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Data Tidak Ditemukkan", 
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data Berhasil Di Update",
	})
}

func DeleteTodo(c *gin.Context){
	var todo models.Todo
	
	var input struct{
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, _ := input.Id.Int64()
	user, _ := c.Get("user")
	idUser := user.(models.User).Id
	if initializers.DB.Where("user_id = ?", idUser).Delete(&todo, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Data Tidak Ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data Berhasil Dihapus",
	})
}