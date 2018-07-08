package fs

import (
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin"
	rcloneFS "github.com/ncw/rclone/fs"
	"github.com/ncw/rclone/fs/object"
)

func put(fs rcloneFS.Fs, file *multipart.FileHeader) error {
	f, err := file.Open()
	if err != nil {
		panic(err)
	}
	dst := file.Filename
	objInfo := object.NewStaticObjectInfo(dst,
		time.Now(), file.Size, false, nil, fs)

	_, err = fs.Put(f, objInfo)
	if err != nil {
		panic(err)
	}
	return nil

}

func Put(c *gin.Context) {
	fsrc, err := rcloneFS.NewFs(remote)
	if err != nil {
		panic(err)
	}

	form, _ := c.MultipartForm()
	files, _ := form.File["file[]"]
	for _, file := range files {
		_ = put(fsrc, file)
	}

	c.JSON(200, gin.H{
		"status": "OK",
	})
}
