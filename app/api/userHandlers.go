package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/accessControl/app/models"
	"net/http"
	"strconv"
)

type userParams struct {
	Name string `validation:"required"`
}

func (a *Api) createUser(c echo.Context) (err error) {
	uP := new(userParams)
	c.Bind(uP)
	if err = c.Validate(uP); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	user := &models.User{Name: uP.Name}
	if err = a.DB.CreateUser(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, "created")
}

func (a *Api) getUsers(c echo.Context) error {
	users, err := a.DB.GetUsers()
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (a *Api) getUser(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	user, err := a.DB.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (a *Api) updateUser(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err := a.DB.GetUser(id); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	uP := new(userParams)
	c.Bind(uP)
	if err = c.Validate(uP); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var user = &models.User{ID: id, Name: uP.Name}
	if err = a.DB.UpdateUser(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "updated")
}

func (a *Api) removeUser(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if _, err := a.DB.GetUser(id); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	user := &models.User{ID: id}
	if err = a.DB.DeleteUser(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "remove")
}
