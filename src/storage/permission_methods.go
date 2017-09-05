package storage

import (
	"gitlab/nefco/access-management-system/src/models"
)

func (dataBase SqlDB) CreatePermission(
	permission *models.Permission,
) error {
	err := dataBase.Create(permission).Error
	return err
}

func (dataBase SqlDB) GetPermission(
	id int,
) (*models.Permission, error) {
	permissions := new(models.Permission)

	err := dataBase.Where(id).Find(permissions).Error

	return permissions, err
}

func (dataBase SqlDB) UpdatePermission(
	permission *models.Permission,
) error {
	err := dataBase.Save(permission).Error
	return err
}

func (dataBase SqlDB) RemovePermission(
	permission *models.Permission,
) error {
	err := dataBase.Delete(permission).Error
	return err
}
