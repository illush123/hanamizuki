package handler

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func StoreFiles(c *gin.Context) {
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
