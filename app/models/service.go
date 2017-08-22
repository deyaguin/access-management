package models

type Service struct {
	ID      int `json:"id"`
	UUID    string
	Name    string `validate:"required"`
	Actions []Action
}

func (s *Service) Equals(comparedS Service) bool {
	result := s.Name == comparedS.Name
	return result
}
