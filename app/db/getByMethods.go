package db

import "gitlab/nefco/accessControl/app/models"

func (dataBase SqlDB) GetPoliciesByUser(u *models.User, p *[]models.Policy, c string) (e error) {
	e = dataBase.db.Model(u).Related(p, c).Error
	return e
}

func (dataBase SqlDB) GetPoliciesByGroup(g *models.Group, p *[]models.Policy, c string) (e error) {
	e = dataBase.db.Model(g).Related(p, c).Error
	return e
}

func (dataBase SqlDB) GetGroupsByUser(u *models.User, g *[]models.Group, c string) (e error) {
	e = dataBase.db.Model(u).Related(g, c).Error
	return e
}

func (dataBase SqlDB) GetPermissionsByPolicy(pol *models.Policy, per *[]models.Permission, c string) (e error) {
	e = dataBase.db.Model(pol).Related(per).Error
	return e
}
