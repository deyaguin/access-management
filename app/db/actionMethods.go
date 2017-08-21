package db

import "gitlab/nefco/accessControl/app/models"

func (dataBase SqlDB) GetAction(key int) (a *models.Action, e error) {
	a = new(models.Action)
	e = dataBase.db.Where(key).Find(a).Error
	return a, e
}
