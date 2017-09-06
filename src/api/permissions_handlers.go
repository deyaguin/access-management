package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"net/http"
	"strconv"
)

func (a *API) updatePermission(c echo.Context) error {
	permissionID, err := strconv.Atoi(c.Param("permissionID"))
	if err != nil {
		return NewInvalidQueryError(
			"PermissionID",
			c.Param("permissionID"),
		)
	}

	permissionUpdating := &models.Permission{ID: permissionID}
	if err := c.Bind(permissionUpdating); err != nil {
		return NewUnprocessableBodyError()
	}

	permission, err := a.UpdatePermission(permissionUpdating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, permission)
}

func (a *API) RemovePermission(c echo.Context) error {
	return *new(error)
}
