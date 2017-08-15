package api

import (
	"github.com/labstack/echo"
	"app/models"
	"net/http"
)

type checkParams struct {
	Resourse string
	Action uint
	ID uint
	Name string
}

func (a *Api) createUser(c echo.Context) (err error) {
	user := new(models.User)
	c.Bind(user)
	if err = c.Validate(user); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.CreateUser(user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

func (a *Api) getUsers(c echo.Context) error {
	users, err := a.DB.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (a *Api) deleteUser(c echo.Context) (err error) {
	user := new(models.User)
	c.Bind(user)
	if err = c.Validate(user); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.DeleteUser(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (a *Api) createGroup(c echo.Context) (err error) {
	group := new(models.Group)
	c.Bind(group)
	if err = c.Validate(group); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.CreateGroup(group); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, group)
}

func (a *Api) getGroups(c echo.Context) error {
	groups, err := a.DB.GetGroups()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, groups)
}

func (a *Api) createPolicy(c echo.Context) (err error) {
	policy := new(models.Policy)
	c.Bind(policy)
	for _, permission := range policy.Permissions {
		if err = c.Validate(permission); err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
	}
	if err = a.DB.CreatePolicy(policy); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, policy)
}

func (a *Api) getPolicies(c echo.Context) error {
	policies, err := a.DB.GetPolicies()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, policies)
}

func (a *Api) userPermissions(c echo.Context) (err error) {
	userAct := new(checkParams)
	c.Bind(userAct)
	access, err := a.check(userAct)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, access)
}
