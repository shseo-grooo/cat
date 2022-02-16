package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/meow", func(c *gin.Context) {
		mode := c.Request.Header["X-Mode"]
		mode = append(mode, "")
		c.String(http.StatusOK, "[Cat Server v3 - %s] 야옹", mode[0])
	})

	r.Run()
}
