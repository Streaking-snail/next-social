package repository

import (
	"context"
	"next-social/server/model"
	"time"
)

var FridRepository = new(fridRepository)

type fridRepository struct {
	baseRepository
}

func (r fridRepository) FindAll(c context.Context, id string) (o []model.User, err error) {
	var user_one []model.User
	err = r.GetDB(c).Table("user_relation rel").Select("user.id,users.username,users.nickname,users.mail").Where("rel.user_id = ?", id).
		Joins("left JOIN users on users.ID = rel.friend_id").Scan(&user_one).Error
	if err != nil {
		return
	}
	o = append(o, user_one...)
	var user_two []model.User
	err = r.GetDB(c).Table("user_relation rel").Select("user.id,users.username,users.nickname,users.mail").Where("rel.friend_id = ?", id).
		Joins("left JOIN users on users.ID = rel.user_id").Scan(&user_two).Error
	o = append(o, user_two...)
	return
}

func (r fridRepository) FindAllApply(c context.Context, userId string) (o []model.UserApply, err error) {
	err = r.GetDB(c).Where("Created > ?", time.Now().AddDate(0, -1, 0)).Find(&o).Error
	return
}

// 好友申请
func (r fridRepository) Create(c context.Context, o *model.UserApply) error {
	return r.GetDB(c).Create(o).Error
}

func (r fridRepository) FindByStatus(c context.Context, o *model.UserApply) (userapply model.UserApply, err error) {
	//userapply = model.UserApply{}
	//var count int8
	err = r.GetDB(c).Where(&o).First(&userapply).Error
	return
	// if err != nil {
	// 	return false, err
	// }
	// return count <= 0, nil
}

func (r fridRepository) Update(c context.Context, o *model.UserApply) (err error) {
	return r.GetDB(c).Updates(o).Error
}

// 增加好友关系
func (r fridRepository) HandApple(c context.Context, o *model.UserRelation) (err error) {
	return r.GetDB(c).Create(o).Error
}

func (r fridRepository) DeleteFrid(c context.Context, userId string, friend_id string) (err error) {
	err = r.GetDB(c).Where("(user_id = ? and FriendID = ?) or (user_id = ? and FriendID = ?)", userId, friend_id, friend_id, userId).Delete(&model.UserRelation{}).Error
	if err == nil {
		err = r.GetDB(c).Table("user_apply").Where("(user_id = ? and friend_id = ?) or (user_id = ? and friend_id = ?)", userId, friend_id, friend_id, userId).Update("Status", 4).Error
	}
	return
}

func (r fridRepository) ExistByFrid(c context.Context, userId, friend_id string) (exist bool, err error) {
	userRelation := model.UserRelation{}
	var count uint64
	err = r.GetDB(c).Table(userRelation.TableName()).Select("count(*)").
		Where("(user_id = ? and friend_id = ?) or (user_id = ? and friend_id = ?)", userId, friend_id, friend_id, userId).
		Find(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 超过三天未处理请求设置为过期
func (r fridRepository) AutoExpireEndpoint(c context.Context) {
	r.GetDB(c).Table("user_apply").Where("Created > ? and Status = ?", time.Now().AddDate(0, 0, -3), 0).Update("Status", 3)
}
