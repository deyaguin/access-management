package db

import (
	"app/models"
	"github.com/jinzhu/gorm"
)

type DB interface {
	CreateUser(user *models.User)
	GetUsers(users *[]models.User)

	CreateGroup(group *models.Group)
	GetGroups(groups *[]models.Group)

	CreatePolicy(policy *models.Policy)
	GetPolicies(policies *[]models.Policy)
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

func (dataBase SqlDB) CreateGroup(group *models.Group) {
	dataBase.db.NewRecord(&group)
	dataBase.db.Create(&group)
}

func (dataBase SqlDB) GetGroups(groups *[]models.Group) {
	dataBase.db.Find(&groups)
}

func (dataBase SqlDB) CreatePolicy(policy *models.Policy) {
	dataBase.db.NewRecord(&policy)
	dataBase.db.Create(&policy)
}

func (dataBase SqlDB) GetPolicies(policies *[]models.Policy) {
	dataBase.db.Find(&policies)
}
