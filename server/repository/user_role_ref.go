package repository

import (
	"context"

	"next-social/server/model"
)

var UserRoleRefRepository = new(userRoleRefRepository)

type userRoleRefRepository struct {
	baseRepository
}

func (r userRoleRefRepository) Create(c context.Context, m *model.UserRoleRef) error {
	return r.GetDB(c).Create(m).Error
}
