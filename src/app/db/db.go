package db

import (
	"app/models"
	"github.com/jinzhu/gorm"
)

type DB interface {
	CreateUser(user *models.User)
	GetUsers(users *[]models.User)
}

type SqlDB struct {
	db *gorm.DB
}

func (dataBase *SqlDB) Connect(vendor, url string) {
	db, err := gorm.Open(vendor, url)
	if err != nil {
		panic(err)
	}
	dataBase.db = db
}

func (dataBase SqlDB) CreateUser(user *models.User) {
	dataBase.db.NewRecord(&user)
	dataBase.db.Create(&user)
}

func (dataBase SqlDB) GetUsers(users *[]models.User) {
	dataBase.db.Find(&users)
}
