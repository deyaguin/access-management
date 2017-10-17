package storage

import (
	"gitlab/nefco/access-management-system/src/logger"
	"gitlab/nefco/access-management-system/src/models"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var Log *zap.Logger = logger.NewLogger()

type DB interface {
	CreateUser(*models.User) error
	GetUsers(int, int, string) (*[]models.User, error)
	GetUsersByEntry(string) (*[]models.User, error)
	GetAllUsers() (*[]models.User, error)
	GetUsersCount() (int, error)
	GetUser(int) (*models.User, error)
	UpdateUser(*models.User) error
	RemoveUser(*models.User) error
	AttachPoliciesByUser(*models.User, *[]models.Policy) error
	DetachPolicyByUser(*models.User, *models.Policy) error
	GetPoliciesByUser(*models.User) (*[]models.Policy, error)
	GetGroupsByUser(*models.User) (*[]models.Group, error)

	CreateGroup(*models.Group) error
	GetGroups(int, int, string) (*[]models.Group, error)
	GetGroupsByEntry(string) (*[]models.Group, error)
	GetAllGroups() (*[]models.Group, error)
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
	GetPolicies(int, int, string) (*[]models.Policy, error)
	GetPoliciesByEntry(string) (*[]models.Policy, error)
	GetAllPolicies() (*[]models.Policy, error)
	GetPoliciesCount() (int, error)
	GetPolicy(int) (*models.Policy, error)
	UpdatePolicy(*models.Policy) error
	RemovePolicy(*models.Policy) error
	AddPermissionToPolicy(*models.Policy, *models.Permission) error
	CreatePermission(*models.Permission) error
	GetPermission(int) (*models.Permission, error)
	UpdatePermission(*models.Permission) error
	RemovePermission(*models.Permission) error
	GetPermissionsByPolicy(*models.Policy) (*[]models.Permission, error)
	GetUsersByPolicy(*models.Policy) (*[]models.User, error)
	GetGroupsByPolicy(*models.Policy) (*[]models.Group, error)

	CreateAction(*models.Action) error
	GetActions() (*[]models.Action, error)
	GetAction(int) (*models.Action, error)
	UpdateAction(*models.Action) error
	RemoveAction(*models.Action) error

	GetService(int) (*models.Service, error)
	GetAllServices() (*[]models.Service, error)
	GetActionsByService(*models.Service) (*[]models.Action, error)
}

type SqlDB struct {
	*gorm.DB
}

func SqlDBCreator(vendor, url string) DB {
	db, err := gorm.Open(vendor, url)
	if err != nil {
		Log.Error(
			"DB start failed",
			zap.Error(err),
		)
	}

	db.LogMode(true)

	Log.Info("DB start successfully")

	DB := &SqlDB{
		db,
	}

	return DB
}
