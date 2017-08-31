package storage

import (
	"gitlab/nefco/access-management-system/src/models"
)

func (dataBase SqlDB) CreateUser(
	user *models.User,
) error {
	return dataBase.db.Create(user).Error
}

func (dataBase SqlDB) GetUsers(
	page int,
	perPage int,
) (*[]models.User, error) {
	users := new([]models.User)

	err := dataBase.db.
		Limit(perPage).Offset(page * perPage).
		Find(users).Error

	return users, err
}

func (dataBase SqlDB) GetUsersCount() (int, error) {
	var count int

	err := dataBase.db.
		Table("users").
		Count(&count).Error

	return count, err
}

func (dataBase SqlDB) GetUser(
	id int,
) (*models.User, error) {
	user := new(models.User)

	err := dataBase.db.
		Where(id).Find(user).Error

	return user, err
}

func (dataBase SqlDB) UpdateUser(
	user *models.User,
) error {
	return dataBase.db.
		Save(user).Error
}

func (dataBase SqlDB) RemoveUser(
	user *models.User,
) error {
	return dataBase.db.
		Delete(user).Error
}

func (dataBase SqlDB) AttachPoliciesByUser(
	user *models.User,
	policies *[]models.Policy,
) error {
	err := dataBase.db.
		Model(user).Association("policies").
		Append(policies).Error

	return err
}

func (dataBase SqlDB) DetachPolicyByUser(
	user *models.User,
	policy *models.Policy,
) error {
	err := dataBase.db.
		Model(user).Association("policies").
		Delete(policy).Error

	return err
}

//func (dataBase SqlDB) GetPoliciesByUser(
//	user *models.User,
//) (*[]models.Policy, error) {
//	policies := new([]models.Policy)
//
//	err := dataBase.db.
//		Model(user).Related(policies, "policies").Error
//
//	return policies, err
//}

func (dataBase SqlDB) GetPoliciesByUser(
	user *models.User,
	page *int,
	perPage *int,
) (*[]models.Policy, int, error) {
	policies := new([]models.Policy)

	if page == nil || perPage == nil {
		err := dataBase.db.
			Model(user).Related(policies, "policies").Error
		if err != nil {
			return nil, 0, err
		}

		return policies, 0, nil
	}

	err := dataBase.db.
		Limit(*perPage).Offset(*page * *perPage).
		Model(user).Related(policies, "policies").Error

	count := dataBase.db.
		Model(user).Association("policies").
		Count()

	return policies, count, err
}

//func (dataBase SqlDB) GetGroupsByUser(
//	user *models.User,
//) (*[]models.Group, error) {
//	groups := new([]models.Group)
//
//	err := dataBase.db.
//		Model(user).Related(groups, "groups").Error
//
//	return groups, err
//}

func (dataBase SqlDB) GetGroupsByUser(
	user *models.User,
	page *int,
	perPage *int,
) (*[]models.Group, int, error) {
	groups := new([]models.Group)

	if page == nil || perPage == nil {
		err := dataBase.db.
			Model(user).Related(groups, "groups").Error
		if err != nil {
			return nil, 0, err
		}

		return groups, 0, nil
	}

	err := dataBase.db.
		Limit(*perPage).Offset(*page * *perPage).
		Model(user).Related(groups, "groups").Error

	count := dataBase.db.
		Model(user).Association("groups").
		Count()

	return groups, count, err
}
