package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		time := time.Now()
		path := c.Request.URL.Path
		verb := c.Request.Method

		// Process request
		c.Next()

		var size int
		if c.Writer != nil {
			size = c.Writer.Size()
		}
		fmt.Printf("time: %v\npath: localhost:8080%s | verb: %s | size: %d\n", time, path, verb, size)
	}
}
