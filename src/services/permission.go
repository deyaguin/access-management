package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type PermissionService interface {
	UpdatePermission(*models.Permission) (*models.Permission, error)
	RemovePermission(*models.Permission) error
}

type permissionService struct {
	storage storage.DB
}

func NewPermissionService(storage storage.DB) PermissionService {
	return &permissionService{
		storage,
	}
}

func (service *permissionService) UpdatePermission(permissionUpdating *models.Permission) (*models.Permission, error) {
	permission, err := service.storage.GetPermission(permissionUpdating.ID)
	if err != nil {
		return permission, NewEntityNotFoundError("permission", permissionUpdating.ID)
	}

	if err := validator.Validate(permissionUpdating); err != nil {
		return permission, NewValidationError(err.Error())
	}

	permission.SetFields(permissionUpdating)

	return permission, service.storage.UpdatePermission(permission)
}

func (service *permissionService) RemovePermission(permission *models.Permission) error {
	if _, err := service.storage.GetPermission(permission.ID); err != nil {
		return NewEntityNotFoundError("permission", permission.ID)
	}

	return service.storage.RemovePermission(permission)
}
