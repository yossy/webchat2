package main

import (
	"./controller"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"./data"
)

func main() {
	data.DbInit()
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("SessionName", store))
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.TweetIndex)
	router.GET("/signup", controller.UserSignupForm)
	router.POST("/signup", controller.UserSignup)
	router.GET("/signin", controller.UserSigninForm)
	router.POST("/signin", controller.UserSignin)
	router.POST("/signout", controller.UserSignOut)
	router.GET("/user/mypage/:id", controller.UserMypage)
	router.GET("/new", controller.TweetNew)
	router.POST("/new", controller.TweetPost)
	router.GET("/show/:id", controller.TweetShow)
	router.GET("/edit/:id", controller.TweetEdit)
	router.POST("/edit/:id", controller.TweetUpdate)
	router.POST("/delete/:id", controller.TweetDestroy)
	router.POST("/tweet/:id/comments", controller.CommentNew)
	router.Run()
}
