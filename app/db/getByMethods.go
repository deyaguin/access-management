package db

import "gitlab/nefco/accessControl/app/models"

func (dataBase SqlDB) GetPoliciesByUser(user *models.User, policies *[]models.Policy) error {
	err := dataBase.db.Model(user).Related(policies, "Policies").Error
	return err
}

func (dataBase SqlDB) GetGroupsByUser(user *models.User, groups *[]models.Group) error {
	err := dataBase.db.Model(user).Related(groups, "Groups").Error
	return err
}

func (dataBase SqlDB) GetUsersByGroup(group *models.Group, users *[]models.User) error {
	err := dataBase.db.Model(group).Related(users, "Users").Error
	return err
}

func (dataBase SqlDB) GetPoliciesByGroup(group *models.Group, policies *[]models.Policy) error {
	err := dataBase.db.Model(group).Related(policies, "Policies").Error
	return err
}

func (dataBase SqlDB) GetPermissionsByPolicy(policy *models.Policy, permissions *[]models.Permission) error {
	err := dataBase.db.Model(policy).Related(permissions, "Permissions").Error
	return err
}

func (dataBase SqlDB) GetUsersByPolicy(policy *models.Policy, users *[]models.User) error {
	err := dataBase.db.Model(policy).Related(users, "Users").Error
	return err
}

func (dataBase SqlDB) GetGroupsByPolicy(policy *models.Policy, groups *[]models.Group) error {
	err := dataBase.db.Model(policy).Related(groups, "Groups").Error
	return err
}
