package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yashmehrotra/neodyeus/api/v1/fs"
)

func main() {
	r := gin.Default()

	r.GET("/api/v1/list", fs.List)
	r.Run()
}
