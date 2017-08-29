package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/services"
)

type Api struct {
	userService       services.UserService
	groupService      services.GroupService
	policyService     services.PolicyService
	permissionService services.PermissionService
	relationsService  services.RelationsService
	permissionsCheckService services.PermissionsCheckService
}

func NewAPI(
	userService services.UserService,
	groupService services.GroupService,
	policyService services.PolicyService,
	permissionService services.PermissionService,
	relationsService services.RelationsService,
	permissionsCheckService services.PermissionsCheckService,
) {
	api := &Api{
		userService,
		groupService,
		policyService,
		permissionService,
		relationsService,
permissionsCheckService,
	}
	e := echo.New()

	e.POST("/users", api.createUser)
	e.GET("/users", api.getUsers)
	e.GET("/users/:id", api.getUser)
	e.PATCH("/users/:id", api.updateUser)
	e.DELETE("/users/:id", api.removeUser)
	e.POST("/groups", api.createGroup)
	e.GET("/groups", api.getGroups)
	e.GET("/groups/:id", api.getGroup)
	e.PATCH("/groups/:id", api.updateGroup)
	e.DELETE("/groups/:id", api.removeGroup)
	e.POST("/policies", api.createPolicy)
	e.GET("/policies", api.getPolicies)
	e.GET("/policies/:id", api.getPolicy)
	e.PATCH("/policies/:id", api.updatePolicy)
	e.DELETE("/policies/:id", api.removePolicy)
	e.PATCH("/permissions/:id", api.updatePermission)

	e.PUT("/groups/:id/users", api.addUsersToGroup)
	e.GET("/groups/:id/users", api.getUsersByGroupHandler)
	e.DELETE("/groups/:groupId/users/:userId", api.removeUserFromGroup)
	e.PUT("/policies/:id/permissions", api.addPermissionsToPolicy)
	e.GET("/policies/:id/permissions", api.getPermissionsByPolicyHandler)
	e.DELETE("/policies/:policyId/permissions/:permissionId", api.removePermissionFromPolicy)
	e.PUT("/users/:id/policies", api.attachPoliciesByUser)
	e.GET("/users/:id/policies", api.getPoliciesByUserHandler)
	e.DELETE("users/:userId/policies/:policyId", api.detachPolicyByUser)
	e.PUT("/groups/:id/policies", api.attachPoliciesByGroup)
	e.GET("/groups/:id/policies", api.getPoliciesByGroupHandler)
	e.DELETE("/groups/:groupId/policies/:policyId", api.detachPolicyByGroup)

	e.HTTPErrorHandler = api.errorHandler

	e.POST("/check_permissions", api.userPermissions)
	e.Logger.Fatal(e.Start(":1535"))
}
