package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type RelationsService interface {
	AddUsersToGroup(*models.Group, *[]models.User) error
	AddPermissionsToPolicy(*models.Policy, *[]models.Permission) error
	AttachPoliciesByUser(*models.User, *[]models.Policy) error
	AttachPoliciesByGroup(*models.Group, *[]models.Policy) error

	RemoveUserFromGroup(*models.Group, *models.User) error
	RemovePermissionFromPolicy(*models.Policy, *models.Permission) error
	DetachPolicyByUser(*models.User, *models.Policy) error
	DetachPolicyByGroup(*models.Group, *models.Policy) error

	GetUsersByGroup(*models.Group) (*[]models.User, error)
	GetGroupsByUser(*models.User) (*[]models.Group, error)
	GetPoliciesByUser(*models.User) (*[]models.Policy, error)
	GetPoliciesByGroup(*models.Group) (*[]models.Policy, error)
	GetUsersByPolicy(*models.Policy) (*[]models.User, error)
	GetGroupsByPolicy(*models.Policy) (*[]models.Group, error)
	GetPermissionsByPolicy(*models.Policy) (*[]models.Permission, error)
}

type relationsService struct {
	storage storage.DB
}

func NewRelationsService(storage storage.DB) RelationsService {
	return &relationsService{
		storage,
	}
}

func (service *relationsService) AddUsersToGroup(group *models.Group, users *[]models.User) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewEntityNotFoundError("group", group.ID)
	}

	for _, user := range *users {
		if _, err := service.storage.GetUser(user.ID); err != nil {
			return NewEntityNotFoundError("user", user.ID)
		}
	}

	if err := service.storage.AddUsersToGroup(group, users); err != nil {
		return err
	}

	return nil
}

func (service *relationsService) AddPermissionsToPolicy(policy *models.Policy, permissions *[]models.Permission) error {
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

func (service *relationsService) AttachPoliciesByUser(user *models.User, policies *[]models.Policy) error {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewEntityNotFoundError("user", user.ID)
	}

	for _, policy := range *policies {
		if _, err := service.storage.GetPolicy(policy.ID); err != nil {
			return NewEntityNotFoundError("policy", policy.ID)
		}
	}

	if err := service.storage.AttachPoliciesByUser(user, policies); err != nil {
		return err
	}

	return nil
}

func (service *relationsService) AttachPoliciesByGroup(group *models.Group, policies *[]models.Policy) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewEntityNotFoundError("group", group.ID)
	}

	for _, policy := range *policies {
		if _, err := service.storage.GetPolicy(policy.ID); err != nil {
			return NewEntityNotFoundError("policy", policy.ID)
		}
		if policy.ID == 0 {
			return *new(error)
		}
	}

	if err := service.storage.AttachPoliciesByGroup(group, policies); err != nil {
		return err
	}

	return nil
}

func (service *relationsService) RemoveUserFromGroup(group *models.Group, user *models.User) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewEntityNotFoundError("group", group.ID)
	}
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewEntityNotFoundError("user", user.ID)
	}

	if err := service.storage.RemoveUserFromGroup(group, user); err != nil {
		return err
	}

	return nil
}

func (service *relationsService) RemovePermissionFromPolicy(policy *models.Policy, permission *models.Permission) error {
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

func (service *relationsService) DetachPolicyByUser(user *models.User, policy *models.Policy) error {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewEntityNotFoundError("user", user.ID)
	}
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewEntityNotFoundError("policy", policy.ID)
	}

	if err := service.storage.DetachPolicyByUser(user, policy); err != nil {
		return err
	}

	return nil
}

func (service *relationsService) DetachPolicyByGroup(group *models.Group, policy *models.Policy) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewEntityNotFoundError("group", group.ID)
	}
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewEntityNotFoundError("policy", policy.ID)
	}

	if err := service.storage.DetachPolicyByGroup(group, policy); err != nil {
		return err
	}

	return nil
}

func (service *relationsService) GetGroupsByUser(user *models.User) (*[]models.Group, error) {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return nil, NewEntityNotFoundError("user", user.ID)
	}

	groups, err := service.storage.GetGroupsByUser(user)
	if err != nil {
		return nil, err
	}

	return groups, err
}

func (service *relationsService) GetUsersByGroup(group *models.Group) (*[]models.User, error) {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return nil, NewEntityNotFoundError("group", group.ID)
	}

	users, err := service.storage.GetUsersByGroup(group)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *relationsService) GetPoliciesByUser(user *models.User) (*[]models.Policy, error) {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return nil, NewEntityNotFoundError("user", user.ID)
	}

	policies, err := service.storage.GetPoliciesByUser(user)
	if err != nil {
		return nil, err
	}

	return policies, nil
}

func (service *relationsService) GetPoliciesByGroup(group *models.Group) (*[]models.Policy, error) {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return nil, NewEntityNotFoundError("group", group.ID)
	}

	policies, err := service.storage.GetPoliciesByGroup(group)
	if err != nil {
		return nil, err
	}

	return policies, nil
}

func (service *relationsService) GetUsersByPolicy(policy *models.Policy) (*[]models.User, error) {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return nil, NewEntityNotFoundError("policy", policy.ID)
	}

	users, err := service.storage.GetUsersByPolicy(policy)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *relationsService) GetGroupsByPolicy(policy *models.Policy) (*[]models.Group, error) {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return nil, NewEntityNotFoundError("policy", policy.ID)
	}

	groups, err := service.storage.GetGroupsByPolicy(policy)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (service *relationsService) GetPermissionsByPolicy(policy *models.Policy) (*[]models.Permission, error) {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return nil, NewEntityNotFoundError("policy", policy.ID)
	}

	permissions, err := service.storage.GetPermissionsByPolicy(policy)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
