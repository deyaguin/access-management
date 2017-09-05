package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/access-management-system/src/services"
	"gitlab/nefco/access-management-system/src/logger"
	"go.uber.org/zap"
)

var Log *zap.Logger = logger.NewLogger()

type Api struct {
	userService             services.UserService
	groupService            services.GroupService
	policyService           services.PolicyService
	permissionService       services.PermissionService
	permissionsCheckService services.PermissionsCheckService
	address string
}

func NewAPI(
	userService services.UserService,
	groupService services.GroupService,
	policyService services.PolicyService,
	permissionService services.PermissionService,
	permissionsCheckService services.PermissionsCheckService,
	address string,
) {
	api := &Api{
		userService,
		groupService,
		policyService,
		permissionService,
		permissionsCheckService,
		address,
	}
	e := echo.New()

	e.POST("/users", api.createUser)
	e.GET("/users", api.getUsers)
	e.GET("/users/:id", api.getUser)
	e.PATCH("/users/:id", api.updateUser)
	e.DELETE("/users/:id", api.removeUser)
	e.POST("/groups", api.createGroup)
	e.GET("/groups", api.getGroups)
	e.GET("/groups/:groupId", api.getGroup)
	e.PATCH("/groups/:groupId", api.updateGroup)
	e.DELETE("/groups/:groupId", api.removeGroup)
	e.POST("/policies", api.createPolicy)
	e.GET("/policies", api.getPolicies)
	e.GET("/policies/:id", api.getPolicy)
	e.PATCH("/policies/:id", api.updatePolicy)
	e.DELETE("/policies/:id", api.removePolicy)
	e.PATCH("/permissions/:id", api.updatePermission)

	e.PUT("/groups/:groupId/users", api.addUsersToGroup)
	e.GET("/groups/:groupId/users", api.getUsersByGroup)
	e.DELETE("/groups/:groupId/users/:userId", api.removeUserFromGroup)
	e.PUT("/policies/:id/permissions", api.addPermissionsToPolicy)
	e.GET("/policies/:id/permissions", api.getPermissionsByPolicy)
	e.DELETE("/policies/:policyId/permissions/:permissionId", api.removePermissionFromPolicy)
	e.PUT("/users/:id/policies", api.attachPoliciesByUser)
	e.GET("/users/:id/policies", api.getPoliciesByUser)
	e.DELETE("users/:userId/policies/:policyId", api.detachPolicyByUser)
	e.PUT("/groups/:id/policies", api.attachPoliciesByGroup)
	e.GET("/groups/:id/policies", api.getPoliciesByGroupHandler)
	e.DELETE("/groups/:groupId/policies/:policyId", api.detachPolicyByGroup)

	e.POST("/check_permissions", api.userPermissions)

	e.HTTPErrorHandler = api.errorHandler

	err := e.Start(api.address)

	if err != nil {
		Log.Fatal(
			"Api start failed",
			zap.Error(err),
		)
	}

	Log.Info("Api start successfully")
}
