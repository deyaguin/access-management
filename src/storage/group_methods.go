package storage

import (
	"gitlab/nefco/access-management-system/src/models"
)

func (dataBase SqlDB) CreateGroup(
	group *models.Group,
) error {
	err := dataBase.Create(group).Error
	return err
}

func (dataBase SqlDB) GetGroups(
	page int,
	perPage int,
	name string,
) (*[]models.Group, error) {
	groups := new([]models.Group)

	if name == "" {
		if err := dataBase.Limit(perPage).Offset((page - 1) * perPage).
			Find(groups).Error; err != nil {
			return nil, err
		}

		return groups, nil
	}

	err := dataBase.Where("name LIKE ?", "%"+name+"%").
		Limit(perPage).Offset((page - 1) * perPage).Find(groups).Error

	return groups, err
}

func (dataBase SqlDB) GetAllGroups() (*[]models.Group, error) {
	groups := new([]models.Group)
	err := dataBase.Find(groups).Error
	return groups, err
}

func (dataBase SqlDB) GetGroupsByEntry(name string) (*[]models.Group, error) {
	groups := new([]models.Group)
	err := dataBase.Where("name LIKE ?", "%"+name+"%").Find(groups).Error
	return groups, err
}

func (dataBase SqlDB) GetGroupsCount() (int, error) {
	var count int

	err := dataBase.Table("groups").Count(&count).Error

	return count, err
}

func (dataBase SqlDB) GetGroup(
	id int,
) (*models.Group, error) {
	group := new(models.Group)

	err := dataBase.Where(id).Find(group).Error

	return group, err
}

func (dataBase SqlDB) UpdateGroup(
	group *models.Group,
) error {
	err := dataBase.Save(group).Error

	return err
}

func (dataBase SqlDB) RemoveGroup(
	group *models.Group,
) error {
	err := dataBase.Delete(group).Error

	return err
}

func (dataBase SqlDB) AddUsersToGroup(
	group *models.Group,
	users *[]models.User,
) error {
	err := dataBase.Model(group).Association("users").Append(users).Error

	return err
}

func (dataBase SqlDB) RemoveUserFromGroup(
	group *models.Group,
	user *models.User,
) error {
	err := dataBase.Model(group).Association("users").Delete(user).Error

	return err
}

func (dataBase SqlDB) GetUsersByGroup(
	group *models.Group,
) (*[]models.User, error) {
	users := new([]models.User)

	err := dataBase.Model(group).Related(users, "users").Error

	return users, err
}

func (dataBase SqlDB) AttachPoliciesByGroup(
	group *models.Group,
	policies *[]models.Policy,
) error {
	err := dataBase.Model(group).Association("policies").
		Append(policies).Error

	return err
}

func (dataBase SqlDB) DetachPolicyByGroup(
	group *models.Group,
	policy *models.Policy,
) error {
	err := dataBase.Model(group).Association("policies").
		Delete(policy).Error

	return err
}

func (dataBase SqlDB) GetPoliciesByGroup(
	group *models.Group,
) (*[]models.Policy, error) {
	policies := new([]models.Policy)

	err := dataBase.Model(group).Related(policies, "policies").Error

	return policies, err
}
