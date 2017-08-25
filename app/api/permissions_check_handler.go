package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/accessControl/app/models"
	"net/http"
)

type checkParams struct {
	Resourse string `validation:"required"`
	Action   int    `validation:"required"`
	ID       int    `validation:"required"`
}

func (a *Api) checkPermissions(userAct *checkParams) (bool, error) {
	user := &models.User{ID: userAct.ID}
	self, err := a.getUserPermissions(user)
	if err != nil {
		return false, err
	}
	group, err := a.getGroupPermissions(user)
	access, has := a.comparePermissions(self, userAct)
	if !has {
		access, _ = a.comparePermissions(group, userAct)
	}
	return access, err
}

func (a *Api) getUserPermissions(u *models.User) ([]models.Permission, error) {
	var permissions []models.Permission
	policies, err := a.relationsService.GetPoliciesByUser(u)
	if err == nil {
		for _, policy := range *policies {
			p, e := a.relationsService.GetPermissionsByPolicy(&policy)
			permissions = append(permissions, *p...)
			if e != nil {
				err = e
			}
		}
	}
	return permissions, err
}

func (a *Api) getGroupPermissions(u *models.User) ([]models.Permission, error) {
	var (
		policies    []models.Policy
		permissions []models.Permission
	)
	groups, err := a.relationsService.GetGroupsByUser(u)
	for _, group := range *groups {
		p, e := a.relationsService.GetPoliciesByGroup(&group)
		policies = append(policies, *p...)
		if e != nil {
			err = e
			break
		}
	}
	for _, policy := range policies {
		p, e := a.relationsService.GetPermissionsByPolicy(&policy)
		permissions = append(permissions, *p...)
		if e != nil {
			err = e
			break
		}
	}
	return permissions, err
}

func (a *Api) comparePermissions(p []models.Permission, userAct *checkParams) (result bool, has bool) {
	result = false
	has = false
	for _, p := range p {
		if p.ActionID == userAct.Action && p.Resourse == userAct.Resourse {
			has = true
			result = p.Access
			if !result {
				break
			}
		}
	}
	return result, has
}

func (a *Api) userPermissions(c echo.Context) error {
	cParams := new(checkParams)
	c.Bind(cParams)
	access, err := a.checkPermissions(cParams)
	if err != nil {
		c.Logger().Error(err)
		return err
	}
	return c.JSON(http.StatusOK, access)
}
