package controllers

import (
	"example.com/m/initializers"
	"example.com/m/models"
	"github.com/gin-gonic/gin"
)

func UsersCreate(c *gin.Context) {

	var body struct {
		Username string
		Password string
	}

	c.Bind(&body)

	user := models.User{Username: body.Username, Password: body.Password}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersIndex(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"user": users,
	})
}

func UsersShow(c *gin.Context) {
	id := c.Param("id")

	var users []models.User
	initializers.DB.First(&users, id)

	c.JSON(200, gin.H{
		"user": users,
	})
}

func UsersUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Username string
		Password string
	}

	c.Bind(&body)

	var user models.User
	initializers.DB.First(&user, id)

	initializers.DB.Model(&user).Updates(models.User{
		Username: body.Username,
		Password: body.Password,
	})

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.User{}, id)

	c.Status(200)
}
