package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
)

func (a *Api) updatePermission(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return NewInvalidQueryError("PermissionID", string(id))
	}

	permissionUpdating := &models.Permission{ID: id}
	if err := c.Bind(permissionUpdating); err != nil {
		return NewUnprocessableBodyError("body is unprocessable")
	}

	permission, err := a.permissionService.UpdatePermission(permissionUpdating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, permission)
}

func (a *Api) RemovePermission(c echo.Context) error {
	return *new(error)
}
