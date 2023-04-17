package api

import (
	"context"
	"next-social/server/common/maps"
	"next-social/server/model"
	"next-social/server/repository"
	"next-social/server/service"
	"strconv"
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

func (u UserApi) PagingEndpoint(c *gin.Context) {
	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	username := c.Query("username")
	nickname := c.Query("nickname")
	mail := c.Query("mail")

	order := c.Query("order")
	field := c.Query("field")
	online := c.Query("online")

	items, total, err := repository.UserRepository.Find(context.TODO(), pageIndex, pageSize, username, nickname, mail, online, "", order, field)
	if err != nil {
		ShowError(c, err)
		return
	}

	Success(c, maps.Map{
		"total": total,
		"items": items,
	})
	return
}

func (u UserApi) AllEndpoint(c *gin.Context) {
	users, err := repository.UserRepository.FindAll(context.Background())
	if err != nil {
		ShowError(c, err)
		return
	}
	items := make([]maps.Map, len(users))
	for i, user := range users {
		items[i] = maps.Map{
			"id":       user.ID,
			"nickname": user.Nickname,
		}
	}
	Success(c, items)
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
