package api

import (
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (a *API) createUser(c echo.Context) error {
	userCreating := &models.User{}

	if err := c.Bind(userCreating); err != nil {
		return NewUnprocessableBodyError()
	}

	user, err := a.CreateUser(userCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (a *API) getUsers(c echo.Context) error {
	if c.QueryParam("page") == "" || c.QueryParam("per_page") == "" {
		users, err := a.GetUsers(1, 10, c.QueryParam("user_name"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, users)
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

	userName := c.QueryParam("user_name")

	users, err := a.GetUsers(page, perPage, userName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (a *API) getAllUsers(c echo.Context) error {
	users, err := a.GetAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (a *API) getUsersByEntry(c echo.Context) error {
	name := c.QueryParam("name")

	users, err := a.GetUsersByEntry(name)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (a *API) getUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return NewInvalidQueryError(
			"UserID",
			c.Param("userID"),
		)
	}

	user, err := a.GetUser(userID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (a *API) updateUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return NewInvalidQueryError(
			"UserID",
			c.Param("userID"),
		)
	}

	userUpdating := &models.User{ID: userID}
	if err := c.Bind(userUpdating); err != nil {
		return NewUnprocessableBodyError()
	}

	user, err := a.UpdateUser(userUpdating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (a *API) removeUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return NewInvalidQueryError(
			"UserID",
			c.Param("userID"),
		)
	}

	if err := a.RemoveUser(userID); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *API) attachPoliciesByUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return NewInvalidQueryError(
			"UserID",
			c.Param("userID"),
		)
	}
	user := &models.User{ID: userID}

	policies := new(policies)
	if err = c.Bind(policies); err != nil {
		return NewUnprocessableBodyError()
	}

	if err = a.AttachPoliciesByUser(
		user,
		policies.Policies,
	); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *API) detachPolicyByUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return NewInvalidQueryError(
			"UserID",
			c.Param("userID"),
		)
	}

	policyId, err := strconv.Atoi(c.Param("policyID"))
	if err != nil {
		return NewInvalidQueryError(
			"PolicyID",
			c.Param("policyID"),
		)
	}

	if err = a.DetachPolicyByUser(
		userID,
		policyId,
	); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *API) getPoliciesByUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return NewInvalidQueryError(
			"UserID",
			c.Param("userID"),
		)
	}

	policies, err := a.GetPoliciesByUser(userID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}

func (a *API) getGroupsByUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return NewInvalidQueryError(
			"UserID",
			c.Param("userID"),
		)
	}

	groups, err := a.GetGroupsByUser(userID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, groups)
}
