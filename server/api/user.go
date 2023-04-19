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
	//用户类型默认为普通用户
	if item.Type == "" {
		item.Type = "user"
	}
	//昵称默认为用户名
	if item.Nickname == "" {
		item.Nickname = item.Username
	}
	if err := service.UserService.CreateUser(item); err != nil {
		ShowError(c, err)
		return
	}

	Success(c, item)
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
}

func (u UserApi) DeleteEndpoint(c *gin.Context) {
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
}

func (u UserApi) UpdateEndpoint(c *gin.Context) {
	id := c.Param("id")

	var item model.User
	if err := c.Bind(&item); err != nil {
		ShowError(c, err)
		return
	}

	if err := service.UserService.UpdateUser(id, item); err != nil {
		ShowError(c, err)
		return
	}
	Success(c, nil)
}

func (u UserApi) UpdateStatusEndpoint(c *gin.Context) {
	id := c.Param("id")
	status := c.Query("status")
	account, _ := GetCurrentAccount(c)

	if account.Type != "admin" || id != account.ID || service.UserService.IsSuperAdmin(id) {
		Fail(c, -1, "不能修改超级管理员状态或权限不足")
		return
	}

	if err := service.UserService.UpdateStatusById(id, status); err != nil {
		ShowError(c, err)
		return
	}
	Success(c, nil)
}

func (u UserApi) ChangePasswordEndpoint(c *gin.Context) {
	id := c.Param("id")
	password := c.PostForm("password")
	account, _ := GetCurrentAccount(c)

	if password == "" {
		Fail(c, -1, "密码不能为空")
		return
	}
	if account.Type != "admin" || id != account.ID || (service.UserService.IsSuperAdmin(id) && service.UserService.IsSuperAdmin(account.ID)) {
		Fail(c, -1, "不能修改超级管理员密码或权限不足")
		return
	}

	if err := service.UserService.ChangePassword(id, password); err != nil {
		ShowError(c, err)
		return
	}
}

func (u UserApi) DetailsEndpoint(c *gin.Context) {
	id := c.Param("id")
	account, _ := GetCurrentAccount(c)
	if account.Type != "admin" || id != account.ID || (service.UserService.IsSuperAdmin(id) && service.UserService.IsSuperAdmin(account.ID)) {
		Fail(c, -1, "权限不足")
		return
	}
	user, err := repository.UserRepository.FindById(context.Background(), id)
	if err != nil {
		ShowError(c, err)
		return
	}
	Success(c, user)
}
