package storage

import (
	"gitlab/nefco/access-management-system/src/models"
)

func (dataBase SqlDB) CreateUser(user *models.User) error {
	return dataBase.db.Create(user).Error
}

func (dataBase SqlDB) GetUsers(page int) (*[]models.User, error) {
	users := new([]models.User)
	err := dataBase.db.Limit(ITEMS_ON_PAGE).Offset(page * ITEMS_ON_PAGE).Find(users).Error
	return users, err
}

func (dataBase SqlDB) GetUsersCount() (int, error) {
	var count int
	err := dataBase.db.Table("users").Count(&count).Error
	return count, err
}

func (dataBase SqlDB) GetUser(id int) (*models.User, error) {
	user := new(models.User)
	err := dataBase.db.Where(id).Find(user).Error
	return user, err
}

func (dataBase SqlDB) UpdateUser(user *models.User) error {
	return dataBase.db.Save(user).Error
}

func (dataBase SqlDB) RemoveUser(user *models.User) error {
	return dataBase.db.Where("id = ?", user.ID).Delete(user).Error
}
