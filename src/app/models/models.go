package models

import (
	//"github.com/jinzhu/gorm"
)

type User struct {
	ID uint `json:"-"`
	Name string
	Groups []Group `gorm:"many2many:user_groups"json:"-"`
	Policies []Policy `gorm:"many2many:user_policies"json:"-"`
}

type Group struct {
	ID uint `json:"-"`
	Name string
	Policies []Policy `gorm:"many2many:group_policies"json:"-"`
}

type Policy struct {
	ID uint `json:"-"`
	Name string
	Groups []Group `gorm:"many2many:group_policies"json:"-"`
	Users []User `gorm:"many2many:user_policies"json:"-"`
	Permissions []Permission `json:"-"`
}

type Permission struct {
	ID uint `json:"-"`
	Resourse string
	AccessType string
	ActionID uint `gorm:"column:action_id"`
	PolicyID uint `gorm:"column:policy_id"`
}

type Action struct {
	ID uint `json:"-"`
	Name string
	Permissions []Permission `json:"-"`
}

type Service struct {
	ID uint `json:"-"`
	Name string
	Actions []Action `json:"-"`
}

type UserGroup struct {
	UserID uint `gorm:"column:user_id"`
	GroupID uint `gorm:"column:group_id"`
}

type UserPolicy struct {
	UserID uint `gorm:"column:user_id"`
	PolicyID uint `gorm:"column:policy_id"`
}

type GroupPolicy struct {
	GroupID uint `gorm:"column:group_id"`
	PolicyID uint `gorm:"column:policy_id"`
}
