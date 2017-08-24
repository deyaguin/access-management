package storage

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

	AddUsersToGroup(*models.Group, *[]models.User) error
	RemoveUserFromGroup(*models.Group, *models.User) error
	AddPermissionsToPolicy(*models.Policy, *[]models.Permission) error
	//RemovePermissionsFromPolicy(*models.Policy, *models.Permission) error
	AttachPoliciesByUser(*models.User, *[]models.Policy) error
	DetachPolicyByUser(*models.User, *models.Policy) error
	AttachPoliciesByGroup(*models.Group, *[]models.Policy) error
	DetachPolicyByGroup(*models.Group, *models.Policy) error

	GetPoliciesByUser(*models.User, *[]models.Policy) error
	GetPoliciesByGroup(*models.Group, *[]models.Policy) error
	GetGroupsByUser(*models.User, *[]models.Group) error
	GetUsersByGroup(*models.Group, *[]models.User) error
	GetPermissionsByPolicy(*models.Policy, *[]models.Permission) error
	GetUsersByPolicy(*models.Policy, *[]models.User) error
	GetGroupsByPolicy(*models.Policy, *[]models.Group) error
}

type SqlDB struct {
	db *gorm.DB
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
