package fs

import (
	_ "github.com/ncw/rclone/backend/all"
	rcloneFS "github.com/ncw/rclone/fs"
	"path/filepath"
	"time"
)

const (
	remote = "/home/yash/rclone_test"
)

type FSObject struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	LastModified time.Time `json:"last_modified"`
	FullPath     string    `json:"full_path"`
	IsDir        bool      `json:"is_dir"`
}

type FSObjects []FSObject

// TODO
// If the need be, this can also be used,
// Also, replace cmd.NewFsSrc with this
// fsWithPath := rcloneFS.NewFs(remote + ":" + path)

// toFSObject converts rclone's
func toFSObject(f rcloneFS.DirEntry) FSObject {
	isDir := false
	if f.Size() < 0 {
		isDir = true
	}
	fullPath := f.Remote()
	name := filepath.Base(fullPath)

	return FSObject{
		Name:         name,
		Size:         f.Size(),
		LastModified: f.ModTime(),
		FullPath:     fullPath,
		IsDir:        isDir,
	}
}
