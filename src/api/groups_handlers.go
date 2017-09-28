package api

import (
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
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
	if c.QueryParam("page") == "" || c.QueryParam("per_page") == "" {
		groups, err := a.GetGroups(1, 10, c.QueryParam("group_name"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, groups)
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
		return NewInvalidQueryError("per_page",
			c.QueryParam("per_page"),
		)
	}

	if err := checkPaginationParams(page, perPage); err != nil {
		return err
	}

	groupName := c.QueryParam("group_name")

	groups, err := a.GetGroups(page, perPage, groupName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, groups)
}

func (a *API) getAllGroups(c echo.Context) error {
	groups, err := a.GetAllGroups()
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

	userId, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return NewInvalidQueryError(
			"userID",
			c.Param("userID"),
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

	users, err := a.GetUsersByGroup(groupID)
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

	policyId, err := strconv.Atoi(c.Param("policyID"))
	if err != nil {
		return NewInvalidQueryError(
			"policyID",
			c.Param("policyID"),
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

	policies, err := a.GetPoliciesByGroup(groupID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}
