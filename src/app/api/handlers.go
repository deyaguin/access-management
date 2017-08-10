package api

import (
	"github.com/labstack/echo"
	"app/models"
	"net/http"
)

func (a *Api) CreateUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	a.DB.CreateUser(user)
	return c.JSON(http.StatusOK, user)
}

func (a *Api) GetUsers(c echo.Context) error {
	users := a.DB.GetUsers()
	return c.JSON(http.StatusOK, users)
}

func (a *Api) CreateGroup(c echo.Context) error {
	group := new(models.Group)
	c.Bind(group)
	a.DB.CreateGroup(group)
	return c.JSON(http.StatusOK, group)
}

func (a *Api) GetGroups(c echo.Context) error {
	groups := a.DB.GetGroups()
	return c.JSON(http.StatusOK, groups)
}

func (a *Api) CreatePolicy(c echo.Context) error {
	policy := new(models.Policy)
	c.Bind(policy)
	a.DB.CreatePolicy(policy)
	return c.JSON(http.StatusOK, policy)
}

func (a *Api) GetPolicies(c echo.Context) error {
	policies :=	a.DB.GetPolicies()
	return c.JSON(http.StatusOK, policies)
}

func (a *Api) CheckPermissions(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	permissions := a.FormListOfPermissions(user)
	return c.JSON(http.StatusOK, permissions)
}
