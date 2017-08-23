package models

type Service struct {
	ID      int `json:"id"`
	Name    string
	Actions []Action `json:"-"`
}

func (s *Service) Equals(comparedS Service) bool {
	result := s.Name == comparedS.Name
	return result
}
