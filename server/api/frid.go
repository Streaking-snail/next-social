package api

import (
	"context"
	"next-social/server/repository"
	"next-social/server/service"

	"github.com/gin-gonic/gin"
)

type FridApi struct{}

func (frid FridApi) AllFridEndpoint(c *gin.Context) {
	//id := c.Query("id")
	account, found := GetCurrentAccount(c)
	if !found {
		Fail(c, -1, "获取当前登录账户失败")
		return
	}
	users, err := repository.FridRepository.FindAll(context.TODO(), account.ID)
	if err != nil {
		ShowError(c, err)
		return
	}

	Success(c, users)
}

func (frid FridApi) ApplyEndpoint(c *gin.Context) {
	username := c.PostForm("username")
	friendId, err := repository.UserRepository.FindByUsername(context.TODO(), username)
	if err != nil {
		ShowError(c, err)
	}
	account, found := GetCurrentAccount(c)
	if !found {
		Fail(c, -1, "获取当前登录账户失败")
		return
	}

	if err := service.FridService.ApplyForUserId(account.ID, friendId.ID); err != nil {
		ShowError(c, err)
		return
	}
	Success(c, "好友申请已发送")
}
