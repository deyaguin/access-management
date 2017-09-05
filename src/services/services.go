package services

//type servicesConf struct {
//	usersService             UsersService
//	groupsService            GroupsService
//	policiesService           PoliciesService
//	permissionsService       PermissionsService
//	permissionsCheckService PermissionsCheckService
//}
//
//func NewServicesConf()

type items struct {
	Item  interface{} `json:"items"`
	Count int         `json:"count"`
}
