package models

type User struct {
	ID uint `json:"id"`
	Name string `validate:"required"`
	Groups []Group `gorm:"many2many:user_groups"json:"-"`
	Policies []Policy `gorm:"many2many:user_policies"json:"-"`
}

func (u *User) Equals(comparedU *User) bool {
	result := u.Name == comparedU.Name
	return result
}

type Group struct {
	ID uint `json:"id"`
	Name string `validate:"required"`
	Policies []Policy `gorm:"many2many:group_policies"json:"-"`
}

func (g *Group) Equals(comparedG *Group) bool {
	result := g.Name == comparedG.Name
	return result
}

type Policy struct {
	ID uint `json:"id"`
	Name string `validate:"required"`
	Groups []Group `gorm:"many2many:group_policies"json:"-"`
	Users []User `gorm:"many2many:user_policies"json:"-"`
	Permissions []Permission `json:"permissions"`
}

func (p *Policy) Equals(comparedP *Policy) bool {
	result := p.Name == comparedP.Name
	return result
}

type Permission struct {
	ID uint `json:"id"`
	Resourse string `validate:"required"`
	Access bool
	ActionID uint `gorm:"column:action_id"validate:"required"`
	PolicyID uint `gorm:"column:policy_id"`
}

func (p *Permission) Equals(comparedP Permission) bool {
	result := p.Resourse == comparedP.Resourse &&
		p.Access == comparedP.Access &&
		p.ActionID == comparedP.ActionID
	return result
}

type Action struct {
	ID uint `json:"id"`
	Name string `validate:"required"`
	Permissions []Permission `json:"-"`
}

func (a *Action) Equals(comparedA Action) bool {
	result := a.Name == comparedA.Name
	return result
}

type Service struct {
	ID uint `json:"id"`
	Name string `validate:"required"`
	Actions []Action `json:"-"`
}

func (s *Service) Equals(comparedS Service) bool {
	result := s.Name == comparedS.Name
	return result
}
