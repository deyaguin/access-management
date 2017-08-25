package api

//
//import (
//	"github.com/labstack/echo"
//	"gitlab/nefco/accessControl/app/models"
//	"net/http"
//	"strconv"
//)
//
//func (a *Api) getUsersByGroupHandler(c echo.Context) error {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		c.Logger().Error(err)
//		return err
//	}
//	if _, err = a.DB.GetGroup(id); err != nil {
//		return c.JSON(http.StatusNotFound, "group not found")
//	}
//	group := &models.Group{ID: id}
//	users, err := a.getUsersByGroup(group)
//	if err != nil {
//		c.Logger().Error(err)
//		return err
//	}
//	return c.JSON(http.StatusOK, users)
//}
//
//func (a *Api) getPermissionsByPolicyHandler(c echo.Context) error {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		c.Logger().Error(err)
//		return err
//	}
//	if _, err = a.DB.GetPolicy(id); err != nil {
//		return c.JSON(http.StatusNotFound, "policy not found")
//	}
//	policy := &models.Policy{ID: id}
//	permissions, err := a.getPermissionsByPolicy(policy)
//	if err != nil {
//		c.Logger().Error(err)
//		return err
//	}
//	return c.JSON(http.StatusOK, permissions)
//}
//
//func (a *Api) getPoliciesByUserHandler(c echo.Context) error {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		c.Logger().Error(err)
//		return err
//	}
//	if _, err = a.DB.GetUser(id); err != nil {
//		return c.JSON(http.StatusNotFound, "user not found")
//	}
//	user := &models.User{ID: id}
//	policies, err := a.getPoliciesByUser(user)
//	if err != nil {
//		c.Logger().Error(err)
//		return err
//	}
//	return c.JSON(http.StatusOK, policies)
//}
//
//func (a *Api) getPoliciesByGroupHandler(c echo.Context) error {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		c.Logger().Error(err)
//		return err
//	}
//	if _, err = a.DB.GetGroup(id); err != nil {
//		return c.JSON(http.StatusNotFound, "group not found")
//	}
//	group := &models.Group{ID: id}
//	policies, err := a.getPoliciesByGroup(group)
//	if err != nil {
//		c.Logger().Error(err)
//		return err
//	}
//	return c.JSON(http.StatusOK, policies)
//}
