package model

// type UserRelation struct {
// 	User_id    string `gorm:"primary_key,type:varchar(36)" json:"id"`
// 	Friend_id  string `gorm:"primary_key,type:varchar(36)" json:"friend_id"`
// 	Sorted_key string `gorm:"type:varchar(128)" json:"sorted_key"`
// }
type UserRelation struct {
	UserID    string `gorm:"primary_key;type:varchar(36)" json:"id"`
	FriendID  string `gorm:"primary_key;type:varchar(36)" json:"friend_id"`
	SortedKey string `gorm:"type:varchar(128)" json:"sorted_key"`
}

func (r *UserRelation) TableName() string {
	return "user_relation"
}
