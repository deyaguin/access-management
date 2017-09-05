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

func NewActionService(
	storage storage.DB,
) ActionService {
	return &actionService{
		storage,
	}
}

func (service *actionService) CreateAction(
	actionCreating *models.Action,
) (*models.Action, error) {
	if err := validator.Validate(actionCreating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	action := new(models.Action)
	action.SetFields(actionCreating)

	if err := service.storage.CreateAction(action); err != nil {
		return nil, NewEntityCreateError(err.Error())
	}

	return action, nil
}

func (service *actionService) GetAction(
	actionID int,
) (*models.Action, error) {
	action, err := service.storage.GetAction(actionID)
	if err != nil {
		return nil, NewEntityNotFoundError("action", actionID)
	}

	return action, nil
}

func (service *actionService) GetActions() (*[]models.Action, error) {
	action, err := service.storage.GetActions()
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	return action, nil
}

func (service *actionService) UpdateAction(
	actionUpdating *models.Action,
) (*models.Action, error) {
	action, err := service.storage.GetAction(actionUpdating.ID)
	if err != nil {
		return nil, NewEntityNotFoundError(
			"action",
			actionUpdating.ID,
		)
	}

	if err := validator.Validate(actionUpdating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	action.SetFields(actionUpdating)
	if err := service.storage.UpdateAction(action); err != nil {
		return nil, NewEntityUpdateError(err.Error())
	}

	return action, nil
}

func (service *actionService) RemoveAction(
	action *models.Action,
) error {
	if _, err := service.storage.GetAction(action.ID); err != nil {
		return NewEntityNotFoundError(
			"action",
			action.ID,
		)
	}

	if err := service.storage.RemoveAction(action); err != nil {
		return NewEntityRemoveError(err.Error())
	}

	return nil
}
