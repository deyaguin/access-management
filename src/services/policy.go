package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type PolicyService interface {
	CreatePolicy(*models.Policy) (*models.Policy, error)
	GetPolicy(int) (*models.Policy, error)
	GetPolicies() (*[]models.Policy, error)
	UpdatePolicy(*models.Policy) (*models.Policy, error)
	RemovePolicy(*models.Policy) error

	AddPermissionsToPolicy(*models.Policy, *[]models.Permission) error
	RemovePermissionFromPolicy(*models.Policy, *models.Permission) error
	GetUsersByPolicy(*models.Policy) (*[]models.User, error)
	GetGroupsByPolicy(*models.Policy) (*[]models.Group, error)
	GetPermissionsByPolicy(*models.Policy) (*[]models.Permission, error)
}

type policyService struct {
	storage storage.DB
}

func NewPolicyService(storage storage.DB) PolicyService {
	return &policyService{
		storage,
	}
}

func (service *policyService) CreatePolicy(policyCreating *models.Policy) (*models.Policy, error) {
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

func (service *policyService) GetPolicy(id int) (*models.Policy, error) {
	policy, err := service.storage.GetPolicy(id)
	if err != nil {
		return nil, NewEntityNotFoundError("policy", id)
	}

	return policy, nil
}

func (service *policyService) GetPolicies() (*[]models.Policy, error) {
	policies, err := service.storage.GetPolicies()
	if err != nil {
		return nil, err
	}

	return policies, nil
}

func (service *policyService) UpdatePolicy(policyUpdating *models.Policy) (*models.Policy, error) {
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

func (service *policyService) RemovePolicy(policy *models.Policy) error {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewEntityNotFoundError("policy", policy.ID)
	}

	if err := service.storage.RemovePolicy(policy); err != nil {
		return err
	}

	return nil
}

func (service *policyService) AddPermissionsToPolicy(policy *models.Policy, permissions *[]models.Permission) error {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewEntityNotFoundError("policy", policy.ID)
	}

	for _, permission := range *permissions {
		if err := validator.Validate(permission); err != nil {
			return err
		}
	}

	if err := service.storage.AddPermissionsToPolicy(policy, permissions); err != nil {
		return err
	}

	return nil
}

func (service *policyService) RemovePermissionFromPolicy(policy *models.Policy, permission *models.Permission) error {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewEntityNotFoundError("policy", policy.ID)
	}
	if _, err := service.storage.GetPermission(permission.ID); err != nil {
		return NewEntityNotFoundError("permission", permission.ID)
	}

	if err := service.storage.RemovePermission(permission); err != nil {
		return err
	}

	return nil
}

func (service *policyService) GetPermissionsByPolicy(policy *models.Policy) (*[]models.Permission, error) {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return nil, NewEntityNotFoundError("policy", policy.ID)
	}

	permissions, err := service.storage.GetPermissionsByPolicy(policy)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (service *policyService) GetUsersByPolicy(policy *models.Policy) (*[]models.User, error) {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return nil, NewEntityNotFoundError("policy", policy.ID)
	}

	users, err := service.storage.GetUsersByPolicy(policy)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *policyService) GetGroupsByPolicy(policy *models.Policy) (*[]models.Group, error) {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return nil, NewEntityNotFoundError("policy", policy.ID)
	}

	groups, err := service.storage.GetGroupsByPolicy(policy)
	if err != nil {
		return nil, err
	}

	return groups, nil
}
