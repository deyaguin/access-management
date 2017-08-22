package db

import "gitlab/nefco/accessControl/app/models"

func (dataBase SqlDB) AddUsersToGroup(g *models.Group, u *[]models.User) (e error) {
	e = dataBase.db.Model(g).Association("users").Append(u).Error
	return e
}

func (dataBase SqlDB) RemoveUsersFromGroup(g *models.Group, u *[]models.User) (e error) {
	e = dataBase.db.Model(g).Association("users").Delete(u).Error
	return e
}

func (dataBase SqlDB) AddPermissionsToPolicy(pol *models.Policy, per *[]models.Permission) (e error) {
	e = dataBase.db.Model(pol).Association("permissions").Append(per).Error
	return e
}

func (dataBase SqlDB) RemovePermissionsFromPolicy(pol *models.Policy, per *[]models.Permission) (e error) {
	e = dataBase.db.Model(pol).Association("permissions").Delete(per).Error
	return e
}

func (dataBase SqlDB) AttachPoliciesByUser(u *models.User, p *[]models.Policy) (e error) {
	e = dataBase.db.Model(u).Association("policies").Append(p).Error
	return e
}

func (dataBase SqlDB) DetachPoliciesByUser(u *models.User, p *[]models.Policy) (e error) {
	e = dataBase.db.Model(u).Association("policies").Delete(p).Error
	return e
}

func (dataBase SqlDB) AttachPoliciesByGroup(g *models.Group, p *[]models.Policy) (e error) {
	e = dataBase.db.Model(g).Association("policies").Append(p).Error
	return e
}

func (dataBase SqlDB) DetachPoliciesByGroup(g *models.Group, p *[]models.Policy) (e error) {
	e = dataBase.db.Model(g).Association("policies").Delete(p).Error
	return e
}
