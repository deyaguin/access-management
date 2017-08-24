package db

import (
	"gitlab/nefco/accessControl/app/models"
)

func (dataBase SqlDB) CreatePermission(permission *models.Permission) error {
	err := dataBase.db.Create(permission).Error
	return err
}

func (dataBase SqlDB) GetPermission(id int) (*models.Permission, error) {
	permissions := new(models.Permission)
	err := dataBase.db.Where(id).Find(permissions).Error
	return permissions, err
}

func (dataBase SqlDB) UpdatePermission(permission *models.Permission) error {
	err := dataBase.db.Save(permission).Error
	return err
}

func (dataBase SqlDB) DeletePermission(permission *models.Permission) error {
	err := dataBase.db.Where("id = ?", permission.ID).Delete(permission).Error
	return err
}
