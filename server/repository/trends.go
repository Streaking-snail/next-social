package repository

import (
	"context"
	"next-social/server/model"
)

var TrendsRepository = new(trendsRepository)

type trendsRepository struct {
	baseRepository
}

func (r trendsRepository) FindTrends(c context.Context, userId string, pageIndex, pageSize int) (o []model.Trends, err error) {
	err = r.GetDB(c).Select("id, user_id, created, content").
		Where("user_id = ?", userId).Order("id desc").
		Find(&o).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Error
	return
}

func (r trendsRepository) FindComment(c context.Context, trend_ids []int) (o []model.TrendsComment, err error) {
	err = r.GetDB(c).Where("trends_id in (?)", trend_ids).Order("id desc").Find(&o).Error
	return
}

func (r trendsRepository) Create(c context.Context, o *model.Trends) (err error) {
	return r.GetDB(c).Create(&o).Error
}

func (r trendsRepository) GetFrid(c context.Context, trends_id int) (trends model.Trends, err error) {
	err = r.GetDB(c).Select("user_id").Where("id = ?", trends_id).First(&trends).Error
	return
}

func (r trendsRepository) CreateComment(c context.Context, o *model.TrendsComment) (err error) {
	err = r.GetDB(c).Create(&o).Error
	return
}
