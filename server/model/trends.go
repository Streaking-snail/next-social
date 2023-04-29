package model

import "next-social/server/common"

type Trends struct {
	Id      int             `gorm:"primary_key;auto_increment;not null;type:int(11)" json:"id"`
	UserID  string          `gorm:"index;not null;type:varchar(36)" json:"user_id"`
	Created common.JsonTime `json:"created"`
	Content string          `gorm:"type:varchar(1024)" json:"content"`
	//Status  int             `gorm:"type:tinyint(3)" json:"has_delete"`
}

func (t Trends) TableName() string {
	return "trends"
}

type TrendsComment struct {
	Id       int             `gorm:"primary_key;auto_increment;not null;type:int(11)" json:"id"`
	TrendsID string          `gorm:"index;not null;type:int(11)" json:"trends_id"`
	Created  common.JsonTime `json:"created"`
	Content  string          `gorm:"type:varchar(1024)" json:"content"`
	ParendID uint8           `gorm:"type:tinyint(3)" json:"parend_id"`
	//Status   int             `gorm:"type:tinyint(3)" json:"has_delete"`
}

func (t TrendsComment) TableName() string {
	return "trends_comment"
}

type TrendsForPage struct {
	Id      int             `json:"id"`
	UserID  string          `json:"user_id"`
	Created common.JsonTime `json:"created"`
	Content string          `json:"content"`
	Comment []TrendsComment
}
