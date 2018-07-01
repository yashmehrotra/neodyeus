package fs

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ncw/rclone/cmd"
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
	remote := []string{"/home/yash/rclone_test/"}
	fsrc := cmd.NewFsSrc(remote)

	var req MkdirRequest
	c.BindJSON(&req)

	created := mkdir(fsrc, req.Path)
	c.JSON(200, gin.H{
		"status":  "OK",
		"created": created,
	})
}
