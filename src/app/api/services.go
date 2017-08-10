package api

import (
	"app/models"
	"fmt"
)

func (a *Api) GetUserPolicies(user *models.User) *[]models.Policy {
	policies := new([]models.Policy)
	a.DB.GetEntityAssociations(user, policies, "Policies" )
	return policies
}

func (a *Api) GetGroupPolicies(group *models.Group) *[]models.Policy {
	policies := new([]models.Policy)
	a.DB.GetEntityAssociations(group, policies, "Policies")
	return policies
}

func (a *Api) GetUserGroups(user *models.User) *[]models.Group {
	groups := new([]models.Group)
	a.DB.GetEntityAssociations(user, groups, "Groups")
	return groups
}

func (a *Api) GetPolicyPermissions(policy *models.Policy) *[]models.Permission {
	permissions := new([]models.Permission)
	a.DB.GetEntityAssociations(policy, permissions, "Permissions")
	return permissions
}

func (a *Api) FormListOfPermissions(user *models.User) []models.Permission {
	self := a.formListOfSelfPermissions(user)
	group := a.formListOfGroupPermissions(user)
	//p := a.compareByAccessType(self)
	p := a.compareByAccessType(self)
	//fmt.Println(self)
	fmt.Println(p)
	return append(self, group...)
}

func (a *Api) formListOfSelfPermissions(user *models.User) []models.Permission {
	permissions := new([]models.Permission)
	a.GetUserPolicies(user)
	policies := a.GetUserPolicies(user)
	for _, policy := range *policies {
		p := a.GetPolicyPermissions(&policy)
		*permissions = append(*permissions, *p...)
	}
	return *permissions
}

func (a *Api) formListOfGroupPermissions(user *models.User) []models.Permission {
	permissions := new([]models.Permission)
	policies := new([]models.Policy)
	a.GetUserPolicies(user)
	groups := a.GetUserGroups(user)
	for _, group := range *groups {
		p := a.GetGroupPolicies(&group)
		*policies = append(*policies, *p...)
	}
	for _, policy := range *policies {
		p := a.GetPolicyPermissions(&policy)
		*permissions = append(*permissions, *p...)
	}
	return *permissions
}

//func (a *Api) compareSelfAndGroupPermissions(self, group []models.Permission) []models.Permission {
//	result := new([]models.Permission)
//	for _, s := range self {
//		equal := false
//		for _, g := range group {
//			if s.Equals(g) {
//				equal = true
//			}
//		}
//		if
//	}
//}

func (a *Api) compareByAccessType(permissions []models.Permission) []models.Permission {
	result := new([]models.Permission)
	for i, p := range permissions {
		for _, cP := range permissions[i+1:] {
			if p.Resourse == cP.Resourse &&
				p.ActionID == cP.ActionID &&
				p.AccessType != cP.AccessType {
			} else {
				*result = append(*result, p)
			}
		}
	}
	return *result
}

//func (a *Api) onlyUn(permissions []models.Permission) []models.Permission {
//
//}
