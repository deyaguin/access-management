package storage

import (
	"github.com/jinzhu/gorm"
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/logger"
	"go.uber.org/zap"
)

var Log *zap.Logger = logger.NewLogger()

type DB interface {
	CreateUser(*models.User) error
	GetUsers(int, int) (*[]models.User, error)
	GetUsersCount() (int, error)
	GetUser(int) (*models.User, error)
	UpdateUser(*models.User) error
	RemoveUser(*models.User) error
	AttachPoliciesByUser(*models.User, *[]models.Policy) error
	DetachPolicyByUser(*models.User, *models.Policy) error
	GetPoliciesByUser(*models.User, *int, *int) (*[]models.Policy, int, error)
	GetGroupsByUser(*models.User, *int, *int) (*[]models.Group, int, error)

	CreateGroup(*models.Group) error
	GetGroups(int, int) (*[]models.Group, error)
	GetGroupsCount() (int, error)
	GetGroup(int) (*models.Group, error)
	UpdateGroup(*models.Group) error
	RemoveGroup(*models.Group) error
	AddUsersToGroup(*models.Group, *[]models.User) error
	RemoveUserFromGroup(*models.Group, *models.User) error
	GetUsersByGroup(*models.Group, *int, *int) (*[]models.User, int, error)
	AttachPoliciesByGroup(*models.Group, *[]models.Policy) error
	DetachPolicyByGroup(*models.Group, *models.Policy) error
	GetPoliciesByGroup(*models.Group, *int, *int) (*[]models.Policy, int, error)

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
	GetPermissionsByPolicy(*models.Policy, *int, *int) (*[]models.Permission, int, error)
	GetUsersByPolicy(*models.Policy, *int, *int) (*[]models.User, int, error)
	GetGroupsByPolicy(*models.Policy, *int, *int) (*[]models.Group, int, error)

	CreateAction(*models.Action) error
	GetActions() (*[]models.Action, error)
	GetAction(int) (*models.Action, error)
	UpdateAction(*models.Action) error
	RemoveAction(*models.Action) error
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

	Log.Info("DB start successfully")

	DB := &SqlDB{
		db,
	}

	return DB
}
