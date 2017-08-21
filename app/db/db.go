package db

import (
	"github.com/jinzhu/gorm"
	"gitlab/nefco/accessControl/app/models"
)

type DB interface {
	CreateUser(*models.User) error
	GetUsers() (*[]models.User, error)
	GetUser(int) (*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*models.User) error

	CreateGroup(*models.Group) error
	GetGroups() (*[]models.Group, error)
	GetGroup(int) (*models.Group, error)
	UpdateGroup(*models.Group) error
	DeleteGroup(*models.Group) error

	CreatePolicy(*models.Policy) error
	GetPolicies() (*[]models.Policy, error)
	GetPolicy(int) (*models.Policy, error)
	UpdatePolicy(*models.Policy) error
	DeletePolicy(*models.Policy) error

	CreatePermission(*models.Permission) error
	GetPermission(int) (*models.Permission, error)
	UpdatePermission(*models.Permission) error
	DeletePermission(*models.Permission) error

	//CreateAction(*models.Action) error
	//GetActions(*[]models.Action) error
	GetAction(id int) (*models.Action, error)
	//UpdateAction(*models.Action) error
	//DeleteAction(*models.Action) error

	//AddPermissionsToPolicy(*models.Policy) error
	AddUserToGroups(*models.User) error
	AddUsersToGroup(*models.Group) error
	AttachPolicyToUsers(*models.Policy) error
	AttachPolicyToGroups(*models.Policy) error
	AttachPoliciesByUser(*models.User) error
	AttachPoliciesByGroup(*models.Group) error
	RemoveUserFromGroups(*models.User) error
	RemoveUsersFromGroup(*models.Group) error
	DetachPoliciesByUser(*models.User) error
	DetachPoliciesByGroup(*models.Group) error
	DetachUsersFromPolicy(policy *models.Policy) error
	DetachGroupsFromPolicy(policy *models.Policy) error

	GetPoliciesByUser(*models.User, *[]models.Policy, string) error
	GetPoliciesByGroup(*models.Group, *[]models.Policy, string) error
	GetGroupsByUser(*models.User, *[]models.Group, string) error
	GetPermissionsByPolicy(*models.Policy, *[]models.Permission, string) error
}

type SqlDB struct {
	db *gorm.DB
}

///user



///group



///policy



///permission



///action


///getBy



///associations

//func (dataBase SqlDB) AddPermissionsToPolicy(p *models.Policy) (e error) {
//	e = dataBase.db.Model(p).Association("permissions").Append(p.Permissions).Error
//	return e
//}

func (dataBase SqlDB) AddUserToGroups(u *models.User) (e error) {
	e = dataBase.db.Model(u).Association("groups").Append(u.Groups).Error
	return e
}

func (dataBase SqlDB) AddUsersToGroup(g *models.Group) (e error) {
	e = dataBase.db.Model(g).Association("users").Append(g.Users).Error
	return e
}

func (dataBase SqlDB) AttachPolicyToUsers(p *models.Policy) (e error) {
	e = dataBase.db.Model(p).Association("users").Append(p.Users).Error
	return e
}

func (dataBase SqlDB) AttachPolicyToGroups(p *models.Policy) (e error) {
	e = dataBase.db.Model(p).Association("groups").Append(p.Groups).Error
	return e
}

func (dataBase SqlDB) AttachPoliciesByUser(u *models.User) (e error) {
	e = dataBase.db.Model(u).Association("policies").Append(u.Policies).Error
	return e
}

func (dataBase SqlDB) AttachPoliciesByGroup(g *models.Group) (e error) {
	e = dataBase.db.Model(g).Association("policies").Append(g.Policies).Error
	return e
}

func (dataBase SqlDB) RemoveUserFromGroups(u *models.User) (e error) {
	e = dataBase.db.Model(u).Association("groups").Delete(u.Groups).Error
	return e
}

func (dataBase SqlDB) RemoveUsersFromGroup(g *models.Group) (e error) {
	e = dataBase.db.Model(g).Association("users").Delete(g.Users).Error
	return e
}

func (dataBase SqlDB) DetachPoliciesByUser(u *models.User) (e error) {
	e = dataBase.db.Model(u).Association("policies").Delete(u.Policies).Error
	return e
}

func (dataBase SqlDB) DetachPoliciesByGroup(g *models.Group) (e error) {
	e = dataBase.db.Model(g).Association("policies").Delete(g.Policies).Error
	return e
}

func (dataBase SqlDB) DetachUsersFromPolicy(p *models.Policy) (e error) {
	e = dataBase.db.Model(p).Association("users").Delete(p.Users).Error
	return e
}

func (dataBase SqlDB) DetachGroupsFromPolicy(p *models.Policy) (e error) {
	e = dataBase.db.Model(p).Association("groups").Delete(p.Groups).Error
	return e
}

func SqlDBCreator(vendor, url string) (DB, error) {
	db, err := gorm.Open(vendor, url)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	DB := &SqlDB{
		db,
	}
	return DB, err
}
