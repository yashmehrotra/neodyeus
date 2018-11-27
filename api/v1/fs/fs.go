package fs

import (
	"os"
	"path/filepath"
	"time"

	humanize "github.com/dustin/go-humanize"
	_ "github.com/ncw/rclone/backend/local"
	rcloneFS "github.com/ncw/rclone/fs"
)

var (
	remote = os.Getenv("NEODYEUS_PATH")
)

type ObjectMeta struct {
	Thumbnail   string `json:"thumbnail"`
	ContentType string `json:"content_type"`
}

type FSObject struct {
	Name         string     `json:"name"`
	Size         int64      `json:"size"`
	HumanSize    string     `json:"human_size"`
	LastModified time.Time  `json:"last_modified"`
	FullPath     string     `json:"full_path"`
	IsDir        bool       `json:"is_dir"`
	Meta         ObjectMeta `json:"meta"`
}

type FSObjects []FSObject

func extractContentType(path string) string {
	contentTypeMap := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
	}

	if val, exists := contentTypeMap[filepath.Ext(path)]; exists {
		return val
	}
	return ""
}

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

	// For directories, do not count size of the tree yet
	var humanSize string
	if !isDir {
		humanSize = humanize.Bytes(uint64(f.Size()))
	}

	objMeta := ObjectMeta{
		Thumbnail:   "",
		ContentType: extractContentType(name),
	}

	return FSObject{
		Name:         name,
		Size:         f.Size(),
		HumanSize:    humanSize,
		LastModified: f.ModTime(),
		FullPath:     fullPath,
		IsDir:        isDir,
		Meta:         objMeta,
	}
}
