package api

import "gitlab/nefco/access-management-system/src/models"

type users struct {
	Users *[]models.User `validate:"required"`
}

type policies struct {
	Policies *[]models.Policy `validate:"required"`
}

type permissions struct {
	Permissions *[]models.Permission `validate:"required"`
}

type groups struct {
	Groups *[]models.Group `validate:"required"`
}
