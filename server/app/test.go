package app

import (
	"context"
)

type baseRepository struct {
}

func (b *baseRepository) GetDB(c context.Context) string {

	return "aaaaaaa"
}
