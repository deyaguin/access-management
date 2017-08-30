package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
)

func (a *Api) createUser(c echo.Context) error {
	userCreating := &models.User{}

	if err := c.Bind(userCreating); err != nil {
		return NewUnprocessableBodyError()
	}

	user, err := a.userService.CreateUser(userCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (a *Api) getUsers(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return NewInvalidQueryError("page", c.QueryParam("page"))
	}

	perPage, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		return NewInvalidQueryError("per_page", c.QueryParam("per_page"))
	}

	users, err := a.userService.GetUsers(page, perPage)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (a *Api) getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("UserID", c.Param("id"))
	}

	user, err := a.userService.GetUser(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (a *Api) updateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("UserID", c.Param("id"))
	}

	userUpdating := &models.User{ID: id}
	if err := c.Bind(userUpdating); err != nil {
		return NewUnprocessableBodyError()
	}

	user, err := a.userService.UpdateUser(userUpdating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (a *Api) removeUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("UserID", c.Param("id"))
	}

	if err := a.userService.RemoveUser(id); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) attachPoliciesByUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("UserID", c.Param("id"))
	}
	user := &models.User{ID: id}

	policies := new(policies)
	if err = c.Bind(policies); err != nil {
		return NewUnprocessableBodyError()
	}

	if err = a.userService.AttachPoliciesByUser(user, policies.Policies); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) detachPolicyByUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return NewInvalidQueryError("UserID", c.Param("userId"))
	}

	policyId, err := strconv.Atoi(c.Param("policyId"))
	if err != nil {
		return NewInvalidQueryError("PolicyID", c.Param("policyId"))
	}

	if err = a.userService.DetachPolicyByUser(userId, policyId); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) getPoliciesByUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("UserID", c.Param("id"))
	}

	policies, err := a.userService.GetPoliciesByUser(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}
