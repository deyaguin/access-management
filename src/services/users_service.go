package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"

	"gopkg.in/validator.v2"
	"gitlab/nefco/access-management-system/src/utils"
)

type UsersService interface {
	CreateUser(*models.User) (*models.User, error)
	GetUser(int) (*models.User, error)
	GetUsers(int, int, string) (*paginationItems, error)
	GetAllUsers() (*items, error)
	GetUsersByEntry(string) (*items, error)
	UpdateUser(*models.User) (*models.User, error)
	RemoveUser(int) error

	AttachPoliciesByUser(*models.User, *[]models.Policy) error
	DetachPolicyByUser(int, int) error
	GetPoliciesByUser(int) (*items, error)
	GetGroupsByUser(int) (*items, error)
}

type usersService struct {
	storage storage.DB
}

func NewUsersService(
	storage storage.DB,
) UsersService {
	return &usersService{
		storage,
	}
}

func (service *usersService) CreateUser(
	userCreating *models.User,
) (*models.User, error) {
	if err := validator.Validate(userCreating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	user := new(models.User)
	user.SetFields(userCreating)

	if err := service.storage.CreateUser(user); err != nil {
		return nil, NewEntityCreateError(err.Error())
	}

	return user, nil
}

func (service *usersService) GetUser(
	userID int,
) (*models.User, error) {
	user, err := service.storage.GetUser(userID)
	if err != nil {

		return nil, NewEntityNotFoundError("user", userID)
	}

	return user, nil
}

func (service *usersService) GetUsers(
	page int,
	perPage int,
	name string,
) (*paginationItems, error) {
	users, err := service.storage.GetUsers(page, perPage, name)
	if err != nil {
		return nil, err
	}

	total, err := service.storage.GetUsersCount(
		utils.LikeQuery("users", "name", name))
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	response := &paginationItems{
		users,
		total,
		perPage,
		page,
	}

	return response, nil
}

func (service *usersService) GetAllUsers() (*items, error) {
	users, err := service.storage.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return &items{users}, nil
}

func (service *usersService) GetUsersByEntry(name string) (*items, error) {
	users, err := service.GetUsersByEntry(name)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *usersService) UpdateUser(
	userUpdating *models.User,
) (*models.User, error) {
	user, err := service.storage.GetUser(userUpdating.ID)
	if err != nil {
		return nil, NewEntityNotFoundError("user", userUpdating.ID)
	}

	if err := validator.Validate(userUpdating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	user.SetFields(userUpdating)
	if err := service.storage.UpdateUser(user); err != nil {
		return nil, NewEntityUpdateError(err.Error())
	}

	return user, nil
}

func (service *usersService) RemoveUser(
	userID int,
) error {
	user, err := service.storage.GetUser(userID)
	if err != nil {
		return NewEntityNotFoundError("user", userID)
	}

	if err := service.storage.RemoveUser(user); err != nil {
		return NewEntityRemoveError(err.Error())
	}

	return nil
}

func (service *usersService) AttachPoliciesByUser(
	user *models.User,
	policies *[]models.Policy,
) error {
	if _, err := service.storage.GetUser(user.ID); err != nil {
		return NewEntityNotFoundError("user", user.ID)
	}

	for _, policy := range *policies {
		if _, err := service.storage.GetPolicy(policy.ID); err != nil {
			return NewEntityNotFoundError("policy", policy.ID)
		}
	}

	if err := service.storage.AttachPoliciesByUser(user, policies); err != nil {
		return NewEntityCreateError(err.Error())
	}

	return nil
}

func (service *usersService) DetachPolicyByUser(
	userID int,
	policyID int,
) error {
	user, err := service.storage.GetUser(userID)
	if err != nil {
		return NewEntityNotFoundError("user", userID)
	}

	policy, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return NewEntityNotFoundError("policy", policyID)
	}

	if err := service.storage.DetachPolicyByUser(user, policy); err != nil {
		return NewEntityRemoveError(err.Error())
	}

	return nil
}

func (service *usersService) GetPoliciesByUser(
	userID int,
) (*items, error) {
	user, err := service.storage.GetUser(userID)
	if err != nil {
		return nil, NewEntityNotFoundError("user", userID)
	}

	policies, err := service.storage.GetPoliciesByUser(user)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	result := &items{policies}

	return result, nil
}

func (service *usersService) GetGroupsByUser(
	userID int,
) (*items, error) {
	user, err := service.storage.GetUser(userID)
	if err != nil {
		return nil, NewEntityNotFoundError("user", userID)
	}

	groups, err := service.storage.GetGroupsByUser(user)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	result := &items{groups}

	return result, err
}
