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

		// c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// c.Header("Cache-Control", "public, max-age=31536000")
		// c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, ETag, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, client-security-token")
		// c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		c.Next()
	}

}
