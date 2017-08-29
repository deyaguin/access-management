package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
)

type users struct {
	Users *[]models.User `validate:"required"`
}
type policies struct {
	Policies *[]models.Policy `validate:"required"`
}

func (a *Api) createGroup(c echo.Context) error {
	groupCreating := &models.Group{}

	if err := c.Bind(groupCreating); err != nil {
		return NewUnprocessableBodyError("body is unprocessable")
	}
	group, err := a.groupService.CreateGroup(groupCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, group)
}

func (a *Api) getGroups(c echo.Context) error {
	groups, err := a.groupService.GetGroups()
	if err != nil {
		return NewInvalidQueryError("GroupID", string(3))
	}

	return c.JSON(http.StatusOK, groups)
}

func (a *Api) getGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("GroupID", c.Param("id"))
	}

	groups, err := a.groupService.GetGroup(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, groups)
}

func (a *Api) updateGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("GroupID", string(id))
	}

	groupUpdating := &models.Group{ID: id}
	if err := c.Bind(groupUpdating); err != nil {
		return NewUnprocessableBodyError("body is unprocessable")
	}

	group, err := a.groupService.UpdateGroup(groupUpdating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, group)
}

func (a *Api) removeGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("GroupID", string(id))
	}

	group := &models.Group{ID: id}
	if err := a.groupService.RemoveGroup(group); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) addUsersToGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("GroupID", string(id))
	}
	group := &models.Group{ID: id}

	users := new(users)
	if err = c.Bind(users); err != nil {
		return NewUnprocessableBodyError("body is unprocessable")
	}

	if err = a.groupService.AddUsersToGroup(group, users.Users); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) removeUserFromGroup(c echo.Context) error {
	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return NewInvalidQueryError("GroupID", string(groupId))
	}
	group := &models.Group{ID: groupId}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return err
	}
	user := &models.User{ID: userId}

	if err = a.groupService.RemoveUserFromGroup(group, user); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) getUsersByGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("GroupID", string(id))
	}
	group := &models.Group{ID: id}

	users, err := a.groupService.GetUsersByGroup(group)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (a *Api) attachPoliciesByGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("GroupID", string(id))
	}
	group := &models.Group{ID: id}

	policies := new(policies)
	if err = c.Bind(policies); err != nil {
		return NewUnprocessableBodyError("body is unprocessable")
	}

	if err = a.groupService.AttachPoliciesByGroup(group, policies.Policies); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) detachPolicyByGroup(c echo.Context) error {
	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return NewInvalidQueryError("GroupID", string(groupId))
	}
	group := &models.Group{ID: groupId}

	policyId, err := strconv.Atoi(c.Param("policyId"))
	if err != nil {
		return NewInvalidQueryError("PolicyID", string(policyId))
	}
	policy := &models.Policy{ID: policyId}

	if err = a.groupService.DetachPolicyByGroup(group, policy); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) getPoliciesByGroupHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("GroupID", string(id))
	}
	group := &models.Group{ID: id}

	policies, err := a.groupService.GetPoliciesByGroup(group)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}
