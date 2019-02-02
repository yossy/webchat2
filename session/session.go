package session

import (
  "github.com/gin-gonic/contrib/sessions"
  "github.com/gin-gonic/gin"
  "../data"
)

type SessionInfo struct {
  UserID          interface{}
  UNickName       interface{}
  IsSessionAlive  bool
}

func Login(c *gin.Context, user data.User) {
  session := sessions.Default(c)
  session.Set("alive", true)
  session.Set("userID", user.ID)
  session.Set("nickname", user.NickName)
  session.Save()
}

func ClearSession(c *gin.Context) {
  session := sessions.Default(c)
  session.Clear()
  session.Save()
}

func GetSessionInfo(c *gin.Context) SessionInfo {
  var info SessionInfo
  session := sessions.Default(c)
  user_id := session.Get("userID")
  nickname := session.Get("nickname")
  alive := session.Get("alive")
  if user_id == nil && nickname == nil && alive == nil {
    info = SessionInfo {
      UserID: -1, UNickName:"", IsSessionAlive: false,
    }
  } else {
      info = SessionInfo {
        UserID: user_id.(uint),
        UNickName: nickname.(string),
        IsSessionAlive: alive.(bool),
      }
  }
  return info
}
