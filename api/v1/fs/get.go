package fs

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	rcloneFS "github.com/ncw/rclone/fs"
)

// TODO, handle errors, don't just panic
func get(fs rcloneFS.Fs, path string) (*bytes.Reader, FSObject, error) {
	fsObj, err := fs.NewObject(path)
	if err != nil {
		panic(err)
	}

	fp, err := fsObj.Open()
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	b := make([]byte, 0)
	w := bytes.NewBuffer(b)
	_, err = io.Copy(w, fp)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(w.Bytes())
	return r, toFSObject(fsObj), nil
}

func Get(c *gin.Context) {
	fsrc, err := rcloneFS.NewFs(remote)
	if err != nil {
		panic(err)
	}
	path := c.DefaultQuery("path", "")
	w, fsObj, err := get(fsrc, path)
	if err != nil {
		panic(err)
	}
	http.ServeContent(c.Writer, c.Request, fsObj.Name, fsObj.LastModified, w)
	return
}
