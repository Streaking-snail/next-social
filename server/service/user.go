package service

import (
	"context"
	"fmt"
	"next-social/server/common"
	"next-social/server/common/nt"
	"next-social/server/env"
	"next-social/server/model"
	"next-social/server/repository"
	"next-social/server/utils"

	"gorm.io/gorm"
)

const SuperAdminID = `abcdefghijklmnopqrstuvwxyz`

var UserService = new(userService)

type userService struct {
	baseService
}

func (service userService) IsSuperAdmin(userId string) bool {
	return SuperAdminID == userId
}

func (service userService) CreateUser(user model.User) (err error) {
	return env.GetDB().Transaction(func(tx *gorm.DB) error {
		c := service.Context(tx)
		exist, err := repository.UserRepository.ExistByUsername(c, user.Username)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("username %s is already used", user.Username)
		}
		password := user.Password
		var pass []byte
		if pass, err = utils.Encoder.Encode([]byte(password)); err != nil {
			return err
		}
		user.Password = string(pass)

		user.ID = utils.UUID()
		user.Created = common.NowJsonTime()
		user.Status = nt.StatusEnabled

		if err := repository.UserRepository.Create(c, &user); err != nil {
			return err
		}
		if err := service.saveUserRoles(c, user); err != nil {
			return err
		}
		// if err := StorageService.CreateStorageByUser(c, &user); err != nil {
		// 	return err
		// }

		// 		if user.Mail != "" {
		// 			subject := fmt.Sprintf("%s 注册通知", branding.Name)
		// 			text := fmt.Sprintf(`您好，%s。
		// 	管理员为你开通了账户。
		// 	账号：%s
		// 	密码：%s
		// `, user.Username, user.Username, password)
		// 			go MailService.SendMail(user.Mail, subject, text)
		// 		}
		return nil
	})
}

func (service userService) saveUserRoles(c context.Context, user model.User) error {
	for _, role := range user.Roles {
		ref := &model.UserRoleRef{
			ID:     utils.UUID(),
			UserId: user.ID,
			RoleId: role,
		}
		if err := repository.UserRoleRefRepository.Create(c, ref); err != nil {
			return err
		}
	}
	return nil
}
