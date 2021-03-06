package storage

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/utils"
)

func (dataBase SqlDB) CreateUser(
	user *models.User,
) error {
	err := dataBase.Create(user).Error

	return err
}

func (dataBase SqlDB) GetUsers(
	page int,
	perPage int,
	name string,
) (*[]models.User, error) {
	users := new([]models.User)

	if name == "" {
		if err := dataBase.
			Raw(utils.LimitQuery(
				"users",
				"deleted_at IS NULL",
				(page-1)*perPage, perPage+(page-1)*perPage,
			)).
			Scan(users).Error; err != nil {
			return nil, err
		}

		return users, nil
	}

	err := dataBase.Raw(utils.LimitQuery(
		"users",
		LikeQuery("users", "name", name),
		(page-1)*perPage, perPage+(page-1)*perPage)).
		Scan(users).Error

	return users, err
}

func (dataBase SqlDB) GetAllUsers() (*[]models.User, error) {
	users := new([]models.User)
	err := dataBase.Find(users).Error
	return users, err
}

func (dataBase SqlDB) GetUsersByEntry(name string) (*[]models.User, error) {
	users := new([]models.User)
	err := dataBase.Where("name LIKE ?", "%"+name+"%").Find(users).Error
	return users, err
}

func (dataBase SqlDB) GetUsersCount(query string) (int, error) {
	var count int

	err := dataBase.Table("users").
		Where(query).
		Count(&count).Error

	return count, err
}

func (dataBase SqlDB) GetUser(
	id int,
) (*models.User, error) {
	user := new(models.User)

	err := dataBase.Where(id).Find(user).Error

	return user, err
}

func (dataBase SqlDB) UpdateUser(
	user *models.User,
) error {
	err := dataBase.Save(user).Error

	return err
}

func (dataBase SqlDB) RemoveUser(
	user *models.User,
) error {
	err := dataBase.Delete(user).Error

	return err
}

func (dataBase SqlDB) AttachPoliciesByUser(
	user *models.User,
	policies *[]models.Policy,
) error {
	err := dataBase.Model(user).Association("policies").
		Append(policies).Error

	return err
}

func (dataBase SqlDB) DetachPolicyByUser(
	user *models.User,
	policy *models.Policy,
) error {
	err := dataBase.Model(user).Association("policies").
		Delete(policy).Error

	return err
}

func (dataBase SqlDB) GetPoliciesByUser(
	user *models.User,
) (*[]models.Policy, error) {
	policies := new([]models.Policy)

	err := dataBase.Model(user).Related(policies, "policies").Error

	return policies, err
}

func (dataBase SqlDB) GetGroupsByUser(
	user *models.User,
) (*[]models.Group, error) {
	groups := new([]models.Group)

	err := dataBase.Model(user).Related(groups, "groups").Error

	return groups, err
}
