package models

import "time"

type Service struct {
	ID        int `json:"id" gorm:"AUTO_INCREMENT;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:""`
	Name      string
	Actions   []Action `json:"-"`
}

func (s *Service) Equals(comparedS Service) bool {
	result := s.Name == comparedS.Name
	return result
}
