package main

import (
	"github.com/abhaytiket/gomonorepo/lib/msg"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": msg.AddMessage("app2"),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
