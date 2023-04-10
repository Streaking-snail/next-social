package api

import (
	"next-social/server/dto"
	"next-social/server/global/cache"
	"next-social/server/utils/response"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginAccount dto.LoginAccount

	if err := c.Bind(&loginAccount); err != nil {
		//return err
		response.ShowValidatorError(c, err)
		return
	}

	// 存储登录失败次数信息
	loginFailCountKey := c.ClientIP() + loginAccount.Username
	v, ok := cache.LoginFailedKeyManager.Get(loginFailCountKey)
	if !ok {
		v = 1
	}
	count := v.(int)
	if count >= 5 {
		response.ShowError(c, "登录失败次数过多，请等待5分钟后再试")
		//return Fail(c, -1, )
	}

	response.ShowSuccess(c, "登录成功")

}
