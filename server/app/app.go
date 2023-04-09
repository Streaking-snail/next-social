package app

import (
	"next-social/server/api"

	"github.com/gin-gonic/gin"
)

var AuthorisedRepository = new(authorisedRepository)

type authorisedRepository struct {
	baseRepository
}

func Run() {

	//r := gin.Default()
	r := gin.New()

	r.POST("/login", api.Login)

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8282")
}
