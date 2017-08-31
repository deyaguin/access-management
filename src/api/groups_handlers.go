package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
)

func (a *Api) createGroup(c echo.Context) error {
	groupCreating := &models.Group{}

	if err := c.Bind(groupCreating); err != nil {
		return NewUnprocessableBodyError()
	}
	group, err := a.groupService.CreateGroup(groupCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, group)
}

func (a *Api) getGroups(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return NewInvalidQueryError("page", c.QueryParam("page"))
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return NewInvalidQueryError("per_page", c.QueryParam("per_page"))
	}

	groups, err := a.groupService.GetGroups(page, perPage)
	if err != nil {
		return err
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
		return NewInvalidQueryError("GroupID", c.Param("id"))
	}

	groupUpdating := &models.Group{ID: id}
	if err := c.Bind(groupUpdating); err != nil {
		return NewUnprocessableBodyError()
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
		return NewInvalidQueryError("GroupID", c.Param("id"))
	}

	if err := a.groupService.RemoveGroup(id); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) addUsersToGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("GroupID", c.Param("id"))
	}
	group := &models.Group{ID: id}

	users := new(users)
	if err = c.Bind(users); err != nil {
		return NewUnprocessableBodyError()
	}

	if err = a.groupService.AddUsersToGroup(group, users.Users); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) removeUserFromGroup(c echo.Context) error {
	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return NewInvalidQueryError("GroupID", c.Param("groupId"))
	}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return NewInvalidQueryError("UserID", c.Param("userId"))
	}

	if err = a.groupService.RemoveUserFromGroup(groupId, userId); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) getUsersByGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("GroupID", c.Param("id"))
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return NewInvalidQueryError("page", c.QueryParam("page"))
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return NewInvalidQueryError("per_page", c.Param("page"))
	}

	users, err := a.groupService.GetUsersByGroup(id, page, perPage)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (a *Api) attachPoliciesByGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("GroupID", c.Param("id"))
	}
	group := &models.Group{ID: id}


	policies := new(policies)
	if err = c.Bind(policies); err != nil {
		return NewUnprocessableBodyError()
	}

	if err = a.groupService.AttachPoliciesByGroup(group, policies.Policies); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) detachPolicyByGroup(c echo.Context) error {
	groupId, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return NewInvalidQueryError("GroupID", c.Param("groupId"))
	}

	policyId, err := strconv.Atoi(c.Param("policyId"))
	if err != nil {
		return NewInvalidQueryError("PolicyID", c.Param("policyId"))
	}

	if err = a.groupService.DetachPolicyByGroup(groupId, policyId); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) getPoliciesByGroupHandler(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return NewInvalidQueryError("page", c.QueryParam("page"))
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return NewInvalidQueryError("per_page", c.QueryParam("per_page"))
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("GroupID", c.Param("id"))
	}

	policies, err := a.groupService.GetPoliciesByGroup(id, page, perPage)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}
