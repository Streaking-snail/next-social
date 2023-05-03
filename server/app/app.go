package app

import (
	"next-social/server/api"
	mw "next-social/server/app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func Run() {

	//r := gin.Default()
	gin.SetMode(gin.DebugMode) //开发环境
	//gin.SetMode(gin.ReleaseMode) //线上环境
	r := gin.New()
	r.Use(mw.Auth)
	c := cron.New()

	UserApi := new(api.UserApi)
	FridApi := new(api.FridApi)
	TrendsApi := new(api.TrendsApi)
	c.AddFunc("30 0 * * *", FridApi.AutoExpireEndpoint) //好友请求过期
	c.Start()

	r.POST("/login", api.Login)
	r.POST("/users", UserApi.CreateEndpoint)

	//users := r.Group("/users", mw.Admin)
	users := r.Group("/users")
	{
		users.DELETE("/:id", UserApi.DeleteEndpoint)                       //删除用户
		users.GET("", UserApi.AllEndpoint)                                 //获取所有用户
		users.GET("/paging", UserApi.PagingEndpoint)                       //分页查询用户
		users.PUT("/:id", UserApi.UpdateEndpoint)                          //编辑用户
		users.PATCH("/:id/status", UserApi.UpdateStatusEndpoint)           //修改用户状态
		users.POST("/:id/change-password", UserApi.ChangePasswordEndpoint) //修改用户密码
		users.GET("/:id", UserApi.DetailsEndpoint)                         //用户详情
	}

	frid := r.Group("/frid")
	{
		frid.GET("", FridApi.AllFridEndpoint)          //好友列表
		frid.POST("/apply", FridApi.ApplyEndpoint)     //好友申请
		frid.GET("/list", FridApi.ApplyListEndpoint)   //好友申请列表
		frid.PUT("/:id/status", FridApi.HandeEndpoint) //处理好友请求
		frid.DELETE("/:id", FridApi.DeleteEndpoint)    //删除好友
	}

	trends := r.Group("/trends")
	{
		trends.GET("/paging", TrendsApi.AllTrendsEndpoint)    //个人(好友)动态列表
		trends.POST("", TrendsApi.CreateEndpoint)             //个人动态发布
		trends.POST("/comment", TrendsApi.CommentEndpoint)    //评论
		trends.DELETE("/:type/:id", TrendsApi.DeleteEndpoint) //删除
	}

	topics := r.Group("/topics", mw.Admin)
	{
		topics.POST("", TopicsApi.CreateEndpoint) //创建话题
	}

	r.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8282")
}
