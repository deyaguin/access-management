package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type usersResponse struct {
	Users *[]models.User `json:"users"`
	Count int            `json:"count"`
}

type UserService interface {
	CreateUser(*models.User) (*models.User, error)
	GetUser(int) (*models.User, error)
	GetUsers(int) (*usersResponse, error)
	UpdateUser(*models.User) (*models.User, error)
	RemoveUser(*models.User) error

	AttachPoliciesByUser(*models.User, *[]models.Policy) error
	DetachPolicyByUser(*models.User, *models.Policy) error
	GetPoliciesByUser(*models.User) (*[]models.Policy, error)
	GetGroupsByUser(*models.User) (*[]models.Group, error)
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
	if err := validator.Validate(userCreating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	user := new(models.User)
	user.SetFields(userCreating)

	if err := service.storage.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (service *userService) GetUser(id int) (*models.User, error) {
	user, err := service.storage.GetUser(id)
	if err != nil {
		return nil, NewEntityNotFoundError("user", id)
	}

	return user, nil
}

func (service *userService) GetUsers(page int) (*usersResponse, error) {
	users, err := service.storage.GetUsers(page)
	if err != nil {
		return nil, err
	}

	count, err := service.storage.GetUsersCount()
	if err != nil {
		return nil, err
	}

	response := &usersResponse{
		users,
		count,
	}

	return response, nil
}

func (service *userService) UpdateUser(userUpdating *models.User) (*models.User, error) {
	user, err := service.storage.GetUser(userUpdating.ID)
	if err != nil {
		return nil, NewEntityNotFoundError("user", userUpdating.ID)
	}

	if err := validator.Validate(userUpdating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	user.SetFields(userUpdating)
	if err := service.storage.UpdateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (service *userService) RemoveUser(user *models.User) error {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewEntityNotFoundError("user", user.ID)
	}

	if err := service.storage.RemoveUser(user); err != nil {
		return err
	}

	return nil
}

func (service *userService) AttachPoliciesByUser(user *models.User, policies *[]models.Policy) error {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewEntityNotFoundError("user", user.ID)
	}

	for _, policy := range *policies {
		if _, err := service.storage.GetPolicy(policy.ID); err != nil {
			return NewEntityNotFoundError("policy", policy.ID)
		}
	}

	if err := service.storage.AttachPoliciesByUser(user, policies); err != nil {
		return err
	}

	return nil
}

func (service *userService) DetachPolicyByUser(user *models.User, policy *models.Policy) error {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewEntityNotFoundError("user", user.ID)
	}
	if _, err := service.storage.GetPolicy(policy.ID); err != nil {
		return NewEntityNotFoundError("policy", policy.ID)
	}

	if err := service.storage.DetachPolicyByUser(user, policy); err != nil {
		return err
	}

	return nil
}

func (service *userService) GetPoliciesByUser(user *models.User) (*[]models.Policy, error) {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return nil, NewEntityNotFoundError("user", user.ID)
	}

	policies, err := service.storage.GetPoliciesByUser(user)
	if err != nil {
		return nil, err
	}

	return policies, nil
}

func (service *userService) GetGroupsByUser(user *models.User) (*[]models.Group, error) {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return nil, NewEntityNotFoundError("user", user.ID)
	}

	groups, err := service.storage.GetGroupsByUser(user)
	if err != nil {
		return nil, err
	}

	return groups, err
}
