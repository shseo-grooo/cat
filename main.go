package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().Unix())
	meow := []string{"야옹~", "애옹", "냥", "샤하아아아악", "골골골골", "냐!"}

	r := gin.Default()

	r.GET("/meow", func(c *gin.Context) {

		mode := c.Request.Header["X-Mode"]
		mode = append(mode, "")

		c.String(http.StatusOK, "[Cat Server v4 - %s] %s", mode[0], meow[rand.Intn(len(meow))])
	})

	r.Run()
}
