package repository

import (
	"context"
	"next-social/server/model"
)

var TopicsRepository = new(topicsRepository)

type topicsRepository struct {
	baseRepository
}

func (r topicsRepository) Create(c context.Context, o *model.Topics) (err error) {
	return r.GetDB(c).Create(&o).Error
}
