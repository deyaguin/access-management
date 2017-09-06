package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
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

	users, err := a.GetUsers(page, perPage)
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

	policyId, err := strconv.Atoi(c.Param("policyId"))
	if err != nil {
		return NewInvalidQueryError(
			"PolicyID",
			c.Param("policyId"),
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

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return NewInvalidQueryError(
			"UserID",
			c.Param("userID"),
		)
	}

	policies, err := a.GetPoliciesByUser(
		userID,
		page,
		perPage,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}
