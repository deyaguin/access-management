package storage

import "gitlab/nefco/access-management-system/src/models"

func (dataBase SqlDB) CreateAction(
	action *models.Action,
) error {
	err := dataBase.db.
		Create(action).Error

	return err
}

func (dataBase SqlDB) GetAction(
	id int,
) (*models.Action, error) {
	action := new(models.Action)

	err := dataBase.db.
		Where(id).Find(action).Error

	return action, err
}

func (dataBase SqlDB) GetActions() (*[]models.Action, error) {
	actions := new([]models.Action)
	err := dataBase.db.Find(actions).Error
	return actions, err
}

func (dataBase SqlDB) UpdateAction(
	action *models.Action,
) error {
	err := dataBase.db.
		Save(action).Error

	return err
}

func (dataBase SqlDB) RemoveAction(
	action *models.Action,
) error {
	err := dataBase.db.
		Delete(action).Error

	return err
}
