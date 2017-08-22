package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/accessControl/app/models"
	"net/http"
	"strconv"
)

type permissionParams struct {
	Resourse string `validate:"required"`
	Access   bool   `validate:"required"`
	ActionID int    `validate:"required"`
}

func (a *Api) createPermission(c echo.Context) (err error) {
	permission := new(models.Permission)
	id, err := strconv.Atoi(c.Param("id"))
	permission, err = a.DB.GetPermission(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	c.Bind(permission)
	if err = c.Validate(permission); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err = a.DB.UpdatePermission(permission); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "updated")
}

func (a *Api) updatePermission(c echo.Context) (err error) {
	permission := new(models.Permission)
	id, err := strconv.Atoi(c.Param("id"))
	permission, err = a.DB.GetPermission(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}
	c.Bind(permission)
	if err = c.Validate(permission); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err = a.DB.UpdatePermission(permission); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "updated")
}

func (a *Api) removePermission(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.ParamValues()[1])
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err := a.DB.GetPermission(id); err != nil {
		return c.JSON(http.StatusNotFound, "permission not found")
	}
	permission := &models.Permission{ID: id}
	if err = a.DB.DeletePermission(permission); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "remove")
}
