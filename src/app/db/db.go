package db

import (
	"github.com/jinzhu/gorm"
	"app/models"
)

type DB interface {
	CreateUser(u *models.User) error
	GetUsers() (*[]models.User, error)
	//UpdateUser(u *models.User) error
	DeleteUser(u *models.User) error

	CreateGroup(g *models.Group) error
	GetGroups() (*[]models.Group, error)
	//UpdateGroup(g *models.Group) error
	//DeleteGroup(g *models.Group) error

	CreatePolicy(p  *models.Policy) error
	GetPolicies() (*[]models.Policy, error)
	//UpdatePolicy(p *models.Policy) error
	//DeletePolicy(p *models.Policy) error

	GetPoliciesByUser(u *models.User, p *[]models.Policy, c string) error
	GetPoliciesByGroup(g *models.Group, p *[]models.Policy, c string) error
	GetGroupsByUser(u *models.User, g *[]models.Group, c string) error
	GetPermissionsByPolicy(pol *models.Policy, per *[]models.Permission, c string) error
}

type SqlDB struct {
	db *gorm.DB
}

func (dataBase SqlDB) CreateUser(u *models.User) error {
	//dataBase.db.NewRecord(&u)
	err := dataBase.db.Create(&u).Error
	return err
}

func (dataBase SqlDB) GetUsers() (*[]models.User, error) {
	users := new([]models.User)
	err := dataBase.db.Find(&users).Error
	return users, err
}

func (dataBase SqlDB) DeleteUser(u *models.User) error {
	err := dataBase.db.Delete(&u).Error
	return err
}

func (dataBase SqlDB) CreateGroup(g *models.Group) error {
	//dataBase.db.NewRecord(&g)
	err := dataBase.db.Create(&g).Error
	return err
}

func (dataBase SqlDB) GetGroups() (*[]models.Group, error) {
	groups := new([]models.Group)
	err := dataBase.db.Find(&groups).Error
	return groups, err
}

func (dataBase SqlDB) CreatePolicy(p *models.Policy) error {
	//dataBase.db.NewRecord(&p)
	err := dataBase.db.Create(&p).Error
	return err
}

func (dataBase SqlDB) GetPolicies() (*[]models.Policy, error) {
	policies := new([]models.Policy)
	err := dataBase.db.Find(&policies).Error
	return policies, err
}

func (dataBase SqlDB) GetPoliciesByUser(u *models.User, p *[]models.Policy, c string) error {
	err := dataBase.db.Model(u).Related(p, c).Error
	return err
}

func (dataBase SqlDB) GetPoliciesByGroup(g *models.Group, p *[]models.Policy, c string) error {
	err := dataBase.db.Model(g).Related(p, c).Error
	return err
}

func (dataBase SqlDB) GetGroupsByUser(u *models.User, g *[]models.Group, c string) error {
	err := dataBase.db.Model(u).Related(g, c).Error
	return err
}

func (dataBase SqlDB) GetPermissionsByPolicy(pol *models.Policy, per *[]models.Permission, c string) error {
	err := dataBase.db.Model(pol).Related(per).Error
	return err
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
