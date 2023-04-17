package repository

import (
	"context"
	"next-social/server/model"
)

var LoginLogRepository = new(loginLogRepository)

type loginLogRepository struct {
	baseRepository
}

func (r loginLogRepository) FindAliveLoginLogsByUsername(c context.Context, username string) (o []model.LoginLog, err error) {
	err = r.GetDB(c).Where("state = '1' and logout_time is null and username = ?", username).Find(&o).Error
	return
}

func (r loginLogRepository) Create(c context.Context, o *model.LoginLog) (err error) {
	err = r.GetDB(c).Create(&o).Error
	return
}
