package api

import (
	"github.com/labstack/echo"
	"app/models"
	"net/http"
	"strconv"
)

type checkParams struct {
	Resourse string `validation:"required"`
	Action uint `validation:"required"`
	ID uint `validation:"required"`
}

///user

func (a *Api) createUser(c echo.Context) (err error) {
	user := new(models.User)
	c.Bind(user)
	if err = c.Validate(user); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.CreateUser(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	if err = a.DB.AddUserToGroups(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	if err = a.DB.AttachPoliciesByUser(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

func (a *Api) getUsers(c echo.Context) error {
	users, err := a.DB.GetUsers()
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (a *Api) getUser(c echo.Context) (err error) {
	key, err := strconv.Atoi(c.QueryParam("key"))
	user, err := a.DB.GetUser(key)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (a *Api) updateUser(c echo.Context) (err error) {
	user := new(models.User)
	c.Bind(user)
	if err = c.Validate(user); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.UpdateUser(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

func (a *Api) deleteUser(c echo.Context) (err error) {
	user := new(models.User)
	c.Bind(user)
	if err = c.Validate(user); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.DeleteUser(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

///group

func (a *Api) createGroup(c echo.Context) (err error) {
	group := new(models.Group)
	c.Bind(group)
	if err = c.Validate(group); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.CreateGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	if err = a.DB.AddUsersToGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	if err = a.DB.AttachPoliciesByGroup(group); err != nil {
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

func (a *Api) getGroup(c echo.Context) (err error) {
	key, err := strconv.Atoi(c.QueryParam("key"))
	group, err := a.DB.GetGroup(key)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, group)
}

func (a *Api) updateGroup(c echo.Context) (err error) {
	group := new(models.Group)
	c.Bind(group)
	if err = c.Validate(group); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.UpdateGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, group)
}

func (a *Api) deleteGroup(c echo.Context) (err error) {
	group := new(models.Group)
	c.Bind(group)
	if err = c.Validate(group); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.DeleteGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, group)
}

///policy

func (a *Api) createPolicy(c echo.Context) (err error) {
	policy := new(models.Policy)
	c.Bind(policy)
	for _, permission := range policy.Permissions {
		if err = c.Validate(permission); err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}
	}
	if err = a.DB.CreatePolicy(policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	if err = a.DB.AttachPolicyToUsers(policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	if err = a.DB.AttachPolicyToGroups(policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, policy)
}

func (a *Api) getPolicies(c echo.Context) error {
	policies, err := a.DB.GetPolicies()
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, policies)
}

func (a *Api) getPolicy(c echo.Context) (err error) {
	key, err := strconv.Atoi(c.QueryParam("key"))
	policy, err := a.DB.GetPolicy(key)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, policy)
}

func (a *Api) updatePolicy(c echo.Context) (err error) {
	policy := new(models.Policy)
	c.Bind(policy)
	if err = c.Validate(policy); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.UpdatePolicy(policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, policy)
}

func (a *Api) deletePolicy(c echo.Context) (err error) {
	group := new(models.Policy)
	c.Bind(group)
	if err = c.Validate(group); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.DeletePolicy(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, group)
}

///permission

func (a *Api) createPermission(c echo.Context) (err error) {
	permission := new(models.Permission)
	c.Bind(permission)
	if err = a.DB.CreatePermission(permission); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, permission)
}

func (a *Api) updatePermission(c echo.Context) (err error) {
	permission := new(models.Permission)
	c.Bind(permission)
	if err = c.Validate(permission); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.UpdatePermission(permission); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusCreated, permission)
}

func (a *Api) deletePermission(c echo.Context) (err error) {
	permission := new(models.Permission)
	c.Bind(permission)
	if err = c.Validate(permission); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	if err = a.DB.DeletePermission(permission); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, permission)
}

///associations

func (a *Api) addUserToGroups(c echo.Context) (err error) {
	user := new(models.User)
	c.Bind(user)
	if err = a.DB.AddUserToGroups(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (a *Api) addUsersToGroup(c echo.Context) (err error) {
	group := new(models.Group)
	c.Bind(group)
	if err = a.DB.AddUsersToGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, group)
}

func (a *Api) attachPoliciesByUser(c echo.Context) (err error) {
	user := new(models.User)
	c.Bind(user)
	if err = a.DB.AttachPoliciesByUser(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (a *Api) attachPoliciesByGroup(c echo.Context) (err error) {
	group := new(models.Group)
	c.Bind(group)
	if err = a.DB.AttachPoliciesByGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, group)
}

func (a *Api) attachPolicyToUsers(c echo.Context) (err error) {
	policy := new(models.Policy)
	c.Bind(policy)
	if err = a.DB.AttachPolicyToUsers(policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, policy)
}

func (a *Api) attachPolicyToGroups(c echo.Context) (err error) {
	policy := new(models.Policy)
	c.Bind(policy)
	if err = a.DB.AttachPolicyToGroups(policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, policy)
}

func (a *Api) removeUserFromGroups(c echo.Context) (err error) {
	user := new(models.User)
	c.Bind(user)
	if err = a.DB.RemoveUserFromGroups(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (a *Api) removeUsersFromGroup(c echo.Context) (err error) {
	group := new(models.Group)
	c.Bind(group)
	if err = a.DB.RemoveUsersFromGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, group)
}

func (a *Api) detachPoliciesByUser(c echo.Context) (err error) {
	user := new(models.User)
	c.Bind(user)
	if err = a.DB.DetachPoliciesByUser(user); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (a *Api) detachPoliciesByGroup(c echo.Context) (err error) {
	group := new(models.Group)
	c.Bind(group)
	if err = a.DB.DetachPoliciesByGroup(group); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, group)
}

func (a *Api) detachUsersFromPolicy(c echo.Context) (err error) {
	policy := new(models.Policy)
	c.Bind(policy)
	if err = a.DB.DetachUsersFromPolicy(policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, policy)
}

func (a *Api) detachGroupsFromPolicy(c echo.Context) (err error) {
	policy := new(models.Policy)
	c.Bind(policy)
	if err = a.DB.DetachGroupsFromPolicy(policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, policy)
}

func (a *Api) userPermissions(c echo.Context) (err error) {
	cP := new(checkParams)
	c.Bind(cP)
	access, err := a.check(cP)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, access)
}
