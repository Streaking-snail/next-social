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

// 好友申请
func (r fridRepository) Create(c context.Context, o *model.UserApply) error {
	return r.GetDB(c).Create(o).Error
}

func (r fridRepository) FindByStatus(c context.Context, o *model.UserApply) (exist bool, err error) {
	//userapply := model.UserApply{}
	var count int8
	err = r.GetDB(c).Table(o.TableName()).Select("count(*)").Where(&o).Find(&count).Error
	if err != nil {
		return false, err
	}
	return count <= 0, nil
}

func (r fridRepository) Update(c context.Context, o *model.UserApply) (err error) {
	return r.GetDB(c).Updates(o).Error
}

// 增加好友关系
func (r fridRepository) HandApple(c context.Context, o *model.UserRelation) (err error) {
	return r.GetDB(c).Create(o).Error
}
