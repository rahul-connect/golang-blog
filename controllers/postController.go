package controllers

import (
	"go_blog/inits"
	"go_blog/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	var body struct {
		Title  string
		Body   string
		Likes  int
		Draft  bool
		Author string
		UserID uint `json:"user_id"`
	}

	ctx.BindJSON(&body)

	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(500, gin.H{"error": "user not found"})
		return
	}

	body.UserID = user.(models.User).ID

	post := models.Post{
		Title:  body.Title,
		Body:   body.Body,
		Likes:  body.Likes,
		Draft:  body.Draft,
		Author: body.Author,
		UserID: body.UserID,
	}

	result := inits.DB.Create(&post)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	ctx.JSON(200, gin.H{"data": post})
}

func GetPosts(ctx *gin.Context) {
	var posts []models.Post

	result := inits.DB.Find(&posts)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": posts})
}

func GetPost(ctx *gin.Context) {
	var post models.Post

	result := inits.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	ctx.JSON(200, gin.H{"data": post})
}

func UpdatePost(ctx *gin.Context) {
	var body struct {
		Title  string
		Body   string
		Likes  int
		Draft  bool
		Author string
	}

	ctx.BindJSON(&body)

	var post models.Post

	result := inits.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	inits.DB.Model(&post).Updates(models.Post{
		Title:  body.Title,
		Body:   body.Body,
		Likes:  body.Likes,
		Draft:  body.Draft,
		Author: body.Author,
	})

	ctx.JSON(200, gin.H{"data": post})
}

func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")

	inits.DB.Delete(&models.Post{}, id)
	ctx.JSON(200, gin.H{"data": "post has been deleted successsfully"})

}
