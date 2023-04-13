package model

type UserRoleRef struct {
	ID     string `gorm:"primary_key,type:varchar(36)" json:"id"`
	UserId string `gorm:"index,type:varchar(36)" json:"userId"`
	RoleId string `gorm:"index,type:varchar(36)" json:"roleId"`
}

func (r *UserRoleRef) TableName() string {
	return "users_roles_ref"
}
