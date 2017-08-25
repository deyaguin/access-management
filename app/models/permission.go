package models

import (
	"time"
)

type Permission struct {
	ID        int        `json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" gorm:"default:''"`
	Resourse  string     `validate:"required"`
	Access    bool       `validate:"required"`
	ActionID  int        `gorm:"column:action_id" validate:"required"`
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
