package db

import (
	"gitlab/nefco/accessControl/app/models"
)

func (dataBase SqlDB) CreateUser(user *models.User) error {
	err := dataBase.db.Create(user).Error
	return err
}

func (dataBase SqlDB) GetUsers() (*[]models.User, error) {
	users := new([]models.User)
	err := dataBase.db.Find(users).Error
	return users, err
}

func (dataBase SqlDB) GetUser(id int) (*models.User, error) {
	user := new(models.User)
	err := dataBase.db.Where(id).Find(user).Error
	return user, err
}

func (dataBase SqlDB) UpdateUser(user *models.User) error {
	err := dataBase.db.Save(user).Error
	return err
}

func (dataBase SqlDB) DeleteUser(user *models.User) error {
	err := dataBase.db.Where("id = ?", user.ID).Delete(user).Error
	return err
}
