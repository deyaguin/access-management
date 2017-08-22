package models

type Action struct {
	ID          int `json:"id"`
	UUID        string
	Name        string `validate:"required"`
	Permissions []Permission
}

func (a *Action) Equals(comparedA Action) bool {
	result := a.Name == comparedA.Name
	return result
}
