package models

type Action struct {
	ID          int `json:"id"`
	Name        string
	Permissions []Permission `json:"-"`
}

func (a *Action) Equals(comparedA Action) bool {
	result := a.Name == comparedA.Name
	return result
}
