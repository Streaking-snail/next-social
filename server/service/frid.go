package service

import (
	"context"
	"fmt"
	"next-social/server/common"
	"next-social/server/model"
	"next-social/server/repository"
)

var FridService = new(fridService)

type fridService struct {
	baseService
}

// 好友申请
func (frid fridService) ApplyForUserId(userId, friendId string) error {
	var userapply = model.UserApply{
		UserID:   userId,
		FriendID: friendId,
		// Status:   0,
		// Created:  common.NowJsonTime(),
	}
	exist, err := repository.FridRepository.FindByStatus(context.TODO(), &userapply)
	if err != nil {
		return err
	}

	userapply.Status = 0
	userapply.Created = common.NowJsonTime()

	if exist.IsEmpty() {
		if err := repository.FridRepository.Create(context.TODO(), &userapply); err != nil {
			return err
		}
	} else if exist.Status == 1 {
		return fmt.Errorf("已存在好友关系")
	} else {
		if err := repository.FridRepository.Update(context.TODO(), &userapply); err != nil {
			return err
		}
	}

	return nil
}

func (frid fridService) HandleApply(userId string, friendId string, status int) error {
	var userapply = model.UserApply{
		UserID:   userId,
		FriendID: friendId,
		Status:   0,
		//Handle:   common.NowJsonTime(),
	}
	exist, err := repository.FridRepository.FindByStatus(context.TODO(), &userapply)
	if err != nil {
		return err
	}
	if exist.IsEmpty() {
		return fmt.Errorf("请求已处理或不存在")
	} else {
		userapply.Status = status
		userapply.Handle = common.NowJsonTime()
	}
	//修改申请状态
	if err := repository.FridRepository.Update(context.TODO(), &userapply); err != nil {
		return err
	}
	//建立好友关系
	var userrelation = model.UserRelation{
		UserID:    userId,
		FriendID:  friendId,
		SortedKey: "A-B",
	}
	if err := repository.FridRepository.HandApple(context.TODO(), &userrelation); err != nil {
		return err
	}
	return nil
}

func (frid fridService) DeleteFridById(userId, friend_id string) error {
	if err := repository.FridRepository.DeleteFrid(context.TODO(), userId, friend_id); err != nil {
		return err
	}

	// if err := repository.FridRepository.DeleteFrid(context.TODO(), userId, friend_id); err != nil {
	// 	return err
	// }
	return nil
}
