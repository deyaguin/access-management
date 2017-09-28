package models

import "time"

type Action struct {
	ID          int          `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   *time.Time   `json:"-"`
	Name        *string      `json:"name" validate:"nonzero"`
	ServiceID   int          `json:"serviceId"`
	Permissions []Permission `json:"-"`
}

func (a *Action) Equals(action *Action) bool {
	result := a.Name == action.Name &&
		a.ServiceID == action.ServiceID
	return result
}

func (a *Action) SetFields(action *Action) {
	if action.Name != nil {
		a.Name = action.Name
	}
	if action.ServiceID != 0 {
		a.ServiceID = action.ServiceID
	}
}
