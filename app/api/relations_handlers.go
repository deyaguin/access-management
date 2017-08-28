package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/accessControl/app/models"
	"net/http"
	"strconv"
)

type users struct {
	Users *[]models.User
}
type groups struct {
	Groups *[]models.Group `validate:"required"`
}
type policies struct {
	Policies *[]models.Policy `validate:"required"`
}
type permissions struct {
	Permissions *[]models.Permission `validate:"required"`
}

func (a *Api) addUsersToGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	group := &models.Group{ID: id}

	users := new(users)
	c.Bind(users)

	if err = a.relationsService.AddUsersToGroup(group, users.Users); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) removeUserFromGroup(c echo.Context) error {
	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return err
	}
	group := &models.Group{ID: groupId}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return err
	}
	user := &models.User{ID: userId}

	if err = a.relationsService.RemoveUserFromGroup(group, user); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (a *Api) addPermissionsToPolicy(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	policy := &models.Policy{ID: id}

	permissions := new(permissions)
	c.Bind(permissions)

	if err = a.relationsService.AddPermissionsToPolicy(policy, permissions.Permissions); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) removePermissionFromPolicy(c echo.Context) error {
	policyId, err := strconv.Atoi(c.Param("policyId"))
	if err != nil {
		return err
	}
	policy := &models.Policy{ID: policyId}

	permissionId, err := strconv.Atoi(c.Param("permissionId"))
	if err != nil {
		return err
	}
	permission := &models.Permission{ID: permissionId}

	if err = a.relationsService.RemovePermissionFromPolicy(policy, permission); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (a *Api) attachPoliciesByUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	user := &models.User{ID: id}

	policies := new(policies)
	c.Bind(policies)

	if err = a.relationsService.AttachPoliciesByUser(user, policies.Policies); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) detachPolicyByUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return err
	}
	user := &models.User{ID: userId}

	policyId, err := strconv.Atoi(c.Param("policyId"))
	if err != nil {
		return err
	}
	policy := &models.Policy{ID: policyId}

	if err = a.relationsService.DetachPolicyByUser(user, policy); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (a *Api) attachPoliciesByGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	group := &models.Group{ID: id}

	policies := new(policies)
	c.Bind(policies)

	if err = a.relationsService.AttachPoliciesByGroup(group, policies.Policies); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) detachPolicyByGroup(c echo.Context) error {
	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return err
	}
	group := &models.Group{ID: groupId}

	policyId, err := strconv.Atoi(c.Param("policyId"))
	if err != nil {
		return err
	}
	policy := &models.Policy{ID: policyId}

	if err = a.relationsService.DetachPolicyByGroup(group, policy); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
