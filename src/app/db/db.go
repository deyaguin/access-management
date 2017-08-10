package db

import (
	"github.com/jinzhu/gorm"
	"app/models"
)

type DB interface {
	CreateUser(user *models.User)
	GetUsers() *[]models.User

	CreateGroup(group *models.Group)
	GetGroups() *[]models.Group

	CreatePolicy(policy *models.Policy)
	GetPolicies() *[]models.Policy

	//Create(entity interface{})
	//GetAll(entities interface{})
	//Update(entity interface{})
	//Delete(entity interface{})

	GetEntityAssociations(entity interface{}, associations interface{}, column string)
}

type SqlDB struct {
	db *gorm.DB
}

func (dataBase SqlDB) CreateUser(user *models.User) {
	dataBase.db.NewRecord(&user)
	dataBase.db.Create(&user)
}

func (dataBase SqlDB) GetUsers() *[]models.User {
	users := new([]models.User)
	dataBase.db.Find(&users)
	return users
}

func (dataBase SqlDB) CreateGroup(group *models.Group) {
	dataBase.db.NewRecord(&group)
	dataBase.db.Create(&group)
}

func (dataBase SqlDB) GetGroups() *[]models.Group {
	groups := new([]models.Group)
	dataBase.db.Find(&groups)
	return groups
}

func (dataBase SqlDB) CreatePolicy(policy *models.Policy) {
	dataBase.db.NewRecord(&policy)
	dataBase.db.Create(&policy)
}

func (dataBase SqlDB) GetPolicies() *[]models.Policy {
	policies := new([]models.Policy)
	dataBase.db.Find(&policies)
	return policies
}

func (dataBase SqlDB) GetEntityAssociations(entity interface{}, associations interface{}, column string) {
	dataBase.db.Model(entity).Related(associations, column)
}

func SqlDBCreator(vendor, url string) DB {
	db, err := gorm.Open(vendor, url)
	if err != nil {
		panic(err)
	}
	DB := &SqlDB{
		db,
	}
	return DB
}

//func (dataBase SqlDB) Create(entity interface{}) {
//	dataBase.db.NewRecord(entity)
//	dataBase.db.Create(entity)
//}
//
//func (dataBase SqlDB) GetAll(entities interface{}) {
//	dataBase.db.Find(entities)
//}


//func (dataBase *SqlDB) Connect(vendor, url string) {
//	db, err := gorm.Open(vendor, url)
//	if err != nil {
//		panic(err)
//	}
//	dataBase.db = db
//}
