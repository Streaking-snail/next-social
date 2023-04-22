package service

import (
	"context"
	"next-social/server/common"
	"next-social/server/model"
)

var FridService = new(fridService)

type fridService struct {
	baseService
}

func (frid fridService) ApplyForUserId(userId, friendId string) error {
	var userapply = model.UserApply{
		UserID:   userId,
		FriendID: friendId,
		Status:   0,
		Created:  common.NowJsonTime(),
	}
	if err := repository.FridRepository.Create(context.TODO(), &userapply); err != nil {
		return err
	}
	return nil
}
