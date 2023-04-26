package api

import (
	"context"

	"github.com/gin-gonic/gin"
)

type TrendsApi = struct{}

func (t TrendsApi) AllTrendsEndpoint(c *gin.Context) {
	account, found := GetCurrentAccount(c)
	if !found {
		Fail(c, -1, "获取当前登录账户失败")
		return
	}
	items, err := repository.TrendsRepository.FindAll(context.TODO(), account.ID)
	if err != nil {
		ShowError(c, err)
		return
	}
	Success(c, items)
}
