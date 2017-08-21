package models

type Permission struct {
	ID       int    `json:"id"`
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
