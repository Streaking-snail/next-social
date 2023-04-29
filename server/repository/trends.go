package repository

import (
	"context"
	"next-social/server/model"
)

var TrendsRepository = new(trendsRepository)

type trendsRepository struct {
	baseRepository
}

func (r trendsRepository) Find(c context.Context, userId string, pageIndex, pageSize int) (items []model.TrendsForPage, err error) {
	m := model.Trends{}
	db := r.GetDB(c).Table(m.TableName()).Select("trends.id, trends.user_id, trends.created, trends.content, com.trends_id, com.created as comment_created, com.content as comment_content").
		Joins("left join TrendsComment as com on trends.id = com.trends_id")
	err = db.Where("trends.user_id = ?", userId).Order("trends.id desc").Find(&items).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Error
	return
	// err = r.GetDB(c).Find(&o).Error
	// return
}

func (r trendsRepository) Create(c context.Context, o *model.Trends) (err error) {
	return r.GetDB(c).Create(&o).Error
}
