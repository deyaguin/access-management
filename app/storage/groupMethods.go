package storage

import (
	"gitlab/nefco/accessControl/app/models"
)

func (dataBase SqlDB) CreateGroup(group *models.Group) error {
	err := dataBase.db.Create(group).Error
	return err
}

func (dataBase SqlDB) GetGroups() (*[]models.Group, error) {
	groups := new([]models.Group)
	err := dataBase.db.Find(groups).Error
	return groups, err
}

func (dataBase SqlDB) GetGroup(id int) (*models.Group, error) {
	group := new(models.Group)
	err := dataBase.db.Where(id).Find(group).Error
	return group, err
}

func (dataBase SqlDB) UpdateGroup(group *models.Group) error {
	err := dataBase.db.Save(group).Error
	return err
}

func (dataBase SqlDB) DeleteGroup(group *models.Group) error {
	err := dataBase.db.Where("id = ?", group.ID).Delete(group).Error
	return err
}
