package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"

	"gopkg.in/validator.v2"
)

type PoliciesService interface {
	CreatePolicy(*models.Policy) (*models.Policy, error)
	GetPolicy(int) (*models.Policy, error)
	GetPolicies(int, int, string) (*items, error)
	UpdatePolicy(*models.Policy) (*models.Policy, error)
	RemovePolicy(int) error

	AddPermissionsToPolicy(*models.Policy, *[]models.Permission) error
	RemovePermissionFromPolicy(int, int) error
	GetUsersByPolicy(int) (*pureItems, error)
	GetGroupsByPolicy(int) (*pureItems, error)
	GetPermissionsByPolicy(int) (*pureItems, error)
}

type policiesService struct {
	storage storage.DB
}

func NewPoliciesService(
	storage storage.DB,
) PoliciesService {
	return &policiesService{
		storage,
	}
}

func (service *policiesService) CreatePolicy(
	policyCreating *models.Policy,
) (*models.Policy, error) {
	if err := validator.Validate(policyCreating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	policy := new(models.Policy)
	policy.SetFields(policyCreating)

	if err := service.storage.CreatePolicy(policy); err != nil {
		return nil, NewEntityCreateError(err.Error())
	}
	return policy, nil
}

func (service *policiesService) GetPolicy(
	policyID int,
) (*models.Policy, error) {
	policy, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", policyID)
	}

	return policy, nil
}

func (service *policiesService) GetPolicies(
	page int,
	perPage int,
	name string,
) (*items, error) {
	policies, err := service.storage.GetPolicies(page, perPage, name)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	count, err := service.storage.GetPoliciesCount()
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	response := &items{
		policies,
		count,
		perPage,
		page,
	}

	return response, nil
}

func (service *policiesService) UpdatePolicy(
	policyUpdating *models.Policy,
) (*models.Policy, error) {
	policy, err := service.storage.GetPolicy(policyUpdating.ID)
	if err != nil {
		return nil, NewEntityNotFoundError("permission", policyUpdating.ID)
	}

	if err := validator.Validate(policyUpdating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	policy.SetFields(policyUpdating)
	if err := service.storage.UpdatePolicy(policy); err != nil {
		return nil, NewEntityUpdateError(err.Error())
	}

	return policy, nil
}

func (service *policiesService) RemovePolicy(
	policyID int,
) error {
	policy, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return NewEntityNotFoundError("policy", policyID)
	}

	if err := service.storage.RemovePolicy(policy); err != nil {
		return NewEntityRemoveError(err.Error())
	}

	return nil
}

func (service *policiesService) AddPermissionsToPolicy(
	policy *models.Policy,
	permissions *[]models.Permission,
) error {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewEntityNotFoundError("policy", policy.ID)
	}

	for _, permission := range *permissions {
		if err := validator.Validate(permission); err != nil {
			if err != nil {
				return NewValidationError(err.Error())
			}

			return NewEntityCreateError(err.Error())
		}
	}

	if err := service.storage.AddPermissionsToPolicy(
		policy,
		permissions,
	); err != nil {
		return NewEntityCreateError(err.Error())
	}

	return nil
}

func (service *policiesService) RemovePermissionFromPolicy(
	policyID int,
	permissionID int,
) error {
	_, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return NewEntityNotFoundError("policy", policyID)
	}

	permission, err := service.storage.GetPermission(permissionID)
	if err != nil {
		return NewEntityNotFoundError("permission", permissionID)
	}

	if err := service.storage.RemovePermission(permission); err != nil {
		return NewEntityRemoveError(err.Error())
	}

	return nil
}

func (service *policiesService) GetPermissionsByPolicy(
	policyID int,
) (*pureItems, error) {
	policy, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", policyID)
	}

	permissions, err := service.storage.GetPermissionsByPolicy(policy)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	result := &pureItems{permissions}

	return result, nil
}

func (service *policiesService) GetUsersByPolicy(
	policyID int,
) (*pureItems, error) {
	policy, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", policyID)
	}

	users, err := service.storage.GetUsersByPolicy(policy)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	result := &pureItems{users}

	return result, nil
}

func (service *policiesService) GetGroupsByPolicy(
	policyID int,
) (*pureItems, error) {
	policy, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", policyID)
	}

	groups, err := service.storage.GetGroupsByPolicy(policy)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	result := &pureItems{groups}

	return result, nil
}
