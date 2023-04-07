package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name      string `gorm: "type:varchar(20);not null"`
	Telephone string `gorm: "type: varchar(11); not null; unique"`
	Password  string `gorm: "size: 255; not null"`
}

func main() {
	db := InitDB()

	defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		tel := c.PostForm("telephone")

		if len(tel) != 11 {
			// c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{"code": 422, "msg": "手机号必须为11位"})
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"code": http.StatusUnprocessableEntity, "msg": "密码不得小于6位"})
			return
		}
		if len(username) == 0 {
			username = RandomString(10)
		}
		log.Println(username, tel, password)
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "注册成功！"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func RandomString(n int) string {
	letters := []byte("qwertyuiopasdfghjklzxcvbnm")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *sql.DB {
	// host := "localhost"
	// port := 3306
	// database := "ginessential"
	// username := "root"
	// password := "root"
	// charset := "utf8"
	// args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql)

	db, err := sql.Open("mysql", "root:root@/golang")
	// db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db

}
