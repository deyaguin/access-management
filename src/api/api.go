package api

import (
	"gitlab/nefco/access-management-system/src/logger"
	"gitlab/nefco/access-management-system/src/services"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
)

var Log *zap.Logger = logger.NewLogger()

type API struct {
	*services.ServicesConf
	address string
}

func NewAPI(
	servicesConf *services.ServicesConf,
	address string,
) {
	api := &API{
		servicesConf,
		address,
	}
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	e.POST("/users", api.createUser)
	e.GET("/users", api.getUsers)
	e.GET("/users/all", api.getAllUsers)
	e.GET("/users/:userID", api.getUser)
	e.PATCH("/users/:userID", api.updateUser)
	e.DELETE("/users/:userID", api.removeUser)
	e.PUT("/users/:userID/policies", api.attachPoliciesByUser)
	e.GET("/users/:userID/policies", api.getPoliciesByUser)
	e.DELETE("/users/:userID/policies/:policyID", api.detachPolicyByUser)
	e.GET("/users/:userID/groups", api.getGroupsByUser)

	e.POST("/groups", api.createGroup)
	e.GET("/groups", api.getGroups)
	e.GET("/groups/all", api.getAllGroups)
	e.GET("/groups/:groupID", api.getGroup)
	e.PATCH("/groups/:groupID", api.updateGroup)
	e.DELETE("/groups/:groupID", api.removeGroup)
	e.PUT("/groups/:groupID/users", api.addUsersToGroup)
	e.GET("/groups/:groupID/users", api.getUsersByGroup)
	e.DELETE("/groups/:groupID/users/:userID", api.removeUserFromGroup)
	e.PUT("/groups/:groupID/policies", api.attachPoliciesByGroup)
	e.GET("/groups/:groupID/policies", api.getPoliciesByGroupHandler)
	e.DELETE("/groups/:groupID/policies/:policyID", api.detachPolicyByGroup)

	e.POST("/policies", api.createPolicy)
	e.GET("/policies", api.getPolicies)
	e.GET("/policies/all", api.getAllPolicies)
	e.GET("/policies/:policyID", api.getPolicy)
	e.PATCH("/policies/:policyID", api.updatePolicy)
	e.DELETE("/policies/:policyID", api.removePolicy)
	e.PATCH("/permissions/:policyID", api.updatePermission)
	e.PUT("/policies/:policyID/permissions", api.addPermissionsToPolicy)
	e.GET("/policies/:policyID/permissions", api.getPermissionsByPolicy)
	e.DELETE("/policies/:policyID/permissions/:permissionID", api.removePermissionFromPolicy)

	e.POST("/actions", api.createAction)
	e.GET("/actions", api.getActions)
	e.GET("/actions/:actionID", api.getAction)
	e.PATCH("/actions/:actionID", api.updateAction)
	e.DELETE("/actions/:actionID", api.removeAction)

	e.GET("/services", api.getAllServices)
	e.GET("/services/:serviceID/actions", api.getActionsByService)

	e.POST("/check_permissions", api.userPermissions)

	e.HTTPErrorHandler = api.errorHandler

	err := e.Start(api.address)

	if err != nil {
		Log.Fatal(
			"API start failed",
			zap.Error(err),
		)
	}

	Log.Info("API started successfully")
}
