package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/services"
	"net/http"
	"strconv"
)

func (a *Api) updatePermission(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return services.NewInvalidQueryError("query params is not valid")
	}

	permissionUpdating := &models.Permission{ID: id}
	if err := c.Bind(permissionUpdating); err != nil {
		return services.NewUnprocessableBodyError("body is unprocessable")
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
