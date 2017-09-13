package services

type ServicesConf struct {
	UsersService
	GroupsService
	PoliciesService
	PermissionsService
	PermissionsCheckService
	ActionsService
}

func NewServicesConf(
	users UsersService,
	groups GroupsService,
	policies PoliciesService,
	permissions PermissionsService,
	permissionsCheck PermissionsCheckService,
	actions ActionsService,
) *ServicesConf {
	return &ServicesConf{
		users,
		groups,
		policies,
		permissions,
		permissionsCheck,
		actions,
	}
}

type items struct {
	Items       interface{} `json:"items"`
	Total       int         `json:"total"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
}
