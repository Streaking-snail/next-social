package api

import (
	"context"
	"next-social/server/common"
	"next-social/server/model"
	"next-social/server/repository"
	"strconv"

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

	// online := c.Query("online")

	items, err := repository.TrendsRepository.Find(context.TODO(), account.ID, pageIndex, pageSize)
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

// func (t TrendsApi) CommentEndpoint(c *gin.Context) {
// 	// account, found := GetCurrentAccount(c)
// 	// if !found {
// 	// 	Fail(c, -1, "获取当前登录账户失败")
// 	// 	return
// 	// }
// 	trends_id := c.PostForm("id")
// 	Content := c.PostForm("content")
// 	if Content == "" {
// 		Fail(c, -1, "内容不能为空")
// 		return
// 	}
// 	var trendsComment = model.TrendsComment{
// 		TrendsID: trends_id,
// 		//UserID:   account.ID,
// 		Created: common.NowJsonTime(),
// 		Content: Content,
// 		//	ParendID:
// 	}
// 	if err := service.TrendsService.CreateComment(&trendsComment); err != nil {
// 		ShowError(c, err)
// 		return
// 	}
// 	Success(c, "评论成功")
// }
