package models

import (
	"time"
)

type User struct {
	ID        int        `json:"id" gorm:"AUTO_INCREMENT;not null;unique"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
	Name      *string    `json:"name" validate:"nonzero"`
	Groups    []Group    `gorm:"many2many:user_groups;save_associations:false" json:"-"`
	Policies  []Policy   `gorm:"many2many:user_policies;save_associations:false" json:"-"`
}

func (u *User) Equals(user *User) bool {
	result := u.Name == user.Name
	return result
}

func (u *User) SetFields(user *User) {
	if user.Name != nil {
		u.Name = user.Name
	}
}
