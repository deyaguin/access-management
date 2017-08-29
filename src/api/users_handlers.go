package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
)

type groups struct {
	Groups *[]models.Group `validate:"required"`
}

func (a *Api) createUser(c echo.Context) error {
	userCreating := &models.User{}

	if err := c.Bind(userCreating); err != nil {
		return NewUnprocessableBodyError("body is unprocessable")
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
		return NewInvalidQueryError("GroupID", string(3))
	}

	users, err := a.userService.GetUsers(page)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (a *Api) getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("UserID", string(id))
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
		return NewInvalidQueryError("UserID", string(id))
	}

	userUpdating := &models.User{ID: id}
	if err := c.Bind(userUpdating); err != nil {
		return NewUnprocessableBodyError("body is unprocessable")
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
		return NewInvalidQueryError("UserID", string(id))
	}

	user := &models.User{ID: id}
	if err := a.userService.RemoveUser(user); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) attachPoliciesByUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("UserID", string(id))
	}
	user := &models.User{ID: id}

	policies := new(policies)
	if err = c.Bind(policies); err != nil {
		return NewUnprocessableBodyError("body is unprocessable")
	}

	if err = a.userService.AttachPoliciesByUser(user, policies.Policies); err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (a *Api) detachPolicyByUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return NewInvalidQueryError("UserID", string(userId))
	}
	user := &models.User{ID: userId}

	policyId, err := strconv.Atoi(c.Param("policyId"))
	if err != nil {
		return NewInvalidQueryError("PolicyID", string(policyId))
	}
	policy := &models.Policy{ID: policyId}

	if err = a.userService.DetachPolicyByUser(user, policy); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (a *Api) getPoliciesByUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("UserID", string(id))
	}
	user := &models.User{ID: id}

	policies, err := a.userService.GetPoliciesByUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}
