package models

import (
	"github.com/jinzhu/gorm"
)

//type dao interface {
//	createEntity(*interface{})
//
//}

type UserDAO struct {
	db *gorm.DB
}

func (ud *UserDAO) Create(user *User) {
	ud.db.NewRecord(user)
}
