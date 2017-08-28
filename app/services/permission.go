package services

import (
	"gitlab/nefco/accessControl/app/models"
	"gitlab/nefco/accessControl/app/storage"
)

type PermissionService interface {
	UpdatePermission(*models.Permission) error
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

func (service *permissionService) UpdatePermission(permission *models.Permission) error {
	if _, err := service.storage.GetPermission(permission.ID); err != nil {
		return NewPermissionNotFoundError(permission.ID)
	}

	return service.storage.UpdatePermission(permission)
}

func (service *permissionService) RemovePermission(permission *models.Permission) error {
	if _, err := service.storage.GetPermission(permission.ID); err != nil {
		return NewPermissionNotFoundError(permission.ID)
	}

	return service.storage.RemovePermission(permission)
}
