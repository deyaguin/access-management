package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type ActionService interface {
	CreateAction(*models.Action) (*models.Action, error)
	GetAction(int) (*models.Action, error)
	GetActions() (*[]models.Action, error)
	UpdateAction(*models.Action) (*models.Action, error)
	RemoveAction(*models.Action) error
}

type actionService struct {
	storage storage.DB
}

func NewActionService(storage storage.DB) ActionService {
	return &actionService{
		storage,
	}
}

func (service *actionService) CreateAction(actionCreating *models.Action) (*models.Action, error) {
	action := new(models.Action)

	if err := validator.Validate(actionCreating); err != nil {
		return action, NewValidationError(err.Error())
	}

	action.SetFields(actionCreating)

	err := service.storage.CreateAction(action)
	return action, err
}

func (service *actionService) GetAction(id int) (*models.Action, error) {
	action, err := service.storage.GetAction(id)
	if err != nil {
		return action, NewEntityNotFoundError("action", id)
	}

	return action, err
}

func (service *actionService) GetActions() (*[]models.Action, error) {
	return service.storage.GetActions()
}

func (service *actionService) UpdateAction(actionUpdating *models.Action) (*models.Action, error) {
	action, err := service.storage.GetAction(actionUpdating.ID)
	if err != nil {
		return action, err
	}

	if err := validator.Validate(actionUpdating); err != nil {
		return action, NewValidationError(err.Error())
	}

	action.SetFields(actionUpdating)

	return action, err
}

func (service *actionService) RemoveAction(action *models.Action) error {
	if _, err := service.storage.GetAction(action.ID); err != nil {
		return NewEntityNotFoundError("action", action.ID)
	}

	return service.storage.RemoveAction(action)
}
