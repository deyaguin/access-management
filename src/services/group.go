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
	group := new(models.Group)

	if err := validator.Validate(groupCreating); err != nil {
		return group, NewValidationError(err.Error())
	}

	group.SetFields(groupCreating)

	err := service.storage.CreateGroup(group)
	return group, err
}

func (service *groupService) GetGroup(id int) (*models.Group, error) {
	group, err := service.storage.GetGroup(id)
	if err != nil {
		return group, NewEntityNotFoundError("group", id)
	}

	return group, err
}

func (service *groupService) GetGroups() (*[]models.Group, error) {
	return service.storage.GetGroups()
}

func (service *groupService) UpdateGroup(groupUpdating *models.Group) (*models.Group, error) {
	group, err := service.storage.GetGroup(groupUpdating.ID)
	if err != nil {
		return group, NewEntityNotFoundError("group", groupUpdating.ID)
	}

	if err := validator.Validate(groupUpdating); err != nil {
		return group, NewValidationError(err.Error())
	}

	group.SetFields(groupUpdating)

	return group, service.storage.UpdateGroup(group)
}

func (service *groupService) RemoveGroup(group *models.Group) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewEntityNotFoundError("group", group.ID)
	}

	return service.storage.RemoveGroup(group)
}
