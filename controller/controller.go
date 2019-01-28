package controller

import (
	"github.com/gin-gonic/gin"
  "../data"
)


func IndexRouter(c *gin.Context) {
  tweet := data.GetAll()
	c.HTML(200, "index.html", gin.H{
    "tweet": tweet,
    })
}

func NewRouter(c *gin.Context) {
  c.HTML(200, "new.html", nil)
}

func PostRouter(c *gin.Context) {
  text := c.PostForm("text")
  image := c.PostForm("image")
  data.Create(text, image)
  c.Redirect(302, "/")
}

func ShowRouter(c *gin.Context) {
  id := c.Param("id")
  tweet := data.TweetFind(id)
  c.HTML(200, "show.html", gin.H{
    "tweet": tweet,
    })
}
