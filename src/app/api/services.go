package api

import (
	"app/models"
	"fmt"
)

func (a *Api) GetPoliciesByUser(user *models.User) *[]models.Policy {
	policies := new([]models.Policy)
	a.DB.GetPoliciesByUser(user, policies, "Policies" )
	return policies
}

func (a *Api) GetPoliciesByGroup(group *models.Group) *[]models.Policy {
	policies := new([]models.Policy)
	a.DB.GetPoliciesByGroup(group, policies, "Policies")
	return policies
}

func (a *Api) GetGroupsByUser(user *models.User) *[]models.Group {
	groups := new([]models.Group)
	a.DB.GetGroupsByUser(user, groups, "Groups")
	return groups
}

func (a *Api) GetPermissionsByPolicy(policy *models.Policy) *[]models.Permission {
	permissions := new([]models.Permission)
	a.DB.GetPermissionsByPolicy(policy, permissions, "Permissions")
	return permissions
}

func (a *Api) Check(userAct *userAction) bool {
	user := new(models.User)
	user.Name = userAct.Name
	user.ID = userAct.ID
	self := a.getSelfPermissions(user)
	group := a.getGroupPermissions(user)
	fmt.Println(group)
	access := a.checkPermissions(self, userAct)
	return access
}

func (a *Api) getSelfPermissions(user *models.User) []models.Permission {
	var permissions []models.Permission
	a.GetPoliciesByUser(user)
	policies := a.GetPoliciesByUser(user)
	for _, policy := range *policies {
		p := a.GetPermissionsByPolicy(&policy)
		permissions = append(permissions, *p...)
	}
	return permissions
}

func (a *Api) getGroupPermissions(user *models.User) []models.Permission {
	var permissions []models.Permission
	var policies []models.Policy
	a.GetPoliciesByUser(user)
	groups := a.GetGroupsByUser(user)
	for _, group := range *groups {
		p := a.GetPoliciesByGroup(&group)
		policies = append(policies, *p...)
	}
	for _, policy := range policies {
		p := a.GetPermissionsByPolicy(&policy)
		permissions = append(permissions, *p...)
	}
	return permissions
}

func (a *Api) checkPermissions(permissions []models.Permission, userAct *userAction) bool {
	result := false
	for _, p := range permissions {
		fmt.Println(p)
		if p.ActionID == userAct.Action && p.Resourse == userAct.Resourse {
			result = p.Access
			if !result {
				break
			}
		}
	}
	return result
}
