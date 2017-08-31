package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"
	"gopkg.in/validator.v2"
)

type UserService interface {
	CreateUser(*models.User) (*models.User, error)
	GetUser(int) (*models.User, error)
	GetUsers(int, int) (*items, error)
	UpdateUser(*models.User) (*models.User, error)
	RemoveUser(int) error

	AttachPoliciesByUser(*models.User, *[]models.Policy) error
	DetachPolicyByUser(int, int) error
	GetPoliciesByUser(int, int, int) (*items, error)
	GetGroupsByUser(int, int, int) (*items, error)
}

type userService struct {
	storage storage.DB
}

func NewUserService(
	storage storage.DB,
) UserService {
	return &userService{
		storage,
	}
}

func (service *userService) CreateUser(
	userCreating *models.User,
) (*models.User, error) {
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

func (service *userService) GetUser(
	id int,
) (*models.User, error) {
	user, err := service.storage.GetUser(id)
	if err != nil {
		return nil, NewEntityNotFoundError("user", id)
	}

	return user, nil
}

func (service *userService) GetUsers(
	page int,
	perPage int,
) (*items, error) {
	users, err := service.storage.GetUsers(page, perPage)
	if err != nil {
		return nil, err
	}

	count, err := service.storage.GetUsersCount()
	if err != nil {
		return nil, err
	}

	response := &items{
		users,
		count,
	}

	return response, nil
}

func (service *userService) UpdateUser(
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
		return nil, err
	}

	return user, nil
}

func (service *userService) RemoveUser(
	userId int,
) error {
	user, err := service.storage.GetUser(userId)
	if err != nil {
		return NewEntityNotFoundError("user", userId)
	}

	if err := service.storage.RemoveUser(user); err != nil {
		return err
	}

	return nil
}

func (service *userService) AttachPoliciesByUser(
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
		return err
	}

	return nil
}

func (service *userService) DetachPolicyByUser(
	userId int,
	policyId int,
) error {
	user, err := service.storage.GetUser(userId)
	if err != nil {
		return NewEntityNotFoundError("user", userId)
	}

	policy, err := service.storage.GetPolicy(policyId)
	if err != nil {
		return NewEntityNotFoundError("policy", policyId)
	}

	if err := service.storage.DetachPolicyByUser(user, policy); err != nil {
		return err
	}

	return nil
}

func (service *userService) GetPoliciesByUser(
	userId int,
	page int,
	perPage int,
) (*items, error) {
	user, err := service.storage.GetUser(userId)
	if err != nil {
		return nil, NewEntityNotFoundError("user", userId)
	}

	policies, count, err := service.storage.GetPoliciesByUser(
		user,
		&page,
		&perPage,
	)
	if err != nil {
		return nil, err
	}

	items := &items{
		policies,
		count,
	}

	return items, nil
}

func (service *userService) GetGroupsByUser(
	userId int,
	page int,
	perPage int,
) (*items, error) {
	user, err := service.storage.GetUser(userId)
	if err != nil {
		return nil, NewEntityNotFoundError("user", userId)
	}

	groups, count, err := service.storage.GetGroupsByUser(
		user,
		&page,
		&perPage,
	)
	if err != nil {
		return nil, err
	}

	items := &items{
		groups,
		count,
	}

	return items, err
}
