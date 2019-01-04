package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	vmiss "github.com/illush123/hanamizuki/server/vmiss/handler"
)

func Register(e *gin.Engine) {
	v1 := e.Group("/api/v1")

	// vmiss
	{
		r := v1.Group("/vmiss")
		r.POST("/upload", vmiss.StoreFiles)
	}

	// health checker
	e.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
}
