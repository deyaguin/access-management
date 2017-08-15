package api

import (
	"app/models"
)

func (a *Api) getPoliciesByUser(u *models.User) (*[]models.Policy, error) {
	p := new([]models.Policy)
	e := a.DB.GetPoliciesByUser(u, p, "Policies" )
	return p, e
}

func (a *Api) getPoliciesByGroup(g *models.Group) (*[]models.Policy, error) {
	p := new([]models.Policy)
	e := a.DB.GetPoliciesByGroup(g, p, "Policies")
	return p, e
}

func (a *Api) getGroupsByUser(u *models.User) (*[]models.Group, error) {
	g := new([]models.Group)
	e := a.DB.GetGroupsByUser(u, g, "Groups")
	return g, e
}

func (a *Api) getPermissionsByPolicy(pol *models.Policy) (*[]models.Permission, error) {
	per := new([]models.Permission)
	e := a.DB.GetPermissionsByPolicy(pol, per, "Permissions")
	return per, e
}

func (a *Api) check(userAct *checkParams) (bool, error) {
	user := new(models.User)
	user.Name = userAct.Name
	user.ID = userAct.ID
	self, err := a.getSelfPermissions(user)
	if err != nil {
		return false, err
	}
	group, err := a.getGroupPermissions(user)
	access, has := a.checkPermissions(self, userAct)
	if !has {
		access, _ = a.checkPermissions(group, userAct)
	}
	return access, err
}

func (a *Api) getSelfPermissions(u *models.User) ([]models.Permission, error) {
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
		policies []models.Policy
		permissions []models.Permission
	)
	groups, err := a.getGroupsByUser(u)
	for _, group := range *groups {
		p, e := a.getPoliciesByGroup(&group)
		policies = append(policies, *p...)
		if e != nil {
			err = e
		}
	}
	for _, policy := range policies {
		p, e := a.getPermissionsByPolicy(&policy)
		permissions = append(permissions, *p...)
		if e != nil {
			err = e
		}
	}
	return permissions, err
}

func (a *Api) checkPermissions(p []models.Permission, userAct *checkParams) (result bool, has bool) {
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
