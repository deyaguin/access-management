package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID       int      `json:"id" gorm:"AUTO_INCREMENT;not null;unique"`
	Name     string   `validate:"required"`
	Groups   []Group  `gorm:"many2many:user_groups;save_associations:false"`
	Policies []Policy `gorm:"many2many:user_policies;save_associations:false"`
}

func (u *User) Equals(comparedU *User) bool {
	result := u.Name == comparedU.Name
	return result
}
