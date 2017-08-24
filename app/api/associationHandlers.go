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

func (a *Api) addUsersToGroup(c echo.Context) error {
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
	return c.NoContent(http.StatusOK)
}

func (a *Api) removeUserFromGroup(c echo.Context) error {
	gid, err := strconv.Atoi(c.Param("gid"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetGroup(gid); err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetUser(uid); err != nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	group := &models.Group{ID: gid}
	user := &models.User{ID: uid}
	if err = a.DB.RemoveUserFromGroup(group, user); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (a *Api) addPermissionsToPolicy(c echo.Context) error {
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
	return c.NoContent(http.StatusOK)
}

func (a *Api) attachPoliciesByUser(c echo.Context) error {
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
		}
	}
	if err = a.DB.AttachPoliciesByUser(user, &policies.Policies); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (a *Api) detachPolicyByUser(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetUser(uid); err != nil {
		return c.JSON(http.StatusNotFound, "user not found")
	}
	pid, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetPolicy(pid); err != nil {
		return c.JSON(http.StatusNotFound, "policy not found")
	}
	user := &models.User{ID: uid}
	policy := &models.Policy{ID: pid}
	if err = a.DB.DetachPolicyByUser(user, policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (a *Api) attachPoliciesByGroup(c echo.Context) error {
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
		}
	}
	if err = a.DB.AttachPoliciesByGroup(group, &policies.Policies); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (a *Api) detachPolicyByGroup(c echo.Context) error {
	gid, err := strconv.Atoi(c.Param("gid"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetGroup(gid); err != nil {
		return c.JSON(http.StatusNotFound, "group not found")
	}
	pid, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	if _, err = a.DB.GetPolicy(pid); err != nil {
		return c.JSON(http.StatusNotFound, "policy not found")
	}
	group := &models.Group{ID: gid}
	policy := &models.Policy{ID: pid}
	if err = a.DB.DetachPolicyByGroup(group, policy); err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.NoContent(http.StatusOK)
}
