package api

import (
	"context"
	"next-social/server/common/nt"
	"next-social/server/dto"
	"next-social/server/global/cache"
	"next-social/server/repository"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginAccount dto.LoginAccount

	if err := c.Bind(&loginAccount); err != nil {
		//return err
		ShowError(c, err)
	}

	// 存储登录失败次数信息
	loginFailCountKey := c.ClientIP() + loginAccount.Username
	v, ok := cache.LoginFailedKeyManager.Get(loginFailCountKey)
	if !ok {
		v = 1
	}
	count := v.(int)
	if count >= 5 {
		Fail(c, -1, "登录失败次数过多，请等待5分钟后再试")
		return
	}

	user, err := repository.UserRepository.FindByUsername(context.TODO(), loginAccount.Username)
	if err != nil {
		count++
		cache.LoginFailedKeyManager.Set(loginFailCountKey, count, cache.LoginLockExpiration)
		// 保存登录日志
		// if err := service.UserService.SaveLoginLog(c.RealIP(), c.Request().UserAgent(), loginAccount.Username, false, loginAccount.Remember, "", "账号或密码不正确"); err != nil {
		// 	return err
		// }
		FailWithData(c, -1, "您输入的账号或密码不正确", count)
		return
	}

	if user.Status == nt.StatusDisabled {
		Fail(c, -1, "该账户已停用")
		return
	}
}
