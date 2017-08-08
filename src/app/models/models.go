package models

import (
	//"github.com/jinzhu/gorm"
)

type User struct {
	ID int
	Name string
}

type Group struct {
	ID int
	Name string
}

type Policy struct {
	ID int
	Name string
}

type GroupPolicy struct {
	GroupId int `gorm:"column:group_id;"`
	PolicyId int `gorm:"column:policy_id"`
}

type UserPolicy struct {
	UserId int `gorm:"column:user_id"`
	PolicyId int `gorm:"column:policy_id"`
}

type UserGroup struct {
	UserId int `gorm:"column:user_id"`
	GroupId int `gorm:"column:groups_id"`
}
