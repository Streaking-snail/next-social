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

func (r topicsRepository) DeleteById(c context.Context, id int) (err error) {
	return r.GetDB(c).Where("id = ?", id).Delete(&model.Topics{}).Error
}

func (r topicsRepository) Update(c context.Context, o *model.Topics) (err error) {
	return r.GetDB(c).Updates(&o).Error
}
