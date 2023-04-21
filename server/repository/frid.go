package repository

import (
	"context"
	"next-social/server/model"
)

var FridRepository = new(fridRepository)

type fridRepository struct {
	baseRepository
}

func (r fridRepository) FindAll(c context.Context, id string) (o []model.User, err error) {
	// err = r.GetDB(c).Raw("SELECT friend_id FROM user_relation WHERE user_id = ? UNION SELECT user_id FROM user_relation WHERE friend_id = ?", id, id).Scan(&friends).Error
	// return
	var user_one []model.User
	err = r.GetDB(c).Table("user_relation rel").Select("users.username,users.nickname,users.mail").Where("rel.user_id = ?", id).
		Joins("left JOIN users on users.ID = rel.friend_id").Scan(&user_one).Error
	if err != nil {
		return
	}
	o = append(o, user_one...)
	var user_two []model.User
	err = r.GetDB(c).Table("user_relation rel").Select("users.*").Where("rel.friend_id = ?", id).
		Joins("left JOIN users on users.ID = rel.user_id").Scan(&user_two).Error
	o = append(o, user_two...)
	return
}
