package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"vietanh.com/gohero/controllers"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("templates/*")
	router.GET("/comments", controllers.GetComments)
	router.GET("/comments/:id", controllers.GetComment)
	router.POST("/comments", controllers.CreateComment)
	router.PUT("/comments/:id", controllers.UpdateComment)
}
