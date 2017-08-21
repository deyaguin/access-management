package db

import "gitlab/nefco/accessControl/app/models"

func (dataBase SqlDB) CreateGroup(g *models.Group) (e error) {
	e = dataBase.db.Create(g).Error
	return e
}

func (dataBase SqlDB) GetGroups() (g *[]models.Group, e error) {
	g = new([]models.Group)
	e = dataBase.db.Find(g).Error
	return g, e
}

func (dataBase SqlDB) GetGroup(key int) (g *models.Group, e error) {
	g = new(models.Group)
	e = dataBase.db.Where(key).Find(g).Error
	return g, e
}

func (dataBase SqlDB) UpdateGroup(g *models.Group) (e error) {
	e = dataBase.db.Save(g).Error
	return e
}

func (dataBase SqlDB) DeleteGroup(g *models.Group) (e error) {
	e = dataBase.db.Delete(g).Error
	return e
}
