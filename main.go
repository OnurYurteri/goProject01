package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

var n = 5

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/getNumber", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"n": n,
		})
	})
	return r
}

func main() {
	r := setupRouter()
	go increment(&n)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func increment(n *int) {
	for i := 0; i < 100; i++ {
		*n++
		time.Sleep(time.Millisecond * 200)
	}

}
