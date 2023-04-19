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
		users.DELETE("/:id", UserApi.DeleteEndpoint)                       //删除用户
		users.GET("", UserApi.AllEndpoint)                                 //获取所有用户
		users.GET("/paging", UserApi.PagingEndpoint)                       //分页查询用户
		users.PUT("/:id", UserApi.UpdateEndpoint)                          //编辑用户
		users.PATCH("/:id/status", UserApi.UpdateStatusEndpoint)           //修改用户状态
		users.POST("/:id/change-password", UserApi.ChangePasswordEndpoint) //修改用户密码
		users.GET("/:id", UserApi.DetailsEndpoint)                         //用户详情

	}

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8282")
}
