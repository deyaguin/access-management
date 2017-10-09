package services

import (
	"fmt"
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"regexp"
)

type CheckingParams struct {
	Resourse string `validation:"nonzero" json:"resourse"`
	Action   int    `validation:"nonzero" json:"action_id"`
	ID       int    `validation:"nonzero" json:"user_id"`
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

	return permissions, nil
}

func (service *permissionsCheckService) getGroupPermissions(
	user *models.User,
) ([]models.Permission, error) {
	var (
		policies    []models.Policy
		permissions []models.Permission
	)

	groups, err := service.storage.GetGroupsByUser(user)

	if err != nil {
		return nil, err
	}

	for _, group := range *groups {
		policy, err := service.storage.GetPoliciesByGroup(&group)
		policies = append(policies, *policy...)
		if err != nil {
			return nil, err
		}
	}

	for _, policy := range policies {
		permission, err := service.storage.GetPermissionsByPolicy(&policy)
		permissions = append(permissions, *permission...)
		if err != nil {
			return nil, err
		}
	}

	return permissions, nil
}

func (service *permissionsCheckService) comparePermissions(
	permissions *[]models.Permission,
	checkingParams *CheckingParams,
) (bool, bool) {
	result := false
	has := false
	for _, p := range *permissions {
		if *p.ActionID == checkingParams.Action && compareResourses(checkingParams.Resourse, *p.Resourse) {
			has = true
			result = *p.Access
			if !result {
				break
			}
		}
	}
	return result, has
}

func compareResourses(
	resourse1, resourse2 string,
) bool {
	if resourse1 == resourse2 {
		return true
	}

	if resourse2[len(resourse2)-1] == '*' {
		r, _ := regexp.Compile("^" + resourse2[:len(resourse2)-1])
		return r.MatchString(resourse1)
	}

	return false
}
