package storage

import "gitlab/nefco/accessControl/app/models"

func (dataBase SqlDB) AddUsersToGroup(group *models.Group, users *[]models.User) error {
	err := dataBase.db.Model(group).Association("users").Append(users).Error
	return err
}

func (dataBase SqlDB) RemoveUserFromGroup(group *models.Group, user *models.User) error {
	err := dataBase.db.Model(group).Association("users").Delete(user).Error
	return err
}

func (dataBase SqlDB) AddPermissionsToPolicy(policy *models.Policy, permissions *[]models.Permission) error {
	err := dataBase.db.Model(policy).Association("permissions").Append(permissions).Error
	return err
}

func (dataBase SqlDB) AttachPoliciesByUser(user *models.User, policies *[]models.Policy) error {
	err := dataBase.db.Model(user).Association("policies").Append(policies).Error
	return err
}

func (dataBase SqlDB) DetachPolicyByUser(user *models.User, policy *models.Policy) error {
	err := dataBase.db.Model(user).Association("policies").Delete(policy).Error
	return err
}

func (dataBase SqlDB) AttachPoliciesByGroup(group *models.Group, policies *[]models.Policy) error {
	err := dataBase.db.Model(group).Association("policies").Append(policies).Error
	return err
}

func (dataBase SqlDB) DetachPolicyByGroup(group *models.Group, policy *models.Policy) error {
	err := dataBase.db.Model(group).Association("policies").Delete(policy).Error
	return err
}

func (dataBase SqlDB) GetPoliciesByUser(user *models.User) (*[]models.Policy, error) {
	policies := new([]models.Policy)
	err := dataBase.db.Model(user).Related(policies, "Policies").Error
	return policies, err
}

func (dataBase SqlDB) GetGroupsByUser(user *models.User) (*[]models.Group, error) {
	groups := new([]models.Group)
	err := dataBase.db.Model(user).Related(groups, "Groups").Error
	return groups, err
}

func (dataBase SqlDB) GetUsersByGroup(group *models.Group) (*[]models.User, error) {
	users := new([]models.User)
	err := dataBase.db.Model(group).Related(users, "Users").Error
	return users, err
}

func (dataBase SqlDB) GetPoliciesByGroup(group *models.Group) (*[]models.Policy, error) {
	policies := new([]models.Policy)
	err := dataBase.db.Model(group).Related(policies, "Policies").Error
	return policies, err
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
