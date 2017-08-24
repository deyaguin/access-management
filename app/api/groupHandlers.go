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

func (a *Api) createGroup(c echo.Context) error {
	gParams := new(groupParams)
	c.Bind(gParams)
	if err := c.Validate(gParams); err != nil {
		return c.JSON(http.StatusCreated, err.Error())
	}
	group := &models.Group{Name: gParams.Name}
	if err := a.DB.CreateGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, group)
}

func (a *Api) getGroups(c echo.Context) error {
	groups, err := a.DB.GetGroups()
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, groups)
}

func (a *Api) getGroup(c echo.Context) error {
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

func (a *Api) updateGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err := a.DB.GetGroup(id); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	gParams := new(groupParams)
	c.Bind(gParams)
	if err = c.Validate(gParams); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var group = &models.Group{ID: id, Name: gParams.Name}
	if err = a.DB.UpdateGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, group)
}

func (a *Api) removeGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if _, err := a.DB.GetGroup(id); err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	group := &models.Group{ID: id}
	if err = a.DB.DeleteGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.NoContent(http.StatusOK)
}
