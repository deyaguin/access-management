package services

import (
	"reflect"
	"gopkg.in/validator.v2"
)

type ServicesConf struct {
	UsersService
	GroupsService
	PoliciesService
	PermissionsService
	PermissionsCheckService
	ActionsService
	ServicesService
}

func NewServicesConf(
	users UsersService,
	groups GroupsService,
	policies PoliciesService,
	permissions PermissionsService,
	permissionsCheck PermissionsCheckService,
	actions ActionsService,
	services ServicesService,
) *ServicesConf {
	validator.SetValidationFunc("isBool", isBool)
	return &ServicesConf{
		users,
		groups,
		policies,
		permissions,
		permissionsCheck,
		actions,
		services,
	}
}

type items struct {
	Items       interface{} `json:"items"`
	Total       int         `json:"total"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
}

type pureItems struct {
	Items interface{} `json:"items"`
}

func isBool(value interface{}, param string) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Bool {
		return validator.ErrUnsupported
	}
	return nil
}


