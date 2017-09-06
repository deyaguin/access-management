package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
)

func (a *API) createGroup(c echo.Context) error {
	groupCreating := &models.Group{}

	if err := c.Bind(groupCreating); err != nil {
		return NewUnprocessableBodyError()
	}

	group, err := a.CreateGroup(groupCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, group)
}

func (a *API) getGroups(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return NewInvalidQueryError(
			"page",
			c.QueryParam("page"),
		)
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return NewInvalidQueryError("per_page",
			c.QueryParam("per_page"),
		)
	}

	if err := checkPaginationParams(page, perPage); err != nil {
		return err
	}

	groups, err := a.GetGroups(page, perPage)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, groups)
}

func (a *API) getGroup(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return NewInvalidQueryError(
			"groupID",
			c.Param("groupID"),
		)
	}

	group, err := a.GetGroup(groupID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, group)
}

func (a *API) updateGroup(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return NewInvalidQueryError(
			"groupID",
			c.Param("groupID"),
		)
	}

	groupUpdating := &models.Group{ID: groupID}
	if err := c.Bind(groupUpdating); err != nil {
		return NewUnprocessableBodyError()
	}

	group, err := a.UpdateGroup(groupUpdating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, group)
}

func (a *API) removeGroup(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return NewInvalidQueryError(
			"groupID",
			c.Param("groupID"),
		)
	}

	if err := a.RemoveGroup(groupID); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *API) addUsersToGroup(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return NewInvalidQueryError(
			"groupID",
			c.Param("groupID"),
		)
	}
	group := &models.Group{ID: groupID}

	users := new(users)
	if err = c.Bind(users); err != nil {
		return NewUnprocessableBodyError()
	}

	if err = a.AddUsersToGroup(group, users.Users); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *API) removeUserFromGroup(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return NewInvalidQueryError(
			"groupID",
			c.Param("groupID"),
		)
	}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return NewInvalidQueryError(
			"userID",
			c.Param("userId"),
		)
	}

	if err = a.RemoveUserFromGroup(groupID, userId); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *API) getUsersByGroup(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return NewInvalidQueryError(
			"groupID",
			c.Param("groupID"),
		)
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return NewInvalidQueryError(
			"page",
			c.QueryParam("page"),
		)
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return NewInvalidQueryError(
			"per_page",
			c.Param("page"),
		)
	}

	if err := checkPaginationParams(page, perPage); err != nil {
		return err
	}

	users, err := a.GetUsersByGroup(
		groupID,
		page,
		perPage,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (a *API) attachPoliciesByGroup(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return NewInvalidQueryError(
			"groupID",
			c.Param("groupID"),
		)
	}
	group := &models.Group{ID: groupID}

	policies := new(policies)
	if err = c.Bind(policies); err != nil {
		return NewUnprocessableBodyError()
	}

	if err = a.AttachPoliciesByGroup(
		group,
		policies.Policies,
	); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *API) detachPolicyByGroup(c echo.Context) error {
	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return NewInvalidQueryError(
			"groupID",
			c.Param("groupID"),
		)
	}

	policyId, err := strconv.Atoi(c.Param("policyId"))
	if err != nil {
		return NewInvalidQueryError(
			"policyID",
			c.Param("policyId"),
		)
	}

	if err = a.DetachPolicyByGroup(groupID, policyId); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *API) getPoliciesByGroupHandler(c echo.Context) error {

	groupID, err := strconv.Atoi(c.Param("groupID"))
	if err != nil {
		return NewInvalidQueryError(
			"groupID",
			c.Param("groupID"),
		)
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return NewInvalidQueryError(
			"page",
			c.QueryParam("page"),
		)
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return NewInvalidQueryError(
			"per_page",
			c.QueryParam("per_page"),
		)
	}

	if err := checkPaginationParams(page, perPage); err != nil {
		return err
	}

	policies, err := a.GetPoliciesByGroup(
		groupID,
		page,
		perPage)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}
