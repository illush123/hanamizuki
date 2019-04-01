package main

import (
	"github.com/gin-gonic/gin"
	"github.com/illush123/hanamizuki/server/config"
	"github.com/illush123/hanamizuki/server/internal/executor"
	"github.com/illush123/hanamizuki/server/router"
)

func init() {
	config.Init()
	executor.Init()
}

func main() {
	api := gin.Default()
	router.Register(api)

	api.Run(":8888")
}
