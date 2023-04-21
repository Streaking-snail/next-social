package api

import (
	"context"
	"next-social/server/common/maps"
	"next-social/server/common/nt"
	"next-social/server/dto"
	"next-social/server/global/cache"
	"next-social/server/model"
	"next-social/server/repository"
	"next-social/server/service"
	"next-social/server/utils"

	"github.com/gin-gonic/gin"
)

type AccountInfo struct {
	Id         string   `json:"id"`
	Username   string   `json:"username"`
	Nickname   string   `json:"nickname"`
	Type       string   `json:"type"`
	EnableTotp bool     `json:"enableTotp"`
	Menus      []string `json:"menus"`
}

func Login(c *gin.Context) {
	var loginAccount dto.LoginAccount

	if err := c.Bind(&loginAccount); err != nil {
		ShowError(c, err)
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
		Fail(c, -1, "登录失败次数过多，请等待5分钟后再试")
		return
	}

	user, err := repository.UserRepository.FindByUsername(context.TODO(), loginAccount.Username)
	if err != nil {
		count++
		cache.LoginFailedKeyManager.Set(loginFailCountKey, count, cache.LoginLockExpiration)
		// 保存登录日志
		if err := service.UserService.SaveLoginLog(c.ClientIP(), c.Request.UserAgent(), loginAccount.Username, false, loginAccount.Remember, "", "账号或密码不正确"); err != nil {
			ShowError(c, err)
			return
		}
		FailWithData(c, -1, "您输入的账号或密码不正确", count)
		return
	}

	if user.Status == nt.StatusDisabled {
		Fail(c, -1, "该账户已停用")
		return
	}

	if err := utils.Encoder.Match([]byte(user.Password), []byte(loginAccount.Password)); err != nil {
		count++
		cache.LoginFailedKeyManager.Set(loginFailCountKey, count, cache.LoginLockExpiration)
		// 保存登录日志
		if err := service.UserService.SaveLoginLog(c.ClientIP(), c.Request.UserAgent(), loginAccount.Username, false, loginAccount.Remember, "", "账号或密码不正确"); err != nil {
			ShowError(c, err)
			return
		}
		FailWithData(c, -1, "您输入的账号或密码不正确", count)
		return
	}

	token, err := LoginSuccess(loginAccount, user, c.ClientIP())
	if err != nil {
		ShowError(c, err)
		return
	}
	// 保存登录日志
	if err := service.UserService.SaveLoginLog(c.ClientIP(), c.Request.UserAgent(), loginAccount.Username, true, loginAccount.Remember, token, ""); err != nil {
		ShowError(c, err)
	}

	// var menus []string
	// if service.UserService.IsSuperAdmin(user.ID) {
	// 	menus = service.MenuService.GetMenus()
	// } else {
	// 	roles, err := service.RoleService.GetRolesByUserId(user.ID)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	for _, role := range roles {
	// 		items := service.RoleService.GetMenuListByRole(role)
	// 		menus = append(menus, items...)
	// 	}
	// }

	info := AccountInfo{
		Id:         user.ID,
		Username:   user.Username,
		Nickname:   user.Nickname,
		Type:       user.Type,
		EnableTotp: user.TOTPSecret != "" && user.TOTPSecret != "-",
		//Menus:      menus,
	}

	Success(c, maps.Map{
		"info":  info,
		"token": token,
	})
}

func LoginSuccess(loginAccount dto.LoginAccount, user model.User, ip string) (string, error) {

	token := utils.LongUUID()

	authorization := dto.Authorization{
		Token:    token,
		Type:     nt.LoginToken,
		Remember: loginAccount.Remember,
		User:     &user,
	}

	if authorization.Remember {
		// 记住登录有效期两周
		cache.TokenManager.Set(token, authorization, cache.RememberMeExpiration)
	} else {
		cache.TokenManager.Set(token, authorization, cache.NotRememberExpiration)
	}

	b := true
	// 修改登录状态
	err := repository.UserRepository.Update(context.TODO(), &model.User{Online: &b, ID: user.ID})
	return token, err
}
