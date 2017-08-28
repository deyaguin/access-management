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
		if user.ID == 0 {
			return *new(error)
		}
	}

	return service.storage.AddUsersToGroup(group, users)
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

	return service.storage.AddPermissionsToPolicy(policy, permissions)
}

func (service *relationsService) AttachPoliciesByUser(user *models.User, policies *[]models.Policy) error {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewEntityNotFoundError("user", user.ID)
	}

	for _, policy := range *policies {
		if _, err := service.storage.GetPolicy(policy.ID); err != nil {
			return NewEntityNotFoundError("policy", policy.ID)
		}
		if policy.ID == 0 {
			return *new(error)
		}
	}

	return service.storage.AttachPoliciesByUser(user, policies)
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

	return service.storage.AttachPoliciesByGroup(group, policies)
}

func (service *relationsService) RemoveUserFromGroup(group *models.Group, user *models.User) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewEntityNotFoundError("group", group.ID)
	}
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewEntityNotFoundError("user", user.ID)
	}

	return service.storage.RemoveUserFromGroup(group, user)
}

func (service *relationsService) RemovePermissionFromPolicy(policy *models.Policy, permission *models.Permission) error {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewEntityNotFoundError("policy", policy.ID)
	}
	if _, err := service.storage.GetPermission(permission.ID); err != nil {
		return NewEntityNotFoundError("permission", permission.ID)
	}

	return service.storage.RemovePermission(permission)
}

func (service *relationsService) DetachPolicyByUser(user *models.User, policy *models.Policy) error {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewEntityNotFoundError("user", user.ID)
	}
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewEntityNotFoundError("policy", policy.ID)
	}

	return service.storage.DetachPolicyByUser(user, policy)
}

func (service *relationsService) DetachPolicyByGroup(group *models.Group, policy *models.Policy) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewEntityNotFoundError("group", group.ID)
	}
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewEntityNotFoundError("policy", policy.ID)
	}

	return service.storage.DetachPolicyByGroup(group, policy)
}

func (service *relationsService) GetGroupsByUser(user *models.User) (*[]models.Group, error) {
	groups := new([]models.Group)

	if _, err := service.storage.GetUser(user.ID); err != nil {
		return groups, NewEntityNotFoundError("user", user.ID)
	}

	return service.storage.GetGroupsByUser(user)
}

func (service *relationsService) GetUsersByGroup(group *models.Group) (*[]models.User, error) {
	users := new([]models.User)

	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return users, NewEntityNotFoundError("group", group.ID)
	}

	return service.storage.GetUsersByGroup(group)
}

func (service *relationsService) GetPoliciesByUser(user *models.User) (*[]models.Policy, error) {
	policies := new([]models.Policy)

	if _, err := service.storage.GetUser(user.ID); err != nil {
		return policies, NewEntityNotFoundError("user", user.ID)
	}

	return service.storage.GetPoliciesByUser(user)
}

func (service *relationsService) GetPoliciesByGroup(group *models.Group) (*[]models.Policy, error) {
	policies := new([]models.Policy)

	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return policies, NewEntityNotFoundError("group", group.ID)
	}

	return service.storage.GetPoliciesByGroup(group)
}

func (service *relationsService) GetUsersByPolicy(policy *models.Policy) (*[]models.User, error) {
	users := new([]models.User)

	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return users, NewEntityNotFoundError("policy", policy.ID)
	}

	return service.storage.GetUsersByPolicy(policy)
}

func (service *relationsService) GetGroupsByPolicy(policy *models.Policy) (*[]models.Group, error) {
	groups := new([]models.Group)

	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return groups, NewEntityNotFoundError("policy", policy.ID)
	}

	return service.storage.GetGroupsByPolicy(policy)
}

func (service *relationsService) GetPermissionsByPolicy(policy *models.Policy) (*[]models.Permission, error) {
	permissions := new([]models.Permission)

	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return permissions, NewEntityNotFoundError("policy", policy.ID)
	}

	return service.storage.GetPermissionsByPolicy(policy)
}
