package model

import "next-social/server/common"

type Topics struct {
	Id          int             `gorm:"primary_key;auto_increment;not null;type:int(11)" json:"id"`
	Name        int             `gorm:"type:varchar(255)" json:"name"`         //话题名称
	Description string          `gorm:"type:varchar(1024)" json:"description"` //话题的描述
	Created     common.JsonTime `json:"created"`
}

func (t Topics) TableName() string {
	return "topics"
}
