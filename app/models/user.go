package models

import (
	"time"
)

type User struct {
	ID        int        `json:"id" gorm:"AUTO_INCREMENT;not null;unique"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Name      string
	Groups    []Group  `gorm:"many2many:user_groups;save_associations:false" json:"-"`
	Policies  []Policy `gorm:"many2many:user_policies;save_associations:false" json:"-"`
}

func (u *User) Equals(comparedU *User) bool {
	result := u.Name == comparedU.Name
	return result
}
