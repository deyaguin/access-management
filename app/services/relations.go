package services

import (
	"gitlab/nefco/accessControl/app/models"
	"gitlab/nefco/accessControl/app/storage"
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
		return NewGroupNotFoundError(group.ID)
	}

	for _, u := range *users {
		if _, err := service.storage.GetUser(u.ID); err != nil {
			return NewUserNotFoundError(u.ID)
		}
		if u.ID == 0 {
			return *new(error)
		}
	}

	return service.storage.AddUsersToGroup(group, users)
}

func (service *relationsService) AddPermissionsToPolicy(policy *models.Policy, permissions *[]models.Permission) error {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewPolicyNotFoundError(policy.ID)
	}

	for _, p := range *permissions {
		if err := validator.Validate(p); err != nil {
			return err
		}
	}

	return service.storage.AddPermissionsToPolicy(policy, permissions)
}

func (service *relationsService) AttachPoliciesByUser(user *models.User, policies *[]models.Policy) error {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewUserNotFoundError(user.ID)
	}

	for _, p := range *policies {
		if _, err := service.storage.GetPolicy(p.ID); err != nil {
			return NewPolicyNotFoundError(p.ID)
		}
		if p.ID == 0 {
			return *new(error)
		}
	}

	return service.storage.AttachPoliciesByUser(user, policies)
}

func (service *relationsService) AttachPoliciesByGroup(group *models.Group, policies *[]models.Policy) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewGroupNotFoundError(group.ID)
	}

	for _, p := range *policies {
		if _, err := service.storage.GetPolicy(p.ID); err != nil {
			return NewPolicyNotFoundError(p.ID)
		}
		if p.ID == 0 {
			return *new(error)
		}
	}

	return service.storage.AttachPoliciesByGroup(group, policies)
}

func (service *relationsService) RemoveUserFromGroup(group *models.Group, user *models.User) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewGroupNotFoundError(group.ID)
	}
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewUserNotFoundError(user.ID)
	}

	return service.storage.RemoveUserFromGroup(group, user)
}

func (service *relationsService) RemovePermissionFromPolicy(policy *models.Policy, permission *models.Permission) error {
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewPolicyNotFoundError(policy.ID)
	}
	if _, err := service.storage.GetPermission(permission.ID); err != nil {
		return NewPermissionNotFoundError(permission.ID)
	}

	return service.storage.RemovePermission(permission)
}

func (service *relationsService) DetachPolicyByUser(user *models.User, policy *models.Policy) error {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewUserNotFoundError(user.ID)
	}
	if _, err := service.storage.GetPolicy(user.ID); err != nil {
		return NewPolicyNotFoundError(user.ID)
	}

	return service.storage.DetachPolicyByUser(user, policy)
}

func (service *relationsService) DetachPolicyByGroup(group *models.Group, policy *models.Policy) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewGroupNotFoundError(group.ID)
	}
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewPolicyNotFoundError(policy.ID)
	}

	return service.storage.DetachPolicyByGroup(group, policy)
}

func (service *relationsService) GetGroupsByUser(user *models.User) (*[]models.Group, error) {
	groups := new([]models.Group)

	if _, err := service.storage.GetUser(user.ID); err != nil {
		return groups, NewUserNotFoundError(user.ID)
	}

	return service.storage.GetGroupsByUser(user)
}

func (service *relationsService) GetUsersByGroup(group *models.Group) (*[]models.User, error) {
	users := new([]models.User)

	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return users, NewGroupNotFoundError(group.ID)
	}

	return service.storage.GetUsersByGroup(group)
}

func (service *relationsService) GetPoliciesByUser(user *models.User) (*[]models.Policy, error) {
	policies := new([]models.Policy)

	if _, err := service.storage.GetUser(user.ID); err != nil {
		return policies, NewUserNotFoundError(user.ID)
	}

	return service.storage.GetPoliciesByUser(user)
}

func (service *relationsService) GetPoliciesByGroup(group *models.Group) (*[]models.Policy, error) {
	policies := new([]models.Policy)

	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return policies, NewGroupNotFoundError(group.ID)
	}

	return service.storage.GetPoliciesByGroup(group)
}

func (service *relationsService) GetUsersByPolicy(policy *models.Policy) (*[]models.User, error) {
	users := new([]models.User)

	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return users, NewPolicyNotFoundError(policy.ID)
	}

	return service.storage.GetUsersByPolicy(policy)
}

func (service *relationsService) GetGroupsByPolicy(policy *models.Policy) (*[]models.Group, error) {
	groups := new([]models.Group)

	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return groups, NewPolicyNotFoundError(policy.ID)
	}

	return service.storage.GetGroupsByPolicy(policy)
}

func (service *relationsService) GetPermissionsByPolicy(policy *models.Policy) (*[]models.Permission, error) {
	permissions := new([]models.Permission)

	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return permissions, NewPolicyNotFoundError(policy.ID)
	}

	return service.storage.GetPermissionsByPolicy(policy)
}
