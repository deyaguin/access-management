package models

import "time"

type Action struct {
	ID          int          `json:"id" gorm:"AUTO_INCREMENT;not null;unique"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   *time.Time   `json:"-"`
	Name        string       `json:"name" validate:"nonzero"`
	Permissions []Permission `json:"-"`
}

func (a *Action) Equals(action *Action) bool {
	result := a.Name == action.Name
	return result
}

func (a *Action) SetFields(action *Action) {
	a.Name = action.Name
}
