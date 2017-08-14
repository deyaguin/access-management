package api

import (
	"github.com/labstack/echo"
	"app/models"
	"net/http"
)

type userAction struct {
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
	return c.JSON(http.StatusOK, user)
}

func (a *Api) getUsers(c echo.Context) error {
	users, dbError := a.DB.GetUsers()
	if dbError != nil {
		return dbError
	}
	return c.JSON(http.StatusOK, users)
}

func (a *Api) deleteUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	dbError := a.DB.DeleteUser(user)
	if dbError != nil {
		return dbError
	}
	return c.JSON(http.StatusOK, user)
}

func (a *Api) createGroup(c echo.Context) error {
	group := new(models.Group)
	c.Bind(group)
	return c.JSON(http.StatusOK, group)
}

func (a *Api) getGroups(c echo.Context) error {
	groups, dbError := a.DB.GetGroups()
	if dbError != nil {
		return c.JSON(http.StatusOK, groups)
	}
	return dbError
}

func (a *Api) createPolicy(c echo.Context) error {
	policy := new(models.Policy)
	c.Bind(policy)
	return c.JSON(http.StatusOK, policy)
}

func (a *Api) getPolicies(c echo.Context) error {
	policies, err :=	a.DB.GetPolicies()
	if err != nil {

	}
	return c.JSON(http.StatusOK, policies)
}

func (a *Api) userPermissions(c echo.Context) error {
	userAct := new(userAction)
	c.Bind(userAct)
	access := a.check(userAct)
	return c.JSON(http.StatusOK, access)
}
