package models

import (
	"time"
)

type Policy struct {
	ID          int          `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   *time.Time   `json:"-"`
	Name        *string      `json:"name" validate:"nonzero"`
	Description *string      `json:"description" validate:"nonzero"`
	Groups      []Group      `gorm:"many2many:group_policies;" json:"-"`
	Users       []User       `gorm:"many2many:user_policies;save_associations:false" json:"-"`
	Permissions []Permission `json:"-"`
}

func (p *Policy) Equals(policy *Policy) bool {
	result := p.Name == policy.Name && p.Description == policy.Description
	return result
}

func (p *Policy) SetFields(policy *Policy) {
	if policy.Name != nil {
		p.Name = policy.Name
	}
	if policy.Description != nil {
		p.Description = policy.Description
	}
}
