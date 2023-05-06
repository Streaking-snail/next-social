package service

import (
	"context"
	"next-social/server/model"
	"next-social/server/repository"
)

var TopicsService = new(topicsService)

type topicsService struct {
	baseService
}

func (t topicsService) UpdateStatusById(id int) error {
	topics := model.Topics{
		Id:     id,
		Status: "1",
	}
	if err := repository.TopicsRepository.Update(context.TODO(), &topics); err != nil {
		return err
	}
	return nil
}
