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

func NewPermissionsCheckService(
	storage storage.DB,
) PermissionsCheckService {
	return &permissionsCheckService{
		storage,
	}
}

func (service *permissionsCheckService) CheckPermissions(
	checkingParams *CheckingParams,
) (bool, error) {
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

	return access, nil
}

func (service *permissionsCheckService) getUserPermissions(
	user *models.User,
) ([]models.Permission, error) {
	var permissions []models.Permission

	policies, _, err := service.storage.GetPoliciesByUser(user, nil, nil)
	if err == nil {
		for _, policy := range *policies {
			permission, _, err := service.storage.GetPermissionsByPolicy(&policy, nil, nil)
			permissions = append(permissions, *permission...)
			if err != nil {
				return permissions, err
			}
		}
	}

	return permissions, nil
}

func (service *permissionsCheckService) getGroupPermissions(
	user *models.User,
) ([]models.Permission, error) {
	var (
		policies    []models.Policy
		permissions []models.Permission
	)

	groups, _, err := service.storage.GetGroupsByUser(user, nil, nil)

	if err != nil {
		return nil, err
	}

	for _, group := range *groups {
		policy, _, err := service.storage.GetPoliciesByGroup(&group, nil, nil)
		policies = append(policies, *policy...)
		if err != nil {
			return nil, err
		}
	}

	for _, policy := range policies {
		permission, _, err := service.storage.GetPermissionsByPolicy(&policy, nil, nil)
		permissions = append(permissions, *permission...)
		if err != nil {
			return nil, err
		}
	}

	return permissions, nil
}

func (service *permissionsCheckService) comparePermissions(permissions *[]models.Permission, checkingParams *CheckingParams) (bool, bool) {
	result := false
	has := false
	for _, p := range *permissions {
		if *p.ActionID == checkingParams.Action && *p.Resourse == checkingParams.Resourse {
			has = true
			result = *p.Access
			if !result {
				break
			}
		}
	}
	return result, has
}
