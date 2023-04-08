package app

import (
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8282")
}
