package api

import (
	"context"
	"next-social/server/common"
	"next-social/server/model"
	"next-social/server/repository"

	"github.com/gin-gonic/gin"
)

type TopicsApi struct{}

func (t TopicsApi) CreateEndpoint(c *gin.Context) {

	name := c.PostForm("name")
	description := c.PostForm("description")
	topics := model.Topics{
		Name:        name,
		Description: description,
		Created:     common.NowJsonTime(),
	}
	if err := repository.TopicsRepository.Create(context.TODO(), &topics); err != nil {
		ShowError(c, err)
	}
	Success(c, "创建成功")
}
