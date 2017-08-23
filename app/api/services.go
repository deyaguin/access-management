package api

import (
	"gitlab/nefco/accessControl/app/models"
)

func (a *Api) getPoliciesByUser(u *models.User) (*[]models.Policy, error) {
	p := new([]models.Policy)
	e := a.DB.GetPoliciesByUser(u, p)
	return p, e
}

func (a *Api) getPoliciesByGroup(g *models.Group) (*[]models.Policy, error) {
	p := new([]models.Policy)
	e := a.DB.GetPoliciesByGroup(g, p)
	return p, e
}

func (a *Api) getGroupsByUser(u *models.User) (*[]models.Group, error) {
	g := new([]models.Group)
	e := a.DB.GetGroupsByUser(u, g)
	return g, e
}

func (a *Api) getUsersByGroup(g *models.Group) (*[]models.User, error) {
	u := new([]models.User)
	e := a.DB.GetUsersByGroup(g, u)
	return u, e
}

func (a *Api) getPermissionsByPolicy(pol *models.Policy) (*[]models.Permission, error) {
	per := new([]models.Permission)
	e := a.DB.GetPermissionsByPolicy(pol, per)
	return per, e
}

func (a *Api) checkPermissions(userAct *checkParams) (bool, error) {
	user := new(models.User)
	user.ID = userAct.ID
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
	policies, err := a.getPoliciesByUser(u)
	if err == nil {
		for _, policy := range *policies {
			p, e := a.getPermissionsByPolicy(&policy)
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
	groups, err := a.getGroupsByUser(u)
	for _, group := range *groups {
		p, e := a.getPoliciesByGroup(&group)
		policies = append(policies, *p...)
		if e != nil {
			err = e
			break
		}
	}
	for _, policy := range policies {
		p, e := a.getPermissionsByPolicy(&policy)
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
