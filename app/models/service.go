package models

type Service struct {
	ID      int   `json:"id"`
	Name    string `validate:"required"`
	Actions []Action
}

func (s *Service) Equals(comparedS Service) bool {
	result := s.Name == comparedS.Name
	return result
}
