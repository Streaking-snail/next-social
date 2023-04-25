package model

import (
	"next-social/server/common"
)

type ChitChat struct {
	Id        int             `gorm:"primary_key;auto_increment;not null;type:int(11)" json:"id"`
	UserID    string          `gorm:"index;not null;type:varchar(36)" json:"user_id"`
	FriendID  string          `gorm:"index;not null;type:varchar(36)" json:"friend_id"`
	HasRead   int             `gorm:"type:tinyint(3)" json:"has_read"`
	Created   common.JsonTime `json:"created"`
	HasDelete int             `gorm:"type:tinyint(3)" json:"has_delete"`
	Message   string          `gorm:"type:varchar(1024)" json:"message"`
}

func (c ChitChat) TableName() string {
	return "chit_chat"
}
