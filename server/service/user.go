package service

const SuperAdminID = `abcdefghijklmnopqrstuvwxyz`

var UserService = new(userService)

type userService struct {
	baseService
}

func (service userService) IsSuperAdmin(userId string) bool {
	return SuperAdminID == userId
}
