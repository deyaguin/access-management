package storage

import (
	"gitlab/nefco/access-management-system/src/models"
)

func (dataBase SqlDB) CreatePolicy(policy *models.Policy) error {
	err := dataBase.db.Create(policy).Error
	return err
}

func (dataBase SqlDB) GetPolicies() (*[]models.Policy, error) {
	policies := new([]models.Policy)
	err := dataBase.db.Find(policies).Error
	return policies, err
}

func (dataBase SqlDB) GetPolicy(id int) (*models.Policy, error) {
	policy := new(models.Policy)
	err := dataBase.db.Where(id).Find(policy).Error
	return policy, err
}

func (dataBase SqlDB) UpdatePolicy(policy *models.Policy) error {
	err := dataBase.db.Save(policy).Error
	return err
}

func (dataBase SqlDB) RemovePolicy(policy *models.Policy) error {
	err := dataBase.db.Where("id = ?", policy.ID).Delete(policy).Error
	return err
}

func (dataBase SqlDB) AddPermissionsToPolicy(policy *models.Policy, permissions *[]models.Permission) error {
	err := dataBase.db.Model(policy).Association("permissions").Append(permissions).Error
	return err
}

func (dataBase SqlDB) GetUsersByPolicy(policy *models.Policy) (*[]models.User, error) {
	users := new([]models.User)
	err := dataBase.db.Model(policy).Related(users, "Users").Error
	return users, err
}

func (dataBase SqlDB) GetGroupsByPolicy(policy *models.Policy) (*[]models.Group, error) {
	groups := new([]models.Group)
	err := dataBase.db.Model(policy).Related(groups, "Groups").Error
	return groups, err
}

func (dataBase SqlDB) GetPermissionsByPolicy(policy *models.Policy) (*[]models.Permission, error) {
	permissions := new([]models.Permission)
	err := dataBase.db.Model(policy).Related(permissions, "Permissions").Error
	return permissions, err
}
