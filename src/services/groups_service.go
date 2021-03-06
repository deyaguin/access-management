package services

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/storage"

	"gopkg.in/validator.v2"
	"gitlab/nefco/access-management-system/src/utils"
)

type GroupsService interface {
	CreateGroup(*models.Group) (*models.Group, error)
	GetGroup(int) (*models.Group, error)
	GetGroups(int, int, string) (*paginationItems, error)
	GetAllGroups() (*items, error)
	GetGroupsByEntry(string) (*items, error)
	UpdateGroup(*models.Group) (*models.Group, error)
	RemoveGroup(int) error

	AddUsersToGroup(*models.Group, *[]models.User) error
	RemoveUserFromGroup(int, int) error
	GetUsersByGroup(int) (*items, error)
	AttachPoliciesByGroup(*models.Group, *[]models.Policy) error
	DetachPolicyByGroup(int, int) error
	GetPoliciesByGroup(int) (*items, error)
}

type groupsService struct {
	storage storage.DB
}

func NewGroupsService(
	storage storage.DB,
) GroupsService {
	return &groupsService{
		storage,
	}
}

func (service *groupsService) CreateGroup(
	groupCreating *models.Group,
) (*models.Group, error) {
	if err := validator.Validate(groupCreating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	group := new(models.Group)
	group.SetFields(groupCreating)

	if err := service.storage.CreateGroup(group); err != nil {
		return nil, NewEntityCreateError(err.Error())
	}

	return group, nil
}

func (service *groupsService) GetGroup(
	groupID int,
) (*models.Group, error) {
	group, err := service.storage.GetGroup(groupID)
	if err != nil {
		return nil, NewEntityNotFoundError("group", groupID)
	}

	return group, nil
}

func (service *groupsService) GetGroups(
	page int,
	perPage int,
	name string,
) (*paginationItems, error) {
	groups, err := service.storage.GetGroups(page, perPage, name)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	count, err := service.storage.GetGroupsCount(
		utils.LikeQuery("groups", "name", name),
	)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	items := &paginationItems{
		groups,
		count,
		perPage,
		page,
	}

	return items, nil
}

func (service *groupsService) GetAllGroups() (*items, error) {
	groups, err := service.storage.GetAllGroups()
	if err != nil {
		return nil, err
	}

	return &items{groups}, nil
}

func (service *groupsService) GetGroupsByEntry(name string) (*items, error) {
	groups, err := service.storage.GetGroupsByEntry(name)
	if err != nil {
		return nil, err
	}

	return &items{groups}, nil
}

func (service *groupsService) UpdateGroup(
	groupUpdating *models.Group,
) (*models.Group, error) {
	group, err := service.storage.GetGroup(groupUpdating.ID)
	if err != nil {
		return nil, NewEntityNotFoundError("group", groupUpdating.ID)
	}

	if err := validator.Validate(groupUpdating); err != nil {
		return nil, NewValidationError(err.Error())
	}

	group.SetFields(groupUpdating)
	if err := service.storage.UpdateGroup(group); err != nil {
		return nil, NewEntityUpdateError(err.Error())
	}

	return group, nil
}

func (service *groupsService) RemoveGroup(
	groupID int,
) error {
	group, err := service.storage.GetGroup(groupID)
	if err != nil {
		return NewEntityNotFoundError("group", groupID)
	}

	if err := service.storage.RemoveGroup(group); err != nil {
		return NewEntityRemoveError(err.Error())
	}

	return nil
}

func (service *groupsService) AddUsersToGroup(
	group *models.Group,
	users *[]models.User,
) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewEntityNotFoundError("group", group.ID)
	}

	for _, user := range *users {
		if _, err := service.storage.GetUser(user.ID); err != nil {
			return NewEntityNotFoundError("user", user.ID)
		}
	}

	if err := service.storage.AddUsersToGroup(group, users); err != nil {
		return NewEntityCreateError(err.Error())
	}

	return nil
}

func (service *groupsService) RemoveUserFromGroup(
	groupID int,
	userID int,
) error {
	group, err := service.storage.GetGroup(groupID)
	if err != nil {
		return NewEntityNotFoundError("group", groupID)
	}

	user, err := service.storage.GetUser(userID)
	if err != nil {
		return NewEntityNotFoundError("user", userID)
	}

	if err := service.storage.RemoveUserFromGroup(group, user); err != nil {
		return NewEntityRemoveError(err.Error())
	}

	return nil
}

func (service *groupsService) GetUsersByGroup(
	groupID int,
) (*items, error) {
	group, err := service.storage.GetGroup(groupID)
	if err != nil {
		return nil, NewEntityNotFoundError("group", groupID)
	}

	users, err := service.storage.GetUsersByGroup(group)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	result := &items{users}

	return result, nil
}

func (service *groupsService) AttachPoliciesByGroup(
	group *models.Group,
	policies *[]models.Policy,
) error {
	if _, err := service.storage.GetGroup(group.ID); err != nil {
		return NewEntityNotFoundError(
			"group",
			group.ID,
		)
	}

	for _, policy := range *policies {
		if _, err := service.storage.GetPolicy(policy.ID); err != nil {
			return NewEntityNotFoundError(
				"policy",
				policy.ID,
			)
		}
	}

	if err := service.storage.AttachPoliciesByGroup(group, policies); err != nil {
		return NewEntityCreateError(err.Error())
	}

	return nil
}

func (service *groupsService) DetachPolicyByGroup(
	groupID int,
	policyID int,
) error {
	group, err := service.storage.GetGroup(groupID)
	if err != nil {
		return NewEntityNotFoundError("group", groupID)
	}

	policy, err := service.storage.GetPolicy(policyID)
	if err != nil {
		return NewEntityNotFoundError("policy", policyID)
	}

	if err := service.storage.DetachPolicyByGroup(group, policy); err != nil {
		return NewEntityRemoveError(err.Error())
	}

	return nil
}

func (service *groupsService) GetPoliciesByGroup(
	groupID int,
) (*items, error) {
	group, err := service.storage.GetGroup(groupID)
	if err != nil {
		return nil, NewEntityNotFoundError("group", groupID)
	}

	policies, err := service.storage.GetPoliciesByGroup(group)
	if err != nil {
		return nil, NewGetEntitiesError(err.Error())
	}

	result := &items{policies}

	return result, nil
}
