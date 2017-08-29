package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type GroupService interface {
	CreateGroup(*models.Group) (*models.Group, error)
	GetGroup(int) (*models.Group, error)
	GetGroups() (*[]models.Group, error)
	UpdateGroup(*models.Group) (*models.Group, error)
	RemoveGroup(*models.Group) error

	AddUsersToGroup(*models.Group, *[]models.User) error
	RemoveUserFromGroup(*models.Group, *models.User) error
	GetUsersByGroup(*models.Group) (*[]models.User, error)
	AttachPoliciesByGroup(*models.Group, *[]models.Policy) error
	DetachPolicyByGroup(*models.Group, *models.Policy) error
	GetPoliciesByGroup(*models.Group) (*[]models.Policy, error)
}

type groupService struct {
	storage storage.DB
}

func NewGroupService(storage storage.DB) GroupService {
	return &groupService{
		storage,
	}
}

func (service *groupService) CreateGroup(groupCreating *models.Group) (*models.Group, error) {
	if err := validator.Validate(groupCreating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	group := new(models.Group)
	group.SetFields(groupCreating)

	if err := service.storage.CreateGroup(group); err != nil {
		return nil, err
	}

	return group, nil
}

func (service *groupService) GetGroup(id int) (*models.Group, error) {
	group, err := service.storage.GetGroup(id)
	if err != nil {
		return nil, NewEntityNotFoundError("group", id)
	}

	return group, nil
}

func (service *groupService) GetGroups() (*[]models.Group, error) {
	groups, err := service.storage.GetGroups()
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (service *groupService) UpdateGroup(groupUpdating *models.Group) (*models.Group, error) {
	group, err := service.storage.GetGroup(groupUpdating.ID)
	if err != nil {
		return nil, NewEntityNotFoundError("group", groupUpdating.ID)
	}

	if err := validator.Validate(groupUpdating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	group.SetFields(groupUpdating)
	if err := service.storage.UpdateGroup(group); err != nil {
		return nil, err
	}

	return group, nil
}

func (service *groupService) RemoveGroup(group *models.Group) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewEntityNotFoundError("group", group.ID)
	}

	if err := service.storage.RemoveGroup(group); err != nil {
		return err
	}

	return nil
}

func (service *groupService) AddUsersToGroup(group *models.Group, users *[]models.User) error {
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

func (service *groupService) RemoveUserFromGroup(group *models.Group, user *models.User) error {
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

func (service *groupService) GetUsersByGroup(group *models.Group) (*[]models.User, error) {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return nil, NewEntityNotFoundError("group", group.ID)
	}

	users, err := service.storage.GetUsersByGroup(group)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *groupService) AttachPoliciesByGroup(group *models.Group, policies *[]models.Policy) error {
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

func (service *groupService) DetachPolicyByGroup(group *models.Group, policy *models.Policy) error {
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

func (service *groupService) GetPoliciesByGroup(group *models.Group) (*[]models.Policy, error) {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return nil, NewEntityNotFoundError("group", group.ID)
	}

	policies, err := service.storage.GetPoliciesByGroup(group)
	if err != nil {
		return nil, err
	}

	return policies, nil
}
