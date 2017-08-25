package api

import (
	"github.com/labstack/echo"
	"gitlab/nefco/accessControl/app/services"
)

type Api struct {
	userService       services.UserService
	groupService      services.GroupService
	policyService     services.PolicyService
	permissionService services.PermissionService
	relationsService  services.RelationsService
}

func NewAPI(
	userService services.UserService,
	groupService services.GroupService,
	policyService services.PolicyService,
	permissionService services.PermissionService,
	relationsService services.RelationsService,
) {
	api := &Api{
		userService,
		groupService,
		policyService,
		permissionService,
		relationsService,
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
	//
	//e.PUT("/groups/:id/users", a.addUsersToGroup)
	//e.GET("/groups/:id/users", a.getUsersByGroupHandler)
	//e.DELETE("/groups/:gid/users/:uid", a.removeUserFromGroup)
	//e.PUT("/policies/:id/permissions", a.addPermissionsToPolicy)
	//e.GET("/policies/:id/permissions", a.getPermissionsByPolicyHandler)
	//e.DELETE("/policies/:polid/permissions/:perid", a.removePermission)
	//e.PUT("/users/:id/policies", a.attachPoliciesByUser)
	//e.GET("/users/:id/policies", a.getPoliciesByUserHandler)
	//e.DELETE("users/:uid/policies/:pid", a.detachPolicyByUser)
	//e.PUT("/groups/:id/policies", a.attachPoliciesByGroup)
	//e.GET("/groups/:id/policies", a.getPoliciesByGroupHandler)
	//e.DELETE("/groups/:gid/policies/:pid", a.detachPolicyByGroup)
	//
	//e.PATCH("/permissions/:id", a.updatePermission)
	////e.DELETE("/permissions/:id", a.removePermission)
	//
	e.POST("/check_permissions", api.userPermissions)
	e.Logger.Fatal(e.Start(":1535"))
}
