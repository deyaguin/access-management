package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/accessControl/app/models"
	"net/http"
	"strconv"
)

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
	c.Bind(group)
	users := group.Users
	for _, u := range users {
		if _, err = a.DB.GetUser(u.ID); err != nil {
			return c.JSON(http.StatusNotFound, "user not found")
			break
		}
	}
	if err = a.DB.AddUsersToGroup(group, &users); err != nil {
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
	c.Bind(group)
	users := group.Users
	for _, u := range users {
		if _, err = a.DB.GetUser(u.ID); err != nil {
			return c.JSON(http.StatusNotFound, "user not found")
			break
		}
	}
	if err = a.DB.RemoveUsersFromGroup(group, &users); err != nil {
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
	c.Bind(policy)
	permissions := policy.Permissions
	for p := range permissions {
		if err = c.Validate(p); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}
	if err = a.DB.AddPermissionsToPolicy(policy, &permissions); err != nil {
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
	c.Bind(policy)
	permissions := policy.Permissions
	for _, p := range permissions {
		if _, err = a.DB.GetPermission(p.ID); err != nil {
			return c.JSON(http.StatusNotFound, "permission not found")
			break
		}
	}
	if err = a.DB.RemovePermissionsFromPolicy(policy, &permissions); err != nil {
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
	c.Bind(user)
	policies := user.Policies
	for _, p := range policies {
		if _, err = a.DB.GetPolicy(p.ID); err != nil {
			return c.JSON(http.StatusNotFound, "policy not found")
			break
		}
	}
	if err = a.DB.AttachPoliciesByUser(user, &policies); err != nil {
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
	c.Bind(user)
	policies := user.Policies
	for _, p := range policies {
		if _, err = a.DB.GetPolicy(p.ID); err != nil {
			return c.JSON(http.StatusNotFound, "policy not found")
			break
		}
	}
	if err = a.DB.DetachPoliciesByUser(user, &policies); err != nil {
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
	c.Bind(group)
	policies := group.Policies
	for _, p := range policies {
		if _, err = a.DB.GetPolicy(p.ID); err != nil {
			return c.JSON(http.StatusNotFound, "policy not found")
			break
		}
	}
	if err = a.DB.AttachPoliciesByGroup(group, &policies); err != nil {
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
	c.Bind(group)
	policies := group.Policies
	for _, p := range policies {
		if _, err = a.DB.GetPolicy(p.ID); err != nil {
			return c.JSON(http.StatusNotFound, "policy not found")
			break
		}
	}
	if err = a.DB.DetachPoliciesByGroup(group, &policies); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, "detached")
}
