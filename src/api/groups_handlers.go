package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/services"
	"net/http"
	"strconv"
)

func (a *Api) createGroup(c echo.Context) error {
	groupCreating := &models.Group{}

	if err := c.Bind(groupCreating); err != nil {
		return services.NewUnprocessableBodyError("body is unprocessable")
	}
	group, err := a.groupService.CreateGroup(groupCreating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, group)
}

func (a *Api) getGroups(c echo.Context) error {
	groups, err := a.groupService.GetGroups()
	if err != nil {
		return services.NewInvalidQueryError("query params is not valid")
	}

	return c.JSON(http.StatusOK, groups)
}

func (a *Api) getGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return services.NewInvalidQueryError("query params is not valid")
	}

	groups, err := a.groupService.GetGroup(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, groups)
}

func (a *Api) updateGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return services.NewInvalidQueryError("query params is not valid")
	}

	groupUpdating := &models.Group{ID: id}
	if err := c.Bind(groupUpdating); err != nil {
		return services.NewUnprocessableBodyError("body is unprocessable")
	}

	group, err := a.groupService.UpdateGroup(groupUpdating)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, group)
}

func (a *Api) removeGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return services.NewInvalidQueryError("query params is not valid")
	}

	group := &models.Group{ID: id}
	if err := a.groupService.RemoveGroup(group); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
