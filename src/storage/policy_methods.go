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
	policyName string,
) (*[]models.Policy, error) {
	policies := new([]models.Policy)

	if policyName == "" {
		if err := dataBase.Limit(perPage).Offset((page - 1) * perPage).
			Find(policies).Error; err != nil {
			return nil, err
		}

		return policies, nil
	}

	err := dataBase.Where("name = ?", policyName).
		Limit(perPage).Offset((page - 1) * perPage).Find(policies).Error

	return policies, err
}

func (dataBase SqlDB) GetAllPolicies() (*[]models.Policy, error) {
	policies := new([]models.Policy)
	err := dataBase.Find(policies).Error
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

func (dataBase SqlDB) AddPermissionToPolicy(
	policy *models.Policy,
	permission *models.Permission,
) error {
	err := dataBase.Model(policy).Association("permissions").
		Append(permission).Error

	return err
}

func (dataBase SqlDB) GetUsersByPolicy(
	policy *models.Policy,
) (*[]models.User, error) {
	users := new([]models.User)

	err := dataBase.Model(policy).Related(users, "users").Error

	return users, err
}

func (dataBase SqlDB) GetGroupsByPolicy(
	policy *models.Policy,
) (*[]models.Group, error) {
	groups := new([]models.Group)

	err := dataBase.Model(policy).Related(groups, "groups").Error

	return groups, err
}

func (dataBase SqlDB) GetPermissionsByPolicy(
	policy *models.Policy,
) (*[]models.Permission, error) {
	permissions := new([]models.Permission)

	err := dataBase.Model(policy).Related(permissions).Error

	return permissions, err
}
