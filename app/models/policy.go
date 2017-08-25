package models

import (
	"time"
)

type Policy struct {
	ID          int        `json:"id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-" gorm:"default:''"`
	Name        string
	Groups      []Group      `gorm:"many2many:group_policies;save_associations:false" json:"-"`
	Users       []User       `gorm:"many2many:user_policies;save_associations:false" json:"-"`
	Permissions []Permission `json:"-"`
}

func (p *Policy) Equals(policy *Policy) bool {
	result := p.Name == policy.Name
	return result
}

func (p *Policy) SetFields(policy *Policy) {
	p.Name = policy.Name
}
