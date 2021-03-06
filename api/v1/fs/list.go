package fs

import (
	"strings"

	"github.com/gin-gonic/gin"
	rcloneFS "github.com/ncw/rclone/fs"
	"github.com/ncw/rclone/fs/walk"
)

func list(fs rcloneFS.Fs, path string) FSObjects {

	// rclone gives error for '/'
	path = strings.TrimSuffix(path, "/")

	dt, err := walk.NewDirTree(fs, path, true, 1)
	if err != nil {
		panic(err)
	}

	var fsobjs FSObjects
	for _, v := range dt {
		for _, f := range v {
			fsobjs = append(fsobjs, toFSObject(f))
		}
	}
	return fsobjs
}

func List(c *gin.Context) {
	fsrc, err := rcloneFS.NewFs(remote)
	if err != nil {
		panic(err)
	}

	path := c.DefaultQuery("path", "")
	files := list(fsrc, path)
	c.JSON(200, gin.H{
		"status":     "OK",
		"fs_objects": files,
	})
}
