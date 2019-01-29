package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
	"./data"
	)

func main() {
	data.DbInit()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.TweetIndex)
	router.GET("/signup", controller.UserSignupForm)
	router.POST("/signup", controller.UserSignup)
	router.GET("/new", controller.TweetNew)
	router.POST("/new", controller.TweetPost)
	router.GET("/show/:id", controller.TweetShow)
	router.Run()
}
