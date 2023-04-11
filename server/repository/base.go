package repository

import (
	"context"

	"next-social/server/common/nt"
	"next-social/server/env"

	"gorm.io/gorm"
)

type baseRepository struct {
}

func (b *baseRepository) GetDB(c context.Context) *gorm.DB {
	db, ok := c.Value(nt.DB).(*gorm.DB)
	if !ok {
		return env.GetDB()
	}
	return db
}
