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
		return nil, err
	}
	return policy, nil
}

func (service *policyService) GetPolicy(
	id int,
) (*models.Policy, error) {
	policy, err := service.storage.GetPolicy(id)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", id)
	}

	return policy, nil
}

func (service *policyService) GetPolicies(
	page int,
	perPage int,
) (*items, error) {
	policies, err := service.storage.GetPolicies(page, perPage)
	if err != nil {
		return nil, err
	}

	count, err := service.storage.GetPoliciesCount()
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return policy, nil
}

func (service *policyService) RemovePolicy(
	id int,
) error {
	policy, err := service.storage.GetPolicy(id)
	if err != nil {
		return NewEntityNotFoundError("policy", id)
	}

	if err := service.storage.RemovePolicy(policy); err != nil {
		return err
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
			return err
		}
	}

	if err := service.storage.AddPermissionsToPolicy(
		policy,
		permissions,
	); err != nil {
		return err
	}

	return nil
}

func (service *policyService) RemovePermissionFromPolicy(
	policyId int,
	permissionId int,
) error {
	_, err := service.storage.GetPolicy(policyId)
	if err != nil {
		return NewEntityNotFoundError("policy", policyId)
	}

	permission, err := service.storage.GetPermission(permissionId)
	if err != nil {
		return NewEntityNotFoundError("permission", permissionId)
	}

	if err := service.storage.RemovePermission(permission); err != nil {
		return err
	}

	return nil
}

func (service *policyService) GetPermissionsByPolicy(
	policyId int,
	page int,
	perPage int,
) (*items, error) {
	policy, err := service.storage.GetPolicy(policyId)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", policyId)
	}

	permissions, count, err := service.storage.GetPermissionsByPolicy(
		policy,
		&page,
		&perPage,
	)
	if err != nil {
		return nil, err
	}

	items := &items{
		permissions,
		count,
	}

	return items, nil
}

func (service *policyService) GetUsersByPolicy(
	policyId int,
	page int,
	perPage int,
) (*items, error) {
	policy, err := service.storage.GetPolicy(policyId)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", policyId)
	}

	users, count, err := service.storage.GetUsersByPolicy(
		policy,
		&page,
		&perPage,
	)
	if err != nil {
		return nil, err
	}

	items := &items{
		users,
		count,
	}

	return items, nil
}

func (service *policyService) GetGroupsByPolicy(
	policyId int,
	page int,
	perPage int,
) (*items, error) {
	policy, err := service.storage.GetPolicy(policyId)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", policyId)
	}

	groups, count, err := service.storage.GetGroupsByPolicy(
		policy,
		&page,
		&perPage,
	)
	if err != nil {
		return nil, err
	}

	items := &items{
		groups,
		count,
	}

	return items, nil
}
