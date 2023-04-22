package model

import (
	"next-social/server/common"
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
	UserID   string          `gorm:"primary_key;type:varchar(36)" json:"id"`
	FriendID string          `gorm:"primary_key;type:varchar(36)" json:"friend_id"`
	Status   int8            `json:"status"`
	Created  common.JsonTime `json:"created"`
	Handle   common.JsonTime `json:"handle"`
}

func (r *UserApply) TableName() string {
	return "user_apply"
}
