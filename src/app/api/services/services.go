package services

import (
	"app/db"
	"app/models"
)

type Service struct {
	DB db.DB
}

func (s *Service) GetUserPolicies(user *models.User, policies *[]models.Policy) {
	s.DB.GetEntityAssociations(user, policies, "Policies" )
}
