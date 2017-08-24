package models

import (
	"time"
)

type Group struct {
	ID        int        `json:"id" gorm:"AUTO_INCREMENT;not null;unique"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt time.Time `json:"-"`
	Name      string
	Users     []User   `gorm:"many2many:user_groups;save_associations:false" json:"-"`
	Policies  []Policy `gorm:"many2many:group_policies;save_associations:false" json:"-"`
}

func (g *Group) Equals(comparedG *Group) bool {
	result := g.Name == comparedG.Name
	return result
}
