package models

import (
	"time"
)

type Group struct {
	ID        int        `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
	Name      *string    `json:"name" validate:"nonzero"`
	Users     []User     `gorm:"many2many:user_groups;" json:"-"`
	Policies  []Policy   `gorm:"many2many:group_policies;save_associations:false" json:"-"`
}

func (g *Group) Equals(group *Group) bool {
	result := g.Name == group.Name
	return result
}

func (g *Group) SetFields(group *Group) {
	if group.Name != nil {
		g.Name = group.Name
	}
}
