package zero

import "github.com/gin-gonic/gin"

// data/+...
func GETImage(c *gin.Context) {
	path := c.Param("path")
	if path == "" {
		c.Abort()
	}
	c.File("data/" + path)
}
