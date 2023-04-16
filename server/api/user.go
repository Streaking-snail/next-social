package api

import (
	"next-social/server/model"
	"next-social/server/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (u UserApi) CreateEndpoint(c *gin.Context) {
	var item model.User
	if err := c.Bind(&item); err != nil {
		ShowError(c, err)
		return
	}

	if err := service.UserService.CreateUser(item); err != nil {
		ShowError(c, err)
		return
	}

	Success(c, item)
	return
}

func (userApi UserApi) DeleteEndpoint(c *gin.Context) {
	ids := c.Param("id")
	account, found := GetCurrentAccount(c)
	if !found {
		Fail(c, -1, "获取当前登录账户失败")
		return
	}
	split := strings.Split(ids, ",")
	for i := range split {
		userId := split[i]
		if account.ID == userId {
			Fail(c, -1, "不允许删除自身账户")
			return
		}
		if err := service.UserService.DeleteUserById(userId); err != nil {
			ShowError(c, err)
			return
		}
	}

	Success(c, nil)
	return
}
