package storage

import (
	"github.com/jinzhu/gorm"
	"gitlab/nefco/access-management-system/src/models"
)

type DB interface {
	CreateUser(*models.User) error
	GetUsers(int, int) (*[]models.User, error)
	GetUsersCount() (int, error)
	GetUser(int) (*models.User, error)
	UpdateUser(*models.User) error
	RemoveUser(*models.User) error
	AttachPoliciesByUser(*models.User, *[]models.Policy) error
	DetachPolicyByUser(*models.User, *models.Policy) error
	GetPoliciesByUser(*models.User) (*[]models.Policy, error)
	GetGroupsByUser(*models.User) (*[]models.Group, error)

	CreateGroup(*models.Group) error
	GetGroups(int, int) (*[]models.Group, error)
	GetGroupsCount() (int, error)
	GetGroup(int) (*models.Group, error)
	UpdateGroup(*models.Group) error
	RemoveGroup(*models.Group) error
	AddUsersToGroup(*models.Group, *[]models.User) error
	RemoveUserFromGroup(*models.Group, *models.User) error
	GetUsersByGroup(*models.Group) (*[]models.User, error)
	AttachPoliciesByGroup(*models.Group, *[]models.Policy) error
	DetachPolicyByGroup(*models.Group, *models.Policy) error
	GetPoliciesByGroup(*models.Group) (*[]models.Policy, error)

	CreatePolicy(*models.Policy) error
	GetPolicies(int, int) (*[]models.Policy, error)
	GetPoliciesCount() (int, error)
	GetPolicy(int) (*models.Policy, error)
	UpdatePolicy(*models.Policy) error
	RemovePolicy(*models.Policy) error
	AddPermissionsToPolicy(*models.Policy, *[]models.Permission) error

	CreatePermission(*models.Permission) error
	GetPermission(int) (*models.Permission, error)
	UpdatePermission(*models.Permission) error
	RemovePermission(*models.Permission) error
	GetPermissionsByPolicy(*models.Policy) (*[]models.Permission, error)
	GetUsersByPolicy(*models.Policy) (*[]models.User, error)
	GetGroupsByPolicy(*models.Policy) (*[]models.Group, error)

	CreateAction(*models.Action) error
	GetActions() (*[]models.Action, error)
	GetAction(id int) (*models.Action, error)
	UpdateAction(*models.Action) error
	RemoveAction(*models.Action) error
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
