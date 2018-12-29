package main

import (
	"github.com/illush123/hanamizuki/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/", routes.Home)
	router.GET("/aaa", routes.Aaa)

	router.POST("/aaa", routes.TestForm)
	router.Run(":8080")
}
