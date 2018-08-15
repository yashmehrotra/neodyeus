package fs

import (
	"log"

	"github.com/gin-gonic/gin"
	rcloneFS "github.com/ncw/rclone/fs"
)

type MkdirRequest struct {
	Path string `json:"path"`
}

func mkdir(fs rcloneFS.Fs, path string) bool {

	err := fs.Mkdir(path)
	if err != nil {
		log.Print(err)
		return false
	}
	return true
}

func Mkdir(c *gin.Context) {
	fsrc, err := rcloneFS.NewFs(remote)
	if err != nil {
		panic(err)
	}

	var req MkdirRequest
	c.BindJSON(&req)

	created := mkdir(fsrc, req.Path)
	c.JSON(200, gin.H{
		"status":  "OK",
		"created": created,
	})
}
