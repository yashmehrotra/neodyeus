package fs

import (
    "time"
    _ "github.com/ncw/rclone/backend/all"
)

type File struct {
    Name string `json:"name"`
    Size int64 `json:"size"`
    LastModified time.Time `json:"last_modified"`
    FullPath string `json:"full_path"`
}

type Files []File

// TODO: Write a transforming function
