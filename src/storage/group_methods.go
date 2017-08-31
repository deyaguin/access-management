package storage

import (
	"gitlab/nefco/access-management-system/src/models"
)

func (dataBase SqlDB) CreateGroup(
	group *models.Group,
) error {
	err := dataBase.db.
		Create(group).Error

	return err
}

func (dataBase SqlDB) GetGroups(
	page int,
	perPage int,
) (*[]models.Group, error) {
	groups := new([]models.Group)

	err := dataBase.db.
		Limit(perPage).Offset(page * perPage).
		Find(groups).Error

	return groups, err
}

func (dataBase SqlDB) GetGroupsCount() (int, error) {
	var count int

	err := dataBase.db.
		Table("groups").
		Count(&count).Error

	return count, err
}

func (dataBase SqlDB) GetGroup(
	id int,
) (*models.Group, error) {
	group := new(models.Group)

	err := dataBase.db.
		Where(id).Find(group).Error

	return group, err
}

func (dataBase SqlDB) UpdateGroup(
	group *models.Group,
) error {
	err := dataBase.db.
		Save(group).Error

	return err
}

func (dataBase SqlDB) RemoveGroup(
	group *models.Group,
) error {
	err := dataBase.db.
		Delete(group).Error

	return err
}

func (dataBase SqlDB) AddUsersToGroup(
	group *models.Group,
	users *[]models.User,
) error {
	err := dataBase.db.
		Model(group).Association("users").
		Append(users).Error

	return err
}

func (dataBase SqlDB) RemoveUserFromGroup(
	group *models.Group,
	user *models.User,
) error {
	err := dataBase.db.
		Model(group).Association("users").
		Delete(user).Error

	return err
}

func (dataBase SqlDB) GetUsersByGroup(
	group *models.Group,
	page *int,
	perPage *int,
) (*[]models.User, int, error) {
	users := new([]models.User)

	if page == nil || perPage == nil {
		err := dataBase.db.
			Model(group).Related(users, "users").Error
		if err != nil {
			return nil, 0, err
		}

		return users, 0, nil
	}

	err := dataBase.db.
		Limit(*perPage).Offset(*page**perPage).
		Model(group).Related(users, "users").Error

	count := dataBase.db.
		Model(group).Association("users").
		Count()

	return users, count, err
}

func (dataBase SqlDB) AttachPoliciesByGroup(
	group *models.Group,
	policies *[]models.Policy,
) error {
	err := dataBase.db.
		Model(group).Association("policies").
		Append(policies).Error

	return err
}

func (dataBase SqlDB) DetachPolicyByGroup(
	group *models.Group,
	policy *models.Policy,
) error {
	err := dataBase.db.
		Model(group).Association("policies").
		Delete(policy).Error

	return err
}

func (dataBase SqlDB) GetPoliciesByGroup(
	group *models.Group,
	page *int,
	perPage *int,
) (*[]models.Policy, int, error) {
	policies := new([]models.Policy)

	if page == nil || perPage == nil {
		err := dataBase.db.
			Model(group).Related(policies, "policies").Error
		if err != nil {
			return nil, 0, err
		}

		return policies, 0, nil
	}

	err := dataBase.db.
		Limit(*perPage).Offset(*page**perPage).
		Model(group).Related(policies, "policies").Error

	count := dataBase.db.
		Model(group).Association("policies").
		Count()

	return policies, count, err
}
