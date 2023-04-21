package api

import (
	"context"
	"next-social/server/repository"

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
