package models

type Policy struct {
	ID          int    `json:"id"`
	Name        string  `validate:"required"`
	Groups      []Group `gorm:"many2many:group_policies;save_associations:false"`
	Users       []User  `gorm:"many2many:user_policies;save_associations:false"`
	Permissions []Permission
}

func (p *Policy) Equals(comparedP *Policy) bool {
	result := p.Name == comparedP.Name
	return result
}
