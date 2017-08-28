package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type UserService interface {
	CreateUser(*models.User) (*models.User, error)
	GetUser(int) (*models.User, error)
	GetUsers() (*[]models.User, error)
	UpdateUser(*models.User) (*models.User, error)
	RemoveUser(*models.User) error
}

type userService struct {
	storage storage.DB
}

func NewUserService(storage storage.DB) UserService {
	return &userService{
		storage,
	}
}

func (service *userService) CreateUser(userCreating *models.User) (*models.User, error) {
	user := new(models.User)

	if err := validator.Validate(userCreating); err != nil {
		return user, NewValidationError(err.Error())
	}

	user.SetFields(userCreating)

	err := service.storage.CreateUser(user)
	return user, err
}

func (service *userService) GetUser(id int) (*models.User, error) {
	user, err := service.storage.GetUser(id)
	if err != nil {
		return user, NewEntityNotFoundError("user", id)
	}

	return user, err
}

func (service *userService) GetUsers() (*[]models.User, error) {
	return service.storage.GetUsers()
}

func (service *userService) UpdateUser(userUpdating *models.User) (*models.User, error) {
	user, err := service.storage.GetUser(userUpdating.ID)
	if err != nil {
		return user, NewEntityNotFoundError("user", userUpdating.ID)
	}

	if err := validator.Validate(userUpdating); err != nil {
		return user, NewValidationError(err.Error())
	}

	user.SetFields(userUpdating)

	return user, service.storage.UpdateUser(user)
}

func (service *userService) RemoveUser(user *models.User) error {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return err
	}

	return service.storage.RemoveUser(user)
}
