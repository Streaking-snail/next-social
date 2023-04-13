package api

import (
	"next-social/server/model"
	"next-social/server/service"

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
