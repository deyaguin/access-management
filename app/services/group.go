package services

import (
	"gitlab/nefco/accessControl/app/models"
	"gitlab/nefco/accessControl/app/storage"
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
		return groupCreating, err
	}

	group.SetFields(groupCreating)

	err := service.storage.CreateGroup(group)
	return group, err
}

func (service *groupService) GetGroup(id int) (*models.Group, error) {
	group, err := service.storage.GetGroup(id)
	if err != nil {
		return group, NewGroupNotFoundError(id)
	}

	return group, err
}

func (service *groupService) GetGroups() (*[]models.Group, error) {
	return service.storage.GetGroups()
}

func (service *groupService) UpdateGroup(groupUpdating *models.Group) (*models.Group, error) {
	group, err := service.storage.GetGroup(groupUpdating.ID)
	if err != nil {
		return group, err
	}

	if err := validator.Validate(groupUpdating); err != nil {
		return group, err
	}

	group.SetFields(groupUpdating)

	return group, service.storage.UpdateGroup(group)
}

func (service *groupService) RemoveGroup(group *models.Group) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewGroupNotFoundError(group.ID)
	}

	return service.storage.RemoveGroup(group)
}
