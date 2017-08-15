package db

import (
	"github.com/jinzhu/gorm"
	"app/models"
)

type DB interface {
	CreateUser(*models.User) error
	GetUsers() (*[]models.User, error)
	//UpdateUser(*models.User) error
	DeleteUser(*models.User) error

	CreateGroup(*models.Group) error
	GetGroups() (*[]models.Group, error)
	//UpdateGroup(*models.Group) (err error)
	//DeleteGroup(*models.Group) (err error)

	CreatePolicy(*models.Policy) error
	GetPolicies() (*[]models.Policy, error)
	//UpdatePolicy(*models.Policy) (err error)
	//DeletePolicy(*models.Policy) (err error)

	CreatePermission(*models.Permission) error
	//UpdatePermission(*models.Permission) error
	//DeletePermission(*models.Permission) error

	//AttachPolicyToUser(*[]models.User, *[]models.Policy) error
	//AttachPolicyTOGroup(*[]models.Group, *[]models.Policy) error
	//AddUsersToGroups(*[]models.User, *[]models.Group) error
	//
	//AttachPolicyToUsers(*models.Policy, *[]models.User) error
	//AttachPolicyToGroups(*models.Policy, *[]models.Group) error
	//AttachPoliciesByUser(*models.User, *[]models.Policy) error
	//AttachPoliciesByGroup(*models.Group, *[]models.Policy) error
	//AddUserToGroups(*models.User, *[]models.Group) error
	//AddUsersToGroup(*models.Group, *[]models.User) error

	GetPoliciesByUser(*models.User, *[]models.Policy, string) error
	GetPoliciesByGroup(*models.Group, *[]models.Policy, string) error
	GetGroupsByUser(*models.User, *[]models.Group, string) error
	GetPermissionsByPolicy(*models.Policy, *[]models.Permission, string) error
}

type SqlDB struct {
	db *gorm.DB
}

func (dataBase SqlDB) CreateUser(u *models.User) (e error) {
	e = dataBase.db.Create(&u).Error
	return e
}

func (dataBase SqlDB) GetUsers() (u *[]models.User, e error) {
	u = new([]models.User)
	e = dataBase.db.Find(&u).Error
	return u, e
}

func (dataBase SqlDB) DeleteUser(u *models.User) (e error) {
	e = dataBase.db.Delete(&u).Error
	return e
}

func (dataBase SqlDB) CreateGroup(g *models.Group) (e error) {
	e = dataBase.db.Create(&g).Error
	return e
}

func (dataBase SqlDB) GetGroups() (g *[]models.Group, e error) {
	g = new([]models.Group)
	e = dataBase.db.Find(&g).Error
	return g, e
}

func (dataBase SqlDB) CreatePolicy(p *models.Policy) (e error) {
	e = dataBase.db.Create(&p).Error
	return e
}

func (dataBase SqlDB) GetPolicies() (p *[]models.Policy, e error) {
	p = new([]models.Policy)
	e = dataBase.db.Find(&p).Error
	return p, e
}

func (dataBase SqlDB) CreatePermission(p *models.Permission) (e error) {
	e = dataBase.db.Create(&p).Error
	return e
}

func (dataBase SqlDB) GetPoliciesByUser(u *models.User, p *[]models.Policy, c string) (e error) {
	e = dataBase.db.Model(u).Related(p, c).Error
	return e
}

func (dataBase SqlDB) GetPoliciesByGroup(g *models.Group, p *[]models.Policy, c string) (e error) {
	e = dataBase.db.Model(g).Related(p, c).Error
	return e
}

func (dataBase SqlDB) GetGroupsByUser(u *models.User, g *[]models.Group, c string) (e error) {
	e = dataBase.db.Model(u).Related(g, c).Error
	return e
}

func (dataBase SqlDB) GetPermissionsByPolicy(pol *models.Policy, per *[]models.Permission, c string) (e error) {
	e = dataBase.db.Model(pol).Related(per).Error
	return e
}

func SqlDBCreator(vendor, url string) (DB, error) {
	db, err := gorm.Open(vendor, url)
	if err != nil {
		return nil, err
	}
	DB := &SqlDB{
		db,
	}
	return DB, err
}
