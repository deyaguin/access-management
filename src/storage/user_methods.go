package storage

import (
	"gitlab/nefco/access-management-system/src/models"
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
) (*[]models.User, error) {
	users := new([]models.User)

	err := dataBase.Limit(perPage).Offset((page - 1) * perPage).Find(users).Error

	return users, err
}

func (dataBase SqlDB) GetUsersCount() (int, error) {
	var count int

	err := dataBase.Table("users").Count(&count).Error

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
	page *int,
	perPage *int,
) (*[]models.Policy, int, error) {
	policies := new([]models.Policy)

	if page == nil || perPage == nil {
		err := dataBase.Model(user).Related(policies, "policies").Error
		if err != nil {
			return nil, 0, err
		}

		return policies, 0, nil
	}

	err := dataBase.Limit(*perPage).Offset((*page-1)*(*perPage)).
		Model(user).Related(policies, "policies").Error

	count := dataBase.Model(user).Association("policies").Count()

	return policies, count, err
}

func (dataBase SqlDB) GetGroupsByUser(
	user *models.User,
	page *int,
	perPage *int,
) (*[]models.Group, int, error) {
	groups := new([]models.Group)

	if page == nil || perPage == nil {
		err := dataBase.Model(user).Related(groups, "groups").Error
		if err != nil {
			return nil, 0, err
		}

		return groups, 0, nil
	}

	err := dataBase.Limit(*perPage).Offset((*page-1)*(*perPage)).
		Model(user).Related(groups, "groups").Error

	count := dataBase.Model(user).Association("groups").Count()

	return groups, count, err
}
