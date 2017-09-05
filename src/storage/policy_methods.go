package storage

import (
	"gitlab/nefco/access-management-system/src/models"
)

func (dataBase SqlDB) CreatePolicy(
	policy *models.Policy,
) error {
	err := dataBase.Create(policy).Error
	return err
}

func (dataBase SqlDB) GetPolicies(
	page int,
	perPage int,
) (*[]models.Policy, error) {
	policies := new([]models.Policy)

	err := dataBase.Limit(perPage).Offset((page - 1) * perPage).Find(policies).Error

	return policies, err
}

func (dataBase SqlDB) GetPoliciesCount() (int, error) {
	var count int

	err := dataBase.Table("policies").Count(&count).Error

	return count, err
}

func (dataBase SqlDB) GetPolicy(
	id int,
) (*models.Policy, error) {
	policy := new(models.Policy)

	err := dataBase.Where(id).Find(policy).Error

	return policy, err
}

func (dataBase SqlDB) UpdatePolicy(
	policy *models.Policy,
) error {
	err := dataBase.Save(policy).Error

	return err
}

func (dataBase SqlDB) RemovePolicy(
	policy *models.Policy,
) error {
	err := dataBase.Delete(policy).Error

	return err
}

func (dataBase SqlDB) AddPermissionsToPolicy(
	policy *models.Policy,
	permissions *[]models.Permission,
) error {
	err := dataBase.Model(policy).Association("permissions").
		Append(permissions).Error

	return err
}

func (dataBase SqlDB) GetUsersByPolicy(
	policy *models.Policy,
	page *int,
	perPage *int,
) (*[]models.User, int, error) {
	users := new([]models.User)

	if page == nil || perPage == nil {
		err := dataBase.Model(policy).Related(users, "users").Error
		if err != nil {
			return nil, 0, err
		}

		return users, 0, nil
	}

	err := dataBase.Limit(*perPage).Offset((*page - 1) * (*perPage)).
		Model(policy).Related(users, "users").Error

	count := dataBase.Model(policy).Association("users").Count()

	return users, count, err
}

func (dataBase SqlDB) GetGroupsByPolicy(
	policy *models.Policy,
	page *int,
	perPage *int,
) (*[]models.Group, int, error) {
	groups := new([]models.Group)

	if page == nil || perPage == nil {
		err := dataBase.Model(policy).Related(groups, "groups").Error
		if err != nil {
			return nil, 0, err
		}

		return groups, 0, nil
	}

	err := dataBase.Limit(*perPage).Offset((*page - 1) * (*perPage)).
		Model(policy).Related(groups, "groups").Error

	count := dataBase.Model(policy).Association("groups").Count()

	return groups, count, err
}

func (dataBase SqlDB) GetPermissionsByPolicy(
	policy *models.Policy,
	page *int,
	perPage *int,
) (*[]models.Permission, int, error) {
	permissions := new([]models.Permission)

	if page == nil || perPage == nil {
		err := dataBase.Model(policy).Related(permissions, "permissions").Error
		if err != nil {
			return nil, 0, err
		}

		return permissions, 0, nil
	}

	err := dataBase.Limit(*perPage).Offset((*page - 1) * (*perPage)).
		Model(policy).Related(permissions, "permissions").Error

	count := dataBase.Model(policy).Association("permissions").Count()

	return permissions, count, err
}
