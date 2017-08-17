package app

import (
	"app/api"
	"app/db"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {

}

func (app *App) Init() {
	pgDB, err := db.SqlDBCreator("postgres", "host=localhost user=accessControl dbname=accesscontrol password=agryz2010")
	if err != nil {
		panic(err)
	}
	api := api.Api{
		pgDB,
	}
	api.Init()
}
