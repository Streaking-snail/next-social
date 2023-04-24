package model

import (
	"next-social/server/common"
	"reflect"
)

type UserRelation struct {
	UserID    string `gorm:"primary_key;type:varchar(36)" json:"id"`
	FriendID  string `gorm:"primary_key;type:varchar(36)" json:"friend_id"`
	SortedKey string `gorm:"type:varchar(128)" json:"sorted_key"`
}

func (r *UserRelation) TableName() string {
	return "user_relation"
}

type UserApply struct {
	//ID       int             `gorm:"primary_key;auto_increment;not null;type:int(11)" json:"id"`
	UserID   string          `gorm:"primary_key;type:varchar(36)" json:"user_id"`
	FriendID string          `gorm:"primary_key;type:varchar(36)" json:"friend_id"`
	Status   int             `json:"status"`  //0:待处理 1：同意 2：拒绝 3：过期 4：已删除
	Created  common.JsonTime `json:"created"` //申请时间
	Handle   common.JsonTime `json:"handle"`  //处理时间
}

func (u UserApply) IsEmpty() bool {
	return reflect.DeepEqual(u, UserApply{})
}

func (r *UserApply) TableName() string {
	return "user_apply"
}
