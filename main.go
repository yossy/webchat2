package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
	"./data"
	)

func main() {
	data.DbConnect()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.IndexRouter)
	router.GET("/new", controller.NewRouter)
	router.POST("/new", controller.PostRouter)
	router.GET("/show/:id", controller.ShowRouter)
	router.Run()
}
