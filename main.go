package main

import (
	"example.com/m/controllers"
	"example.com/m/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.POST("/users", controllers.UsersCreate)
	r.PUT("/users/:id", controllers.UsersUpdate)
	r.GET("/users", controllers.UsersIndex)
	r.GET("/users/:id", controllers.UsersShow)
	r.DELETE("/users/:id", controllers.UsersDelete)
	r.Run()
}
