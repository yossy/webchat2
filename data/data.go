package data

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
  gorm.Model
  NickName string
  Email string
  Password string
}

type Tweet struct {
  gorm.Model
  Text string
  Image string
  UserID int
}

func DbInit() {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Tweet{})
  if db.Error != nil {
    panic(db.Error)
  }
  db.LogMode(true)
}

func GetAll() []Tweet {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  var tweet []Tweet
  db.Find(&tweet)
  return tweet
}

func TweetCreate(text string, image string) {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  db.Create(&Tweet{Text: text, Image: image})
}

func TweetFind(id string) Tweet {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  var tweet Tweet
  db.Where("id = ?", id).First(&tweet)
  return tweet
}

func PasswordHash(password string) string {
  passwordhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    panic(err)
  }
  return string(passwordhash)
}

func FindLoginUser(email string, password string) User {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  var user User
  db.Where(&User{Email: email}).First(&user)
  passwordhashchecked := user.Password
  missmatch := bcrypt.CompareHashAndPassword([]byte(passwordhashchecked), []byte(password))
  if missmatch != nil {
    panic("failed to check password")
  }
  return user
}

func UserCreate(nickname string, email string, passwordhash string) User {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  var user User
  db.Create(&User{NickName: nickname, Email: email, Password: passwordhash})
  db.Where(&User{Email: email}).First(&user)
  return user
}
