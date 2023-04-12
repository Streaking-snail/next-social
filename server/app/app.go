package app

import (
	"next-social/server/api"

	"github.com/gin-gonic/gin"
)

func Run() {

	//r := gin.Default()
	r := gin.New()

	r.POST("/login", api.Login)

	users := r.Group("/users")
	{
		users.POST("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		// users.GET("", UserApi.AllEndpoint)
		// users.GET("/paging", UserApi.PagingEndpoint)
		// users.POST("", UserApi.CreateEndpoint)
		// users.PUT("/:id", UserApi.UpdateEndpoint)
		// users.PATCH("/:id/status", UserApi.UpdateStatusEndpoint)
		// users.DELETE("/:id", UserApi.DeleteEndpoint)
		// users.GET("/:id", UserApi.GetEndpoint)
		// users.POST("/:id/change-password", UserApi.ChangePasswordEndpoint)
		// users.POST("/:id/reset-totp", UserApi.ResetTotpEndpoint)
	}

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8282")
}

func Admin() {

}
