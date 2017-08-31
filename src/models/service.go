package models

import "time"

type Service struct {
	ID        int        `json:"id" gorm:"AUTO_INCREMENT;not null;unique"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-"`
	Name      *string    `json:"name" validate:"nonzero"`
	Actions   []Action   `json:"-"`
}

func (s *Service) Equals(service Service) bool {
	result := s.Name == service.Name
	return result
}

func (s *Service) SetFields(service Service) {
	if service.Name != nil {
		s.Name = service.Name
	}
}
