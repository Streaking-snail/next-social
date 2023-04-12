package repository

import (
	"context"

	"next-social/server/model"
)

var (
	LoginPolicyUserRefRepository = new(loginPolicyUserRefRepository)
	TimePeriodRepository         = new(timePeriodRepository)
)

type loginPolicyUserRefRepository struct {
	baseRepository
}

func (r loginPolicyUserRefRepository) FindByUserId(c context.Context, userId string) (items []model.LoginPolicyUserRef, err error) {
	err = r.GetDB(c).Where("user_id = ?", userId).Find(&items).Error
	return
}

type timePeriodRepository struct {
	baseRepository
}

func (r timePeriodRepository) FindByLoginPolicyId(c context.Context, loginPolicyId string) (items []model.TimePeriod, err error) {
	err = r.GetDB(c).Where("login_policy_id = ?", loginPolicyId).Find(&items).Error
	return
}
