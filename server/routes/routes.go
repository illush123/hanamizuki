package routes

import (
	"fmt"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func Aaa(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "aaa.html", gin.H{})
}

func TestForm(c *gin.Context) {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields.", file.Filename))
}