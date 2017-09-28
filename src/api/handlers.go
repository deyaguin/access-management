package api

import (
	"gitlab/nefco/access-management-system/src/models"
	"gitlab/nefco/access-management-system/src/services"
)

type users struct {
	Users *[]models.User `validate:"required" json:"users"`
}

type policies struct {
	Policies *[]models.Policy `validate:"required" json:"policies"`
}

type permissions struct {
	Permissions *[]models.Permission `validate:"required"`
}

type groups struct {
	Groups *[]models.Group `validate:"required"`
}

func checkPaginationParams(page, perPage int) error {
	if page < 1 {
		return services.NewValidationError("page not valid")
	}

	if perPage < 1 {
		return services.NewValidationError("perPage not valid")
	}

	return nil
}
