package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type products struct {
	gorm.Model
	Code  string
	Name  string
	Price uint
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	dbInit()
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main Website",
			"info":  "",
		})
	})

	router.GET("/search", func(c *gin.Context) {
		var keyword string
		keyword = c.Query("keyword")
		var info string
		info = read(keyword)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main Website",
			"info":  info,
		})
	})
	router.Run(":8080")
}

func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&products{})

	//create
	db.Create(&products{Code: "L0001", Name: "Apple", Price: 1000})
	db.Create(&products{Code: "L0002", Name: "Orange", Price: 100})
	db.Create(&products{Code: "L0003", Name: "Banana", Price: 2000})
	db.Create(&products{Code: "L0004", Name: "Papaya", Price: 5000})
	db.Create(&products{Code: "L0005", Name: "Mango", Price: 3000})
}

func read(keyword string) string {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("failed to connect database")
	}

	//read
	var product products
	db.First(&product, "Name = ?", keyword)
	var info string
	info += "Name: " + fmt.Sprint(product.Name)
	info += " Price: " + fmt.Sprint(product.Price)
	return info
}
