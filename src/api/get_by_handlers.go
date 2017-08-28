package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/services"
	"net/http"
	"strconv"
)

func (a *Api) getUsersByGroupHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return services.NewInvalidQueryError("query params is not valid")
	}
	group := &models.Group{ID: id}

	users, err := a.relationsService.GetUsersByGroup(group)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (a *Api) getPermissionsByPolicyHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return services.NewInvalidQueryError("query params is not valid")
	}
	policy := &models.Policy{ID: id}

	permissions, err := a.relationsService.GetPermissionsByPolicy(policy)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, permissions)
}

func (a *Api) getPoliciesByUserHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return services.NewInvalidQueryError("query params is not valid")
	}
	user := &models.User{ID: id}

	policies, err := a.relationsService.GetPoliciesByUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}

func (a *Api) getPoliciesByGroupHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return services.NewInvalidQueryError("query params is not valid")
	}
	group := &models.Group{ID: id}

	policies, err := a.relationsService.GetPoliciesByGroup(group)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, policies)
}
