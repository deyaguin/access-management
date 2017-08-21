package models

type Group struct {
	ID       int      `json:"id" gorm:"AUTO_INCREMENT;not null;unique"`
	Name     string   `validate:"required"`
	Users    []User   `gorm:"many2many:user_groups;save_associations:false"`
	Policies []Policy `gorm:"many2many:group_policies;save_associations:false"`
}

func (g *Group) Equals(comparedG *Group) bool {
	result := g.Name == comparedG.Name
	return result
}
