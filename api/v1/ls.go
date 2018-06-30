package fs

import (
    "bytes"
    "os"
    "github.com/ncw/rclone/fs/operations"
    "github.com/ncw/rclone/cmd"
_    "github.com/ncw/rclone/backend/s3"
)

/*
type Manager interface {
    Copy func(src string, dst string)
    List func(src string)
    Move func(src string, dst string)
}
*/


func List(path []string) {
    fsrc := cmd.NewFsSrc(path)
    operations.List(fsrc, os.Stdout)
}

func main() {
    path := []string{"spaces:dagobah"}
    List(path)
}
