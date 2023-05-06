package model

import "next-social/server/common"

type Trends struct {
	Id      int             `gorm:"primary_key;auto_increment;7type:int(11)" json:"id"`
	UserID  string          `gorm:"index;not null;type:varchar(36)" json:"user_id"` //发布者id
	Created common.JsonTime `json:"created"`                                        //发布时间
	Content string          `gorm:"type:varchar(1024)" json:"content"`              //动态内容
}

func (t Trends) TableName() string {
	return "trends"
}

type TrendsComment struct {
	Id       int             `gorm:"primary_key;auto_increment;not null;type:int(11)" json:"id"`
	TrendsID int             `gorm:"index;not null;type:int(11)" json:"trends_id"` //动态id
	UserID   string          `gorm:"type:varchar(36)" json:"user_id"`              //评论者id
	Created  common.JsonTime `json:"created"`                                      //发布时间
	Content  string          `gorm:"type:varchar(1024)" json:"content"`            //评论内容
	ParendID int             `gorm:"type:int(11)" json:"parend_id"`                //回复评论id
}

func (t TrendsComment) TableName() string {
	return "trends_comment"
}

type TrendsLikes struct {
	Id       int             `gorm:"primary_key;auto_increment;not null;type:int(11)" json:"id"`
	TrendsID int             `gorm:"index;not null;type:int(11)" json:"trends_id"` //动态id
	UserID   string          `gorm:"type:varchar(36)" json:"user_id"`              //点赞者id
	Created  common.JsonTime `json:"created"`
}

func (t TrendsLikes) TableName() string {
	return "trends_likes"
}

type TrendsForPage struct {
	Id      int             `json:"id"`
	UserID  string          `json:"user_id"`
	Created common.JsonTime `json:"created"`
	Content string          `json:"content"`
	Comment []TrendsComment
	Likes   []string
}
