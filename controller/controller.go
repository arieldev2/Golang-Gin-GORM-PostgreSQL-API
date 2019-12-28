package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"gin_test/database"
	"gin_test/models"
)

func AllPost(c *gin.Context) {
	db := database.Connection()

	defer db.Close()

	var posts []models.Post
	db.Find(&posts)
	c.JSON(200, gin.H{
		"users": posts,
	})
}

func GetPost(c *gin.Context) {
	db := database.Connection()

	defer db.Close()

	var post models.Post

	id := c.Param("id")
	db.First(&post, id)
	c.JSON(200, gin.H{
		"post": post,
	})
}

func NewPost(c *gin.Context) {
	db := database.Connection()

	defer db.Close()

	var post models.Post
	c.BindJSON(&post)

	db.Create(&post)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	db := database.Connection()

	defer db.Close()

	var post models.Post

	id := c.Param("id")
	db.First(&post, id)

	c.BindJSON(&post)

	db.Save(&post)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	db := database.Connection()

	defer db.Close()

	var post models.Post

	id := c.Param("id")
	db.First(&post, id)
	db.Delete(&post)
	c.JSON(200, gin.H{
		"message": "Post Deleted",
	})
}
