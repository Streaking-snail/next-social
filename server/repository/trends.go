package repository

import (
	"context"
	"next-social/server/model"
)

var TrendsRepository = new(trendsRepository)

type trendsRepository struct {
	baseRepository
}

func (r trendsRepository) FindAll(c context.Context, userId string) (o []model.Trends, err error) {
	err = r.GetDB(c).Find(&o).Error
	return
}
