package vmiss

import (
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/illush123/hanamizuki/proto"
	"github.com/illush123/hanamizuki/server/internal/executor"
)

func StoreFiles(c *gin.Context) {
	f, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	key := c.Query("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not found folder key"})
		return
	}
	req, err := buildRequest(f.File[key])
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	buildDir := c.Query("build_dir")
	if buildDir == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not found build dir"})
		return
	}
	req.BuildDir = buildDir

	_, err = executor.StoreFiles(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, req)
}

func buildRequest(f []*multipart.FileHeader) (*proto.VmissRequest, error) {
	req := &proto.VmissRequest{
		Files: []*proto.CodeFile{},
	}

	for _, file := range f {
		mf, err := file.Open()
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(mf)
		if err != nil {
			return nil, err
		}
		codeFile := &proto.CodeFile{
			FileName: file.Filename,
			Body:     b,
		}
		req.Files = append(req.Files, codeFile)
	}
	return req, nil
}
