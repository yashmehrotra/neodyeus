package fs

import (
    //"fmt"
    "github.com/ncw/rclone/fs/operations"
    "github.com/ncw/rclone/cmd"
    rcloneFS "github.com/ncw/rclone/fs"
    "github.com/gin-gonic/gin"
)

func listrclone(path []string) Files {

    fileChan := make(chan rcloneFS.Object, 100)
    Rand := func (o rcloneFS.Object) {
        fileChan <- o
    }
    fsrc := cmd.NewFsSrc(path)
    var files Files
    go func() {
        for f := range fileChan {
            files = append(files, File{
                Name: f.Remote(),
                Size: f.Size(),
                LastModified: f.ModTime(),
                FullPath: f.Remote(),
            })
        }
    }()
    operations.ListFn(fsrc, Rand)
    close(fileChan)
    return files
}

func List(c *gin.Context) {
    path := []string{"/home/yash/rclone_test/"}
    files := listrclone(path)
    c.JSON(200, gin.H{
        "status": "OK",
        "files": files,
    })
}
