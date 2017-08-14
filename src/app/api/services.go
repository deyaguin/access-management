package api

import (
	"app/models"
)

func (a *Api) getPoliciesByUser(u *models.User) *[]models.Policy {
	policies := new([]models.Policy)
	err := a.DB.GetPoliciesByUser(u, policies, "Policies" )
	if err != nil {

	}
	return policies
}

func (a *Api) getPoliciesByGroup(g *models.Group) *[]models.Policy {
	policies := new([]models.Policy)
	err := a.DB.GetPoliciesByGroup(g, policies, "Policies")
	if err != nil {

	}
	return policies
}

func (a *Api) getGroupsByUser(u *models.User) *[]models.Group {
	groups := new([]models.Group)
	err := a.DB.GetGroupsByUser(u, groups, "Groups")
	if err != nil {

	}
	return groups
}

func (a *Api) getPermissionsByPolicy(p *models.Policy) *[]models.Permission {
	permissions := new([]models.Permission)
	err := a.DB.GetPermissionsByPolicy(p, permissions, "Permissions")
	if err != nil {

	}
	return permissions
}

func (a *Api) check(userAct *userAction) bool {
	user := new(models.User)
	user.Name = userAct.Name
	user.ID = userAct.ID
	self := a.getSelfPermissions(user)
	group := a.getGroupPermissions(user)
	access, has := a.checkPermissions(self, userAct)
	if !has {
		access, _ = a.checkPermissions(group, userAct)
	}
	return access
}

func (a *Api) getSelfPermissions(u *models.User) []models.Permission {
	var permissions []models.Permission
	a.getPoliciesByUser(u)
	policies := a.getPoliciesByUser(u)
	for _, policy := range *policies {
		p := a.getPermissionsByPolicy(&policy)
		permissions = append(permissions, *p...)
	}
	return permissions
}

func (a *Api) getGroupPermissions(u *models.User) []models.Permission {
	var permissions []models.Permission
	var policies []models.Policy
	a.getPoliciesByUser(u)
	groups := a.getGroupsByUser(u)
	for _, group := range *groups {
		p := a.getPoliciesByGroup(&group)
		policies = append(policies, *p...)
	}
	for _, policy := range policies {
		p := a.getPermissionsByPolicy(&policy)
		permissions = append(permissions, *p...)
	}
	return permissions
}

func (a *Api) checkPermissions(p []models.Permission, userAct *userAction) (bool, bool) {
	result := false
	has := false
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
