package data

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "golang.org/x/crypto/bcrypt"
  "strconv"
)

type User struct {
  gorm.Model
  NickName string
  Email string
  Password string
  Tweets []Tweet
  Comments []Comment
}

type Tweet struct {
  gorm.Model
  Text string
  Image string
  UserID uint
  Comments []Comment
}

type Comment struct {
  gorm.Model
  Text string
  UserID uint
  TweetID uint64
}

func DbInit() {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Tweet{}, &Comment{})
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
  db.Order("created_at desc").Find(&tweet)
  return tweet
}

func TweetCreate(text string, image string, userId uint) {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  db.Create(&Tweet{Text: text, Image: image, UserID: userId})
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

func TweetUpdate(id string, text string, image string) {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  var tweet Tweet
  db.Where("id = ?", id).First(&tweet)
  db.Model(&tweet).Updates(Tweet{Text: text, Image: image})
}

func TweetDelete(id string) {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  db.Unscoped().Where("id = ?", id).Delete(&Tweet{})
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

func CommentCreate(id string, comment string, userId uint) {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  tweet_id, _ := strconv.ParseUint(id, 10, 0)
  db.Create(&Comment{Text: comment, TweetID: tweet_id, UserID: userId})
}

func GetComments(id string) []Comment {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  var comments []Comment
  db.Preload("Tweet").Where("tweet_id = ?", id).Find(&comments)
  return comments
}

func MyTweetFind(id string) (User, []Tweet) {
  db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  db.LogMode(true)
  var tweets []Tweet
  var user User
  db.Find(&user, id)
  db.Model(&user).Association("Tweets").Find(&tweets)
  db.Preload("User").Order("created_at desc").Where("user_id = ?", &user.ID).Find(&tweets)
  return user, tweets
}

// func UserFind(id string) User {
//   db, err := gorm.Open("mysql", "root:@/webchat2?charset=utf8&parseTime=True&loc=Local")
//   if err != nil {
//     panic("failed to connect database")
//   }
//   defer db.Close()
//   var user User
//   db.Find(&user, id).Related(&user.Tweets)
//   return user
// }
