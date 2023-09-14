package main

import (
	"go_blog/controllers"
	"go_blog/inits"
	"go_blog/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	r.POST("/", middlewares.RequireAuth, controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)
	r.POST("/user", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/auth", controllers.Validate)
	r.GET("/users", controllers.GetUsers)
	r.GET("/logout", controllers.Logout)

	r.Run()
}
