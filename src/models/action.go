package models

import "time"

type Action struct {
	ID          int `json:"id" gorm:"AUTO_INCREMENT;not null;unique"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time   `json:""`
	Name        string       `json:"name" validate:"nonzero"`
	Permissions []Permission `json:"-"`
}

func (a *Action) Equals(comparedA Action) bool {
	result := a.Name == comparedA.Name
	return result
}
