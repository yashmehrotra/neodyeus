package fs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	rcloneFS "github.com/ncw/rclone/fs"
)

func Preview(c *gin.Context) {
	fsrc, err := rcloneFS.NewFs(remote)
	if err != nil {
		panic(err)
	}
	path := c.Param("path")
	w, fsObj, err := get(fsrc, path)
	if err != nil {
		panic(err)
	}
	http.ServeContent(c.Writer, c.Request, fsObj.Name, fsObj.LastModified, w)
	return
}
