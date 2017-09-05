package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type ActionsService interface {
	CreateAction(*models.Action) (*models.Action, error)
	GetAction(int) (*models.Action, error)
	GetActions(int, int) (*[]models.Action, error)
	UpdateAction(*models.Action) (*models.Action, error)
	RemoveAction(int) error
}

type actionsService struct {
	storage storage.DB
}

func NewActionsService(
	storage storage.DB,
) ActionsService {
	return &actionsService{
		storage,
	}
}

func (service *actionsService) CreateAction(
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

func (service *actionsService) GetAction(
	actionID int,
) (*models.Action, error) {
	action, err := service.storage.GetAction(actionID)
	if err != nil {
		return nil, NewEntityNotFoundError("action", actionID)
	}

	return action, nil
}

func (service *actionsService) GetActions(
	page int,
	perPage int,
) (*[]models.Action, error) {
	actions, err := service.storage.GetActions(page, perPage)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	return actions, nil
}

func (service *actionsService) UpdateAction(
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

func (service *actionsService) RemoveAction(
	actionID int,
) error {
	action, err := service.storage.GetAction(actionID)
	if err != nil {
		return NewEntityNotFoundError("action", actionID)
	}

	if err := service.storage.RemoveAction(action); err != nil {
		return NewEntityRemoveError(err.Error())
	}

	return nil
}
