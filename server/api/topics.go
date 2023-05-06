package api

import (
	"context"
	"next-social/server/common"
	"next-social/server/model"
	"next-social/server/repository"
	"strconv"

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
		return
	}
	Success(c, "创建成功")
}

func (t TopicsApi) DeleteEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	//topics_id := strconv.Atoi(id)
	if err := repository.TopicsRepository.DeleteById(context.TODO(), id); err != nil {
		ShowError(c, err)
		return
	}
	Success(c, "删除成功")
}

func (t TopicsApi) UpdateStatusEndpoint(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	status := c.Param("status")
	if err := Service.TopicsService.UpdateStatusById(id, status); err != nil {
		ShowError(c, err)
		return
	}
	Success(c, "修改成功")
}
