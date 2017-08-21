package db

import (
	"fmt"
	"gitlab/nefco/accessControl/app/models"
)

func (dataBase SqlDB) CreateUser(u *models.User) (e error) {
	e = dataBase.db.Create(u).Error
	return e
}

func (dataBase SqlDB) GetUsers() (u *[]models.User, e error) {
	u = new([]models.User)
	e = dataBase.db.Find(u).Error
	return u, e
}

func (dataBase SqlDB) GetUser(key int) (u *models.User, e error) {
	u = new(models.User)
	e = dataBase.db.Where(key).Find(u).Error
	return u, e
}

func (dataBase SqlDB) UpdateUser(u *models.User) (e error) {
	e = dataBase.db.Save(u).Error
	return e
}

func (dataBase SqlDB) DeleteUser(u *models.User) (e error) {
	fmt.Println(u)
	e = dataBase.db.Where("id = ?", u.ID).Delete(u).Error
	return e
}
