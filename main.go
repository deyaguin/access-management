package main

import (
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab/nefco/access-management-system/src/api"
	"gitlab/nefco/access-management-system/src/migrations"
	"gitlab/nefco/access-management-system/src/services"
	"gitlab/nefco/access-management-system/src/storage"
)

func main() {

	new(migrations.Migration).Init()
	//pgDB, err := storage.SqlDBCreator("postgres", "host=localhost user=access-management-system dbname=accesscontrol password=agryz2010")
	sqliteDB, err := storage.SqlDBCreator("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	userService := services.NewUserService(sqliteDB)
	groupService := services.NewGroupService(sqliteDB)
	policyService := services.NewPolicyService(sqliteDB)
	permissionService := services.NewPermissionService(sqliteDB)
	relationsService := services.NewRelationsService(sqliteDB)
	permissionsCheckService := services.NewPermissionsCheckService(sqliteDB)

	api.NewAPI(
		userService,
		groupService,
		policyService,
		permissionService,
		relationsService,
		permissionsCheckService,
	)
}
