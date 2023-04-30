package service

import (
	"context"
	"fmt"
	"next-social/server/model"
	"next-social/server/repository"
)

var TrendsService = new(trendsService)

type trendsService struct {
	baseService
}

func (trends trendsService) GetTrends(userId string, pageIndex, pageSize int) (items []model.TrendsForPage, err error) {
	trend, err := repository.TrendsRepository.FindTrends(context.TODO(), userId, pageIndex, pageSize)
	if err != nil {
		return nil, err
	}
	//return trend, nil
	if len(trend) == 0 {
		return nil, nil
	}
	var trend_ids []int
	for _, v := range trend {
		trend_ids = append(trend_ids, v.Id)
	}

	trendsComment, err := repository.TrendsRepository.FindComment(context.TODO(), trend_ids)
	if err != nil {
		return nil, err
	}

	for i, j := range trend {
		arr := model.TrendsForPage{Id: j.Id, UserID: j.UserID, Created: j.Created, Content: j.Content}
		items = append(items, arr)
		for _, v := range trendsComment {
			if v.TrendsID == j.Id {
				items[i].Comment = append(items[i].Comment, v)
			}
		}
	}
	return
}

func (trends trendsService) CreateComment(trendsComment *model.TrendsComment, user_id string) error {
	//查询发布者用户id
	trend, err := repository.TrendsRepository.GetFrid(context.TODO(), trendsComment.TrendsID)
	if err != nil {
		return err
	}

	//查询双方是否存在好友关系
	exist, err := repository.FridRepository.ExistByFrid(context.TODO(), trend.UserID, user_id)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("好友关系不存在")
	}

	if err := repository.TrendsRepository.CreateComment(context.TODO(), trendsComment); err != nil {
		return err
	}
	return nil
}
