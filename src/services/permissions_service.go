package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type PermissionService interface {
	UpdatePermission(*models.Permission) (*models.Permission, error)
	RemovePermission(int) error
}

type permissionService struct {
	storage storage.DB
}

func NewPermissionService(
	storage storage.DB,
) PermissionService {
	return &permissionService{
		storage,
	}
}

func (service *permissionService) UpdatePermission(
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
		return nil, err
	}

	return permission, nil
}

func (service *permissionService) RemovePermission(
	permissionId int,
) error {
	permission, err := service.storage.GetPermission(permissionId)
	if err != nil {
		return NewEntityNotFoundError(
			"permission",
			permissionId,
		)
	}

	if err := service.storage.RemovePermission(permission); err != nil {
		return err
	}

	return nil
}
