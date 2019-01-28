package data

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
  gorm.Model
  Name string
  Email string
  Password string
}

type Tweet struct {
  gorm.Model
  Text string
  Image string
  UserID int
}

func DbConnect() {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
      panic("failed to connect database")
    }
    db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Tweet{})
    if db.Error != nil {
      panic(db.Error)
    }
    defer db.Close()
    db.LogMode(true)
}

func GetAll() []Tweet {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  var tweet []Tweet
  db.Find(&tweet)
  defer db.Close()
  return tweet
}

func Create(text string, image string) {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  db.Create(&Tweet{Text: text, Image: image})
  defer db.Close()
}

func TweetFind(id string) Tweet {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  var tweet Tweet
  db.Where("id = ?", id).First(&tweet)
  defer db.Close()
  return tweet
}
