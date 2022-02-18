package main

import (
	"fmt"
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
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET")

		mode := append(c.Request.Header["X-Mode"], "active")[0]

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("[Cat Server v6.4 - %s] %s", mode, meow[rand.Intn(len(meow))])})
	})

	r.Run()
}
