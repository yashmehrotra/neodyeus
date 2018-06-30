package fs

import (
	_ "github.com/ncw/rclone/backend/all"
	rcloneFS "github.com/ncw/rclone/fs"
	"path/filepath"
	"time"
)

type FSObject struct {
	Name         string    `json:"name"`
	Size         int64     `json:"size"`
	LastModified time.Time `json:"last_modified"`
	FullPath     string    `json:"full_path"`
	IsDir        bool      `json:"is_dir"`
}

type FSObjects []FSObject

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
