package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yashmehrotra/neodyeus/api/v1/fs"
)

func main() {
	r := gin.Default()

	r.GET("/api/v1/list", fs.List)
	r.PUT("/api/v1/mkdir", fs.Mkdir)
	r.DELETE("/api/v1/remove", fs.Remove)
	r.PUT("/api/v1/put", fs.Put)
	r.Run()
}
