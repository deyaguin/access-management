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
