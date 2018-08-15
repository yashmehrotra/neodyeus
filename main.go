package main

import (
	"github.com/gin-gonic/gin"

	"github.com/yashmehrotra/neodyeus/api/v1/fs"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/api/v1/list", fs.List)
	r.PUT("/api/v1/mkdir", fs.Mkdir)
	r.DELETE("/api/v1/remove", fs.Remove)
	r.PUT("/api/v1/put", fs.Put)
	r.GET("/api/v1/get", fs.Get)
	r.Run()
}
