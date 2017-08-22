package models

import (
	"time"
)

type Permission struct {
	ID        int `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Resourse string `validate:"required"`
	Access   bool   `validate:"required"`
	ActionID int    `gorm:"column:action_id" validate:"required"`
	PolicyID int    `gorm:"column:policy_id"`
}

func (p *Permission) Equals(comparedP Permission) bool {
	result := p.Resourse == comparedP.Resourse &&
		p.Access == comparedP.Access &&
		p.ActionID == comparedP.ActionID
	return result
}
