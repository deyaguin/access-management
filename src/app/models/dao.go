package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type UserDAO struct {
	DB *gorm.DB
}

func (ud *UserDAO) Create(user *User) {
	ud.DB.NewRecord(&user)
	ud.DB.Create(&user)
}

func (ud *UserDAO) Get(users *[]User) {
	ud.DB.Find(&users)
	fmt.Println(users)
}

type GroupDAO struct {
	DB *gorm.DB
}