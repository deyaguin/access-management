package db

import "gitlab/nefco/accessControl/app/models"

func (dataBase SqlDB) CreatePermission(p *models.Permission) (e error) {
	e = dataBase.db.Create(p).Error
	return e
}

func (dataBase SqlDB) GetPermission(key int) (p *models.Permission, e error) {
	p = new(models.Permission)
	e = dataBase.db.Where(key).Find(p).Error
	return p, e
}

func (dataBase SqlDB) UpdatePermission(p *models.Permission) (e error) {
	e = dataBase.db.Save(p).Error
	return e
}

func (dataBase SqlDB) DeletePermission(p *models.Permission) (e error) {
	e = dataBase.db.Delete(p).Error
	return e
}
