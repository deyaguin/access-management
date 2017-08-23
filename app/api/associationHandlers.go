package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/accessControl/app/models"
	"net/http"
	"strconv"
)

type users struct {
	Users []models.User `validate:"required"`
}
type groups struct {
	Groups []models.Group `validate:"required"`
}
type policies struct {
	Policies []models.Policy `validate:"required"`
}
type permissions struct {
	Permissions []models.Permission `validate:"required"`
}

func (a *Api) addUsersToGroup(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetGroup(id); err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	group := &models.Group{ID: id}
	users := new(users)
	c.Bind(users)
	if err = c.Validate(users); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	for _, u := range users.Users {
		if _, err = a.DB.GetUser(u.ID); err != nil {
			return c.JSON(http.StatusNotFound, "user not found")
		}
	}
	if err = a.DB.AddUsersToGroup(group, &users.Users); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "added")
}

func (a *Api) removeUsersFromGroup(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetGroup(id); err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	group := &models.Group{ID: id}
	users := new(users)
	c.Bind(users)
	for _, u := range users.Users {
		if _, err = a.DB.GetUser(u.ID); err != nil {
			return c.JSON(http.StatusNotFound, "user not found")
		}
	}
	if err = a.DB.RemoveUsersFromGroup(group, &users.Users); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "removed")
}

func (a *Api) addPermissionsToPolicy(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetPolicy(id); err != nil {
		return c.JSON(http.StatusNotFound, "policy not found")
	}
	policy := &models.Policy{ID: id}
	permissions := new(permissions)
	c.Bind(permissions)
	if err = c.Validate(permissions); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	for _, p := range permissions.Permissions {
		if err = c.Validate(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}
	if err = a.DB.AddPermissionsToPolicy(policy, &permissions.Permissions); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "added")
}

func (a *Api) removePermissionsFromPolicy(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetPolicy(id); err != nil {
		return c.JSON(http.StatusNotFound, "policy not found")
	}
	policy := &models.Policy{ID: id}
	permissions := new(permissions)
	c.Bind(permissions)
	for _, p := range permissions.Permissions {
		if _, err = a.DB.GetPermission(p.ID); err != nil {
			return c.JSON(http.StatusNotFound, "permission not found")
			break
		}
	}
	if err = a.DB.RemovePermissionsFromPolicy(policy, &permissions.Permissions); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "removed")
}

func (a *Api) attachPoliciesByUser(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetUser(id); err != nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	user := &models.User{ID: id}
	policies := new(policies)
	c.Bind(policies)
	for _, p := range policies.Policies {
		if _, err = a.DB.GetPolicy(p.ID); err != nil {
			return c.JSON(http.StatusNotFound, "policy not found")
			break
		}
	}
	if err = a.DB.AttachPoliciesByUser(user, &policies.Policies); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "attached")
}

func (a *Api) detachPoliciesByUser(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetUser(id); err != nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	user := &models.User{ID: id}
	policies := new(policies)
	c.Bind(policies)
	for _, p := range policies.Policies {
		if _, err = a.DB.GetPolicy(p.ID); err != nil {
			return c.JSON(http.StatusNotFound, "policy not found")
			break
		}
	}
	if err = a.DB.DetachPoliciesByUser(user, &policies.Policies); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "detached")
}

func (a *Api) attachPoliciesByGroup(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetGroup(id); err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	group := &models.Group{ID: id}
	policies := new(policies)
	c.Bind(policies)
	for _, p := range policies.Policies {
		if _, err = a.DB.GetPolicy(p.ID); err != nil {
			return c.JSON(http.StatusNotFound, "policy not found")
			break
		}
	}
	if err = a.DB.AttachPoliciesByGroup(group, &policies.Policies); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "attached")
}

func (a *Api) detachPoliciesByGroup(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetGroup(id); err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	group := &models.Group{ID: id}
	policies := new(policies)
	c.Bind(policies)
	for _, p := range policies.Policies {
		if _, err = a.DB.GetPolicy(p.ID); err != nil {
			return c.JSON(http.StatusNotFound, "policy not found")
			break
		}
	}
	if err = a.DB.DetachPoliciesByGroup(group, &policies.Policies); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "detached")
}
