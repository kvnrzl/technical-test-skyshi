package router

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	data := []byte(time.Now().String())
	etag := fmt.Sprintf("W/%x", md5.Sum(data))

	return func(c *gin.Context) {
		c.Header("X-Powered-By", "Express")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		c.Header("Connection", "keep-alive")
		c.Header("Keep-Alive", "timeout=5")
		c.Header("ETag", etag)

		if match := c.GetHeader("If-None-Match"); match != "" {
			if strings.Contains(match, etag) {
				c.Status(http.StatusNotModified)
				return
			}
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}

}
