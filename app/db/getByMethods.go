package db

import "gitlab/nefco/accessControl/app/models"

func (dataBase SqlDB) GetPoliciesByUser(u *models.User, p *[]models.Policy) (e error) {
	e = dataBase.db.Model(u).Related(p, "Policies").Error
	return e
}

func (dataBase SqlDB) GetPoliciesByGroup(g *models.Group, p *[]models.Policy) (e error) {
	e = dataBase.db.Model(g).Related(p, "Policies").Error
	return e
}

func (dataBase SqlDB) GetGroupsByUser(u *models.User, g *[]models.Group) (e error) {
	e = dataBase.db.Model(u).Related(g, "Groups").Error
	return e
}

func (dataBase SqlDB) GetUsersByGroup(g *models.Group, u *[]models.User) (e error) {
	e = dataBase.db.Model(g).Related(u, "Users").Error
	return e
}

func (dataBase SqlDB) GetPermissionsByPolicy(pol *models.Policy, per *[]models.Permission) (e error) {
	e = dataBase.db.Model(pol).Related(per, "Permissions").Error
	return e
}
