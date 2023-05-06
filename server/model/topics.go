package model

import "next-social/server/common"

type Topics struct {
	Id          int             `gorm:"primary_key;auto_increment;not null;type:int(11)" json:"id"`
	Name        string          `gorm:"type:varchar(255);not null" json:"name"`          //话题名称
	Description string          `gorm:"type:varchar(1024)" json:"description"`           //话题的描述
	Status      string          `gorm:"type:char(1);not null;default:'0'" json:"status"` //状态 0正常1屏蔽2删除
	Created     common.JsonTime `json:"created"`
}

func (t Topics) TableName() string {
	return "topics"
}
