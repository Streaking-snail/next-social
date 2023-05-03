package api

import (
	"context"
	"next-social/server/common"
	"next-social/server/model"
	"next-social/server/repository"
	"next-social/server/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type TrendsApi struct{}

func (t TrendsApi) AllTrendsEndpoint(c *gin.Context) {
	account, found := GetCurrentAccount(c)
	if !found {
		Fail(c, -1, "获取当前登录账户失败")
		return
	}
	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	user_ids := []string{account.ID}
	fird := c.Query("fird")
	if fird != "" {
		find_ids, err := repository.FridRepository.FindAll(context.TODO(), account.ID)
		if err != nil {
			ShowError(c, err)
			return
		}
		for _, v := range find_ids {
			user_ids = append(user_ids, v.ID)
		}
	}

	items, err := service.TrendsService.GetTrends(user_ids, pageIndex, pageSize)
	if err != nil {
		ShowError(c, err)
		return
	}
	Success(c, items)
}

func (t TrendsApi) CreateEndpoint(c *gin.Context) {
	account, found := GetCurrentAccount(c)
	if !found {
		Fail(c, -1, "获取当前登录账户失败")
		return
	}
	var content = c.PostForm("content")
	if content == "" {
		Fail(c, -1, "内容不能为空")
		return
	}
	var Trends = model.Trends{
		UserID:  account.ID,
		Content: content,
		Created: common.NowJsonTime(),
	}

	if err := repository.TrendsRepository.Create(context.TODO(), &Trends); err != nil {
		ShowError(c, err)
		return
	}
	Success(c, "发布成功")
}

func (t TrendsApi) CommentEndpoint(c *gin.Context) {
	account, found := GetCurrentAccount(c)
	if !found {
		Fail(c, -1, "获取当前登录账户失败")
		return
	}
	trends := c.PostForm("id")
	trends_id, _ := strconv.Atoi(trends)
	Content := c.PostForm("content")
	Content = strings.TrimSpace(Content)
	Parend := c.PostForm("parendID")
	ParendID, _ := strconv.Atoi(Parend)
	if Content == "" {
		Fail(c, -1, "评论内容不能为空")
		return
	}
	var trendsComment = model.TrendsComment{
		TrendsID: trends_id,
		UserID:   account.ID,
		Created:  common.NowJsonTime(),
		Content:  Content,
		ParendID: ParendID,
	}
	if err := service.TrendsService.CreateComment(&trendsComment, account.ID); err != nil {
		ShowError(c, err)
		return
	}
	Success(c, "评论成功")
}

func (t TrendsApi) DeleteEndpoint(c *gin.Context) {
	account, found := GetCurrentAccount(c)
	if !found {
		Fail(c, -1, "获取当前登录账户失败")
		return
	}
	del_type := c.Param("type")
	id := c.Param("id")
	del_id, _ := strconv.Atoi(id)
	if err := repository.TrendsRepository.DeleteById(context.TODO(), del_type, account.ID, del_id); err != nil {
		ShowError(c, err)
		return
	}
	Success(c, "删除成功")

}
