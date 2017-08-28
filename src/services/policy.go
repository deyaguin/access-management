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
	policy := new(models.Policy)

	if err := validator.Validate(policy); err != nil {
		return policy, err
	}

	policy.SetFields(policyCreating)

	err := service.storage.CreatePolicy(policy)
	return policy, err
}

func (service *policyService) GetPolicy(id int) (*models.Policy, error) {
	policy, err := service.storage.GetPolicy(id)
	if err != nil {
		return policy, NewEntityNotFoundError("policy", id)
	}

	return policy, err
}

func (service *policyService) GetPolicies() (*[]models.Policy, error) {
	return service.storage.GetPolicies()
}

func (service *policyService) UpdatePolicy(policyUpdating *models.Policy) (*models.Policy, error) {
	policy, err := service.storage.GetPolicy(policyUpdating.ID)
	if err != nil {
		return policy, NewEntityNotFoundError("permission", policyUpdating.ID)
	}

	if err := validator.Validate(policyUpdating); err != nil {
		return policy, NewValidationError(err.Error())
	}

	policy.SetFields(policyUpdating)

	return policy, service.storage.UpdatePolicy(policy)
}

func (service *policyService) RemovePolicy(policy *models.Policy) error {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewEntityNotFoundError("policy", policy.ID)
	}

	return service.storage.RemovePolicy(policy)
}
