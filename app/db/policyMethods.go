package db

import "gitlab/nefco/accessControl/app/models"

func (dataBase SqlDB) CreatePolicy(p *models.Policy) (e error) {
	e = dataBase.db.Create(p).Error
	return e
}

func (dataBase SqlDB) GetPolicies() (p *[]models.Policy, e error) {
	p = new([]models.Policy)
	e = dataBase.db.Find(p).Error
	return p, e
}

func (dataBase SqlDB) GetPolicy(key int) (p *models.Policy, e error) {
	p = new(models.Policy)
	e = dataBase.db.Where(key).Find(p).Error
	return p, e
}

func (dataBase SqlDB) UpdatePolicy(p *models.Policy) (e error) {
	e = dataBase.db.Save(p).Error
	return e
}

func (dataBase SqlDB) DeletePolicy(p *models.Policy) (e error) {
	e = dataBase.db.Delete(p).Error
	return e
}
