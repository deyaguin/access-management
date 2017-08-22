package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/accessControl/app/models"
	"net/http"
	"strconv"
)

type groupParams struct {
	Name string `validation:"required"`
}

func (a *Api) createGroup(c echo.Context) (err error) {
	gP := new(groupParams)
	c.Bind(gP)
	if err = c.Validate(gP); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	group := &models.Group{Name: gP.Name}
	if err = a.DB.CreateGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, "created")
}

func (a *Api) getGroups(c echo.Context) error {
	groups, err := a.DB.GetGroups()
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, groups)
}

func (a *Api) getGroup(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	group, err := a.DB.GetGroup(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, group)
}

func (a *Api) updateGroup(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err := a.DB.GetGroup(id); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	gP := new(groupParams)
	c.Bind(gP)
	if err = c.Validate(gP); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var group = &models.Group{ID: id, Name: gP.Name}
	if err = a.DB.UpdateGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, "updated")
}

func (a *Api) removeGroup(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if _, err := a.DB.GetGroup(id); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	group := &models.Group{ID: id}
	if err = a.DB.DeleteGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "remove")
}
