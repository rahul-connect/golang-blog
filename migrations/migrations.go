package main

import (
	"go_blog/inits"
	"go_blog/models"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	inits.DB.AutoMigrate(&models.Post{})
	inits.DB.AutoMigrate(&models.User{})
}
