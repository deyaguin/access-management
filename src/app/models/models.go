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
	Group_id int
	Policy_id int
}

type UserPolicy struct {
	User_id int
	Policy_id int
}

type UserGroup struct {
	User_id int
	Group_id int
}
