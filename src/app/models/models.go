package models

import (
	//"github.com/jinzhu/gorm"
)

type entity interface {
	Equals(e *entity) bool
}

type User struct {
	ID uint `json:"id"`
	Name string
	Groups []Group `gorm:"many2many:user_groups"json:"-"`
	Policies []Policy `gorm:"many2many:user_policies"json:"-"`
}

func (u *User) Equals(comparedU *User) bool {
	if u.Name == comparedU.Name && u.ID == comparedU.ID {
		return true
	}
	return false
}

type Group struct {
	ID uint `json:"id"`
	Name string
	Policies []Policy `gorm:"many2many:group_policies"json:"-"`
}

type Policy struct {
	ID uint `json:"id"`
	Name string
	Groups []Group `gorm:"many2many:group_policies"json:"-"`
	Users []User `gorm:"many2many:user_policies"json:"-"`
	Permissions []Permission `json:"-"`
}

type Permission struct {
	ID uint `json:"id"`
	Resourse string
	AccessType string
	ActionID uint `gorm:"column:action_id"`
	PolicyID uint `gorm:"column:policy_id"`
}

func (p *Permission) Equals(comparedP Permission) bool {
	if p.Resourse == comparedP.Resourse &&
		p.AccessType == comparedP.AccessType &&
		p.ActionID == comparedP.ActionID {
		return true
	}
	return false
}

type Action struct {
	ID uint `json:"id"`
	Name string
	Permissions []Permission `json:"-"`
}

type Service struct {
	ID uint `json:"id"`
	Name string
	Actions []Action `json:"-"`
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
