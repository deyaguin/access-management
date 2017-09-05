package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type PermissionsService interface {
	UpdatePermission(*models.Permission) (*models.Permission, error)
	RemovePermission(int) error
}

type permissionsService struct {
	storage storage.DB
}

func NewPermissionsService(
	storage storage.DB,
) PermissionsService {
	return &permissionsService{
		storage,
	}
}

func (service *permissionsService) UpdatePermission(
	permissionUpdating *models.Permission,
) (*models.Permission, error) {
	permission, err := service.storage.GetPermission(permissionUpdating.ID)
	if err != nil {
		return nil, NewEntityNotFoundError(
			"permission",
			permissionUpdating.ID,
		)
	}

	if err := validator.Validate(permissionUpdating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	permission.SetFields(permissionUpdating)
	if err := service.storage.UpdatePermission(permission); err != nil {
		return nil, NewEntityUpdateError(err.Error())
	}

	return permission, nil
}

func (service *permissionsService) RemovePermission(
	permissionID int,
) error {
	permission, err := service.storage.GetPermission(permissionID)
	if err != nil {
		return NewEntityNotFoundError(
			"permission",
			permissionID,
		)
	}

	if err := service.storage.RemovePermission(permission); err != nil {
		return NewEntityRemoveError(err.Error())
	}

	return nil
}
