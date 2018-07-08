package fs

import (
	"log"

	"github.com/gin-gonic/gin"
	rcloneFS "github.com/ncw/rclone/fs"
	"github.com/ncw/rclone/fs/operations"
)

type RemoveRequest struct {
	Path   string `json:"path"`
	DryRun bool   `json:"dry_run"`
}

func remove(fs rcloneFS.Fs, path string, dryRun bool) (FSObjects, bool) {

	fsObjects := list(fs, path)
	if dryRun {
		return fsObjects, false
	}

	err := operations.Purge(fs, path)
	if err != nil {
		log.Print(err)
		return fsObjects, false
	}
	return fsObjects, true
}

func Remove(c *gin.Context) {
	fsrc, err := rcloneFS.NewFs(remote)
	if err != nil {
		panic(err)
	}

	var req RemoveRequest
	c.BindJSON(&req)

	fsObjects, removed := remove(fsrc, req.Path, req.DryRun)
	c.JSON(200, gin.H{
		"status":     "OK",
		"removed":    removed,
		"fs_objects": fsObjects,
	})
}
