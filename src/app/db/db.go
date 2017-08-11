package db

import (
	"github.com/jinzhu/gorm"
	"app/models"
)

type DB interface {
	CreateUser(u *models.User)
	GetUsers() *[]models.User

	CreateGroup(g *models.Group)
	GetGroups() *[]models.Group

	CreatePolicy(p  *models.Policy)
	GetPolicies() *[]models.Policy

	//Create(entity interface{})
	//GetAll(entities interface{})
	//Update(entity interface{})
	//Delete(entity interface{})

	//GetEntityAssociations(e interface{}, a interface{}, column string)

	GetPoliciesByUser(user *models.User, policies *[]models.Policy, column string)
	GetPoliciesByGroup(group *models.Group, policies *[]models.Policy, column string)
	GetGroupsByUser(user *models.User, groups *[]models.Group, column string)
	GetPermissionsByPolicy(policy *models.Policy, permissions *[]models.Permission, column string)
}

type SqlDB struct {
	db *gorm.DB
}

func (dataBase SqlDB) CreateUser(u *models.User) {
	dataBase.db.NewRecord(&u)
	dataBase.db.Create(&u)
}

func (dataBase SqlDB) GetUsers() *[]models.User {
	users := new([]models.User)
	dataBase.db.Find(&users)
	return users
}

func (dataBase SqlDB) CreateGroup(g *models.Group) {
	dataBase.db.NewRecord(&g)
	dataBase.db.Create(&g)
}

func (dataBase SqlDB) GetGroups() *[]models.Group {
	groups := new([]models.Group)
	dataBase.db.Find(&groups)
	return groups
}

func (dataBase SqlDB) CreatePolicy(p *models.Policy) {
	dataBase.db.NewRecord(&p)
	dataBase.db.Create(&p)
}

func (dataBase SqlDB) GetPolicies() *[]models.Policy {
	policies := new([]models.Policy)
	dataBase.db.Find(&policies)
	return policies
}

func (dataBase SqlDB) GetPoliciesByUser(user *models.User, policies *[]models.Policy, column string) {
	dataBase.db.Model(user).Related(policies, column)
}

func (dataBase SqlDB) GetPoliciesByGroup(group *models.Group, policies *[]models.Policy, column string) {
	dataBase.db.Model(group).Related(policies, column)
}

func (dataBase SqlDB) GetGroupsByUser(user *models.User, groups *[]models.Group, column string) {
	dataBase.db.Model(user).Related(groups, column)
}

func (dataBase SqlDB) GetPermissionsByPolicy(policy *models.Policy, permissions *[]models.Permission, column string) {
	dataBase.db.Model(policy).Related(permissions)
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
