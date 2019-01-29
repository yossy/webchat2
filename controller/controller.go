package controller

import (
	"github.com/gin-gonic/gin"
  "../data"
)


func TweetIndex(c *gin.Context) {
  tweet := data.GetAll()
	c.HTML(200, "index.html", gin.H{
    "tweet": tweet,
    })
}

func UserSignupForm(c *gin.Context) {
  c.HTML(200, "user_registration.html", nil)
}

func UserSignup(c *gin.Context) {
  nickname := c.PostForm("nickname")
  email := c.PostForm("email")
  password := c.PostForm("password")
  passwordhash := data.PasswordHash(password)
  data.UserCreate(nickname, email, passwordhash)
  c.Redirect(302, "/")
}

func TweetNew(c *gin.Context) {
  c.HTML(200, "new.html", nil)
}

func TweetPost(c *gin.Context) {
  text := c.PostForm("text")
  image := c.PostForm("image")
  data.TweetCreate(text, image)
  c.Redirect(302, "/")
}

func TweetShow(c *gin.Context) {
  id := c.Param("id")
  tweet := data.TweetFind(id)
  c.HTML(200, "show.html", gin.H{
    "tweet": tweet,
    })
}
