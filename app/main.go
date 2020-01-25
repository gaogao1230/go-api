package main

import (
	"app/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"strconv"
)

func getDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/user/:id", func(c *gin.Context) {
		userId, _ := strconv.Atoi(c.Param("id"))
		db := getDB()
		var user model.User
		if err := db.Where("id = ?", userId).First(&user).Error; gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    "Not Found",
				"message": "レコードが見つかりません",
			})
			return
		}
		c.JSON(http.StatusOK, user)
		defer db.Close()
	})
	return r
}

var router = Router()

func main() {
	r := router
	r.Run(":8080")
}
