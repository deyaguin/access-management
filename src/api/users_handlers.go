package api

import (
	"fmt"
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/services"
	"net/http"
	"strconv"
)

func (a *Api) createUser(c echo.Context) error {
	userCreating := &models.User{}

	if err := c.Bind(userCreating); err != nil {
		return services.NewUnprocessableBodyError("body is unprocessable")
	}

	user, err := a.userService.CreateUser(userCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (a *Api) getUsers(c echo.Context) error {
	users, err := a.userService.GetUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (a *Api) getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return services.NewInvalidQueryError("query params is not valid")
	}

	user, err := a.userService.GetUser(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (a *Api) updateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return services.NewInvalidQueryError("query params is not valid")
	}

	userUpdating := &models.User{ID: id}
	if err := c.Bind(userUpdating); err != nil {
		return services.NewUnprocessableBodyError("body is unprocessable")
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
		return services.NewInvalidQueryError("query params is not valid")
	}

	user := &models.User{ID: id}
	if err := a.userService.RemoveUser(user); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}