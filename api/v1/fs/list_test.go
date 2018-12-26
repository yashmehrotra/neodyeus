package fs

import (
	rcloneFS "github.com/ncw/rclone/fs"
	"testing"
)

func TestList(t *testing.T) {
	fsrc, _ := rcloneFS.NewFs(remote)
	path := ""
	ret := list(fsrc, path)
	if len(ret) < 0 {
		t.Errorf("Tests failed")
	}
}
