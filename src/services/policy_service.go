package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type PolicyService interface {
	CreatePolicy(*models.Policy) (*models.Policy, error)
	GetPolicy(int) (*models.Policy, error)
	GetPolicies(int, int) (*items, error)
	UpdatePolicy(*models.Policy) (*models.Policy, error)
	RemovePolicy(int) error

	AddPermissionsToPolicy(*models.Policy, *[]models.Permission) error
	RemovePermissionFromPolicy(int, int) error
	GetUsersByPolicy(int, int, int) (*items, error)
	GetGroupsByPolicy(int, int, int) (*items, error)
	GetPermissionsByPolicy(int, int, int) (*items, error)
}

type policyService struct {
	storage storage.DB
}

func NewPolicyService(
	storage storage.DB,
) PolicyService {
	return &policyService{
		storage,
	}
}

func (service *policyService) CreatePolicy(
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

func (service *policyService) GetPolicy(
	policyID int,
) (*models.Policy, error) {
	policy, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", policyID)
	}

	return policy, nil
}

func (service *policyService) GetPolicies(
	page int,
	perPage int,
) (*items, error) {
	policies, err := service.storage.GetPolicies(page, perPage)
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
	}

	return response, nil
}

func (service *policyService) UpdatePolicy(
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

func (service *policyService) RemovePolicy(
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

func (service *policyService) AddPermissionsToPolicy(
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

func (service *policyService) RemovePermissionFromPolicy(
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

func (service *policyService) GetPermissionsByPolicy(
	policyID int,
	page int,
	perPage int,
) (*items, error) {
	policy, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", policyID)
	}

	permissions, count, err := service.storage.GetPermissionsByPolicy(
		policy,
		&page,
		&perPage,
	)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	items := &items{
		permissions,
		count,
	}

	return items, nil
}

func (service *policyService) GetUsersByPolicy(
	policyID int,
	page int,
	perPage int,
) (*items, error) {
	policy, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", policyID)
	}

	users, count, err := service.storage.GetUsersByPolicy(
		policy,
		&page,
		&perPage,
	)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	items := &items{
		users,
		count,
	}

	return items, nil
}

func (service *policyService) GetGroupsByPolicy(
	policyID int,
	page int,
	perPage int,
) (*items, error) {
	policy, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", policyID)
	}

	groups, count, err := service.storage.GetGroupsByPolicy(
		policy,
		&page,
		&perPage,
	)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	items := &items{
		groups,
		count,
	}

	return items, nil
}
