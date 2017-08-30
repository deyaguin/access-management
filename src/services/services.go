package services

//type servicesConf struct {
//	userService             UserService
//	groupService            GroupService
//	policyService           PolicyService
//	permissionService       PermissionService
//	permissionsCheckService PermissionsCheckService
//}
//
//func NewServicesConf()

type items struct {
	Item  interface{} `json:"items"`
	Count int         `json:"count"`
}
