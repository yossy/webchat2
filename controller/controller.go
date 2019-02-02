package controller

import (
	"github.com/gin-gonic/gin"
  "../data"
  "../session"
)

func TweetIndex(c *gin.Context) {
  info := session.GetSessionInfo(c)
  tweet := data.GetAll()
	c.HTML(200, "index.html", gin.H{
    "tweet": tweet, "SessionInfo": info,
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
  user := data.UserCreate(nickname, email, passwordhash)
  session.Login(c, user)
  c.Redirect(302, "/")
}

func UserSigninForm(c *gin.Context) {
  c.HTML(200, "user_signin.html", nil)
}

func UserSignin(c *gin.Context) {
  email := c.PostForm("email")
  password := c.PostForm("password")
  user := data.FindLoginUser(email, password)
  session.Login(c, user)
  c.Redirect(302, "/")
}

func UserSignOut(c *gin.Context) {
  session.ClearSession(c)
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
