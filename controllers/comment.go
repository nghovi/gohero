package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vietanh.com/gohero/models"
)

var db = func() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/story_presenter?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("DB error", err)
	}
	return db
}()

func GetComments(c *gin.Context) {
	var comments []models.Comment
	result := db.Where("content LIKE ?", "%%").Order("updated_at desc").Find(&comments)
	if result.Error == nil {
		c.JSON(200, gin.H{
			"message":  "Success",
			"comments": &comments,
		})
	} else {
		c.JSON(200, gin.H{"error": result.Error})
	}
}

func GetComment(c *gin.Context) {
	var story models.Comment
	result := db.First(&story, c.Param("id"))
	if result.Error == nil {
		c.JSON(200, gin.H{
			"message": "Success",
			"story":   &story,
		})
	} else {
		c.JSON(200, gin.H{"error": result.Error})
	}
}

func CreateComment(c *gin.Context) {
	var comment models.Comment
	c.Bind(&comment)
	comment.Ip = c.ClientIP()
	comment.Agent = c.Request.UserAgent()
	result := db.Create(&comment) // pass pointer of data to Create
	if result.Error == nil {
		c.JSON(200, gin.H{
			"message": "Success",
			"comment": comment,
		})
	} else {
		c.JSON(200, gin.H{"error": result.Error})
	}
}

func UpdateComment(c *gin.Context) {
	result := db.Model(&models.Comment{}).Where("id = ?", c.Param("id")).Updates(models.Comment{
		Content: c.PostForm("content"),
		Ip:      c.ClientIP(),
		Agent:   c.Request.UserAgent(),
	})
	if result.Error == nil {
		c.Redirect(http.StatusFound, "/comments")
		// c.JSON(200, gin.H{
		// 	"message": "Success",
		// })
	} else {
		c.JSON(200, gin.H{"error": result.Error})
	}
}

//No reply your own comment
func ReplyComment(c *gin.Context) {
	var parentComment models.Comment
	err := db.First(&parentComment, c.PostForm("parent_id")).Error
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
	}
	var reply models.Comment
	c.Bind(&reply)
	reply.Ip = c.ClientIP()
	reply.Agent = c.Request.UserAgent()
	result := db.Create(&reply) // pass pointer of data to Create
	if result.Error == nil {
		c.JSON(200, gin.H{
			"message": "Success",
			"comment": reply,
		})
	} else {
		c.JSON(200, gin.H{"error": result.Error})
	}
}
