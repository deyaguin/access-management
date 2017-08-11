package models

import (
	//"github.com/jinzhu/gorm"
)

type User struct {
	ID uint `json:"id"`
	Name string
	Groups []Group `gorm:"many2many:user_groups"json:"-"`
	Policies []Policy `gorm:"many2many:user_policies"json:"-"`
}

func (u *User) Equals(comparedU *User) bool {
	result := u.Name == comparedU.Name
	return result
}

type Group struct {
	ID uint `json:"id"`
	Name string
	Policies []Policy `gorm:"many2many:group_policies"json:"-"`
}

func (g *Group) Equals(comparedG *Group) bool {
	result := g.Name == comparedG.Name
	return result
}

type Policy struct {
	ID uint `json:"id"`
	Name string
	Groups []Group `gorm:"many2many:group_policies"json:"-"`
	Users []User `gorm:"many2many:user_policies"json:"-"`
	Permissions []Permission `json:"-"`
}

func (p *Policy) Equals(comparedP *Policy) bool {
	result := p.Name == comparedP.Name
	return result
}

type Permission struct {
	ID       uint `json:"id"`
	Resourse string
	Access   bool
	ActionID uint `gorm:"column:action_id"`
	PolicyID uint `gorm:"column:policy_id"`
}

func (p *Permission) Equals(comparedP Permission) bool {
	result := p.Resourse == comparedP.Resourse &&
		p.Access == comparedP.Access &&
		p.ActionID == comparedP.ActionID
	return result
}

type Action struct {
	ID uint `json:"id"`
	Name string
	Permissions []Permission `json:"-"`
}

func (a *Action) Equals(comparedA Action) bool {
	result := a.Name == comparedA.Name
	return result
}

type Service struct {
	ID uint `json:"id"`
	Name string
	Actions []Action `json:"-"`
}

func (s *Service) Equals(comparedS Service) bool {
	result := s.Name == comparedS.Name
	return result
}

//type UserGroup struct {
//	UserID uint `gorm:"column:user_id"`
//	GroupID uint `gorm:"column:group_id"`
//}
//
//type UserPolicy struct {
//	UserID uint `gorm:"column:user_id"`
//	PolicyID uint `gorm:"column:policy_id"`
//}
//
//type GroupPolicy struct {
//	GroupID uint `gorm:"column:group_id"`
//	PolicyID uint `gorm:"column:policy_id"`
//}
