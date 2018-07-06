package fs

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ncw/rclone/cmd"
	"github.com/ncw/rclone/fs/operations"
)

func copyRclone(srcPath, dstPath string) {
	fsrc, srcFileName, fdst := cmd.NewFsSrcFileDst([]string{srcPath, dstPath})
	println("-----" + srcFileName)
	if srcFileName == "" {
		//		sync.CopyDir(fdst, fsrc)
	} else {
		operations.CopyFile(fdst, fsrc, srcFileName, srcFileName)
	}
	return
}

func Copy(c *gin.Context) {
	srcPath := "/Users/vishesh/rclone_test/abc"
	dstPath := "/Users/vishesh/rclone_test/"
	copyRclone(srcPath, dstPath)
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
