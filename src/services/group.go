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
