package api

import (
	"context"
	"next-social/server/repository"
	"next-social/server/service"
	"strconv"

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
		return
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

func (frid FridApi) HandeEndpoint(c *gin.Context) {
	account, found := GetCurrentAccount(c)
	if !found {
		Fail(c, -1, "获取当前登录账户失败")
		return
	}
	friendId := c.Param("id")
	str_status := c.Query("status")
	status, err := strconv.Atoi(str_status)
	if err != nil {
		ShowError(c, err)
		return
	}
	if err := service.FridService.HandleApply(account.ID, friendId, status); err != nil {
		ShowError(c, err)
		return
	}
	Success(c, "处理完成")
}

func (frid FridApi) ApplyListEndpoint(c *gin.Context) {
	account, found := GetCurrentAccount(c)
	if !found {
		Fail(c, -1, "获取当前登录账户失败")
		return
	}
	list, err := repository.FridRepository.FindAllApply(context.TODO(), account.ID)
	if err != nil {
		ShowError(c, err)
		return
	}
	Success(c, list)
}
func (frid FridApi) DeleteEndpoint(c *gin.Context) {
	friend_id := c.Param("id")
	account, found := GetCurrentAccount(c)
	if !found {
		Fail(c, -1, "获取当前登录账户失败")
		return
	}
	if err := service.FridService.DeleteFridById(account.ID, friend_id); err != nil {
		ShowError(c, err)
		return
	}
	Success(c, "删除成功")
}

// 每天执行，将超过三天未处理的请求设置为过期
func (frid FridApi) AutoExpireEndpoint() {
	repository.FridRepository.AutoExpireEndpoint(context.TODO())
}
