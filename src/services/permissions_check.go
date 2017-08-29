package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
)

type CheckingParams struct {
	Resourse string `validation:"nonzero"`
	Action   int    `validation:"nonzero"`
	ID       int    `validation:"nonzero"`
}

type PermissionsCheckService interface {
	getUserPermissions(*models.User) ([]models.Permission, error)
	getGroupPermissions(*models.User) ([]models.Permission, error)
	comparePermissions(*[]models.Permission, *CheckingParams) (bool, bool)
	CheckPermissions(checkingParams *CheckingParams) (bool, error)
}

type permissionsCheckService struct {
	storage storage.DB
}

func NewPermissionsCheckService(storage storage.DB) PermissionsCheckService {
	return &permissionsCheckService{
		storage,
	}
}

func (service *permissionsCheckService) CheckPermissions(checkingParams *CheckingParams) (bool, error) {
	user := &models.User{ID: checkingParams.ID}
	userPermissions, err := service.getUserPermissions(user)
	if err != nil {
		return false, err
	}

	groupPermissions, err := service.getGroupPermissions(user)
	if err != nil {
		return false, err
	}

	access, has := service.comparePermissions(&userPermissions, checkingParams)
	if !has {
		access, _ = service.comparePermissions(&groupPermissions, checkingParams)
	}

	return access, err
}

func (service *permissionsCheckService) getUserPermissions(user *models.User) ([]models.Permission, error) {
	var permissions []models.Permission

	policies, err := service.storage.GetPoliciesByUser(user)
	if err == nil {
		for _, policy := range *policies {
			permission, err := service.storage.GetPermissionsByPolicy(&policy)
			permissions = append(permissions, *permission...)
			if err != nil {
				return permissions, err
			}
		}
	}

	return permissions, err
}

func (service *permissionsCheckService) getGroupPermissions(user *models.User) ([]models.Permission, error) {
	var (
		policies    []models.Policy
		permissions []models.Permission
	)

	groups, err := service.storage.GetGroupsByUser(user)

	for _, group := range *groups {
		policy, err := service.storage.GetPoliciesByGroup(&group)
		policies = append(policies, *policy...)
		if err != nil {
			return permissions, err
		}
	}

	for _, policy := range policies {
		permission, err := service.storage.GetPermissionsByPolicy(&policy)
		permissions = append(permissions, *permission...)
		if err != nil {
			return permissions, err
		}
	}

	return permissions, err
}

func (service *permissionsCheckService) comparePermissions(permissions *[]models.Permission, checkingParams *CheckingParams) (bool, bool) {
	result := false
	has := false
	for _, p := range *permissions {
		if p.ActionID == checkingParams.Action && p.Resourse == checkingParams.Resourse {
			has = true
			result = p.Access
			if !result {
				break
			}
		}
	}
	return result, has
}
