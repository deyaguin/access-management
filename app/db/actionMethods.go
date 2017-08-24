package db

import "gitlab/nefco/accessControl/app/models"

func (dataBase SqlDB) GetAction(id int) (*models.Action, error) {
	action := new(models.Action)
	err := dataBase.db.Where(id).Find(action).Error
	return action, err
}
