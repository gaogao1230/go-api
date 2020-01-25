package main

import (
	"app/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "user:password@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.Create(&model.User{Id: 1, Name: "山田", Age: 11})
	db.Create(&model.User{Id: 2, Name: "田中", Age: 12})
}
