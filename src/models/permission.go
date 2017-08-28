package models

import (
	"time"
)

type Permission struct {
	ID        int        `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" gorm:"default:''"`
	Resourse  string     `json:"resourse" validate:"nonzero"`
	Access    bool       `json:"access" validate:"nonzero"`
	ActionID  int        `json:"actionId" gorm:"column:action_id" validate:"nonzero"`
	PolicyID  int        `gorm:"column:policy_id" json:"-"`
}

func (p *Permission) Equals(permission Permission) bool {
	result := p.Resourse == permission.Resourse &&
		p.Access == permission.Access &&
		p.ActionID == permission.ActionID
	return result
}

func (p *Permission) SetFields(permission *Permission) {
	p.Resourse = permission.Resourse
	p.Access = permission.Access
	p.ActionID = permission.ActionID
}
