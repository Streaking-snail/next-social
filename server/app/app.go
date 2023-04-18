package app

import (
	"next-social/server/api"
	mw "next-social/server/app/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {

	//r := gin.Default()
	r := gin.New()
	r.Use(mw.Auth)

	UserApi := new(api.UserApi)

	r.POST("/login", api.Login)
	r.POST("/users", UserApi.CreateEndpoint)

	users := r.Group("/users", mw.Admin)
	{
		users.DELETE("/:id", UserApi.DeleteEndpoint)
		users.GET("", UserApi.AllEndpoint)
		users.GET("/paging", UserApi.PagingEndpoint)
		users.PUT("/:id", UserApi.UpdateEndpoint)
		// users.PATCH("/:id/status", UserApi.UpdateStatusEndpoint)
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
