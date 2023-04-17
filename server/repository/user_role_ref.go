package repository

var UserRoleRefRepository = new(userRoleRefRepository)

type userRoleRefRepository struct {
	baseRepository
}

// func (r userRoleRefRepository) Create(c context.Context, m *model.UserRoleRef) error {
// 	return r.GetDB(c).Create(m).Error
// }

// func (r userRoleRefRepository) DeleteByUserId(c context.Context, userId string) error {
// 	return r.GetDB(c).Where("user_id = ?", userId).Delete(model.UserRoleRef{}).Error
// }
