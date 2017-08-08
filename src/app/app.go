package app

import (
	"app/api"
	"app/db"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {

}

func (app *App) Init() {
	pgDB := new(db.SqlDB)
	pgDB.Connect("postgres", "host=localhost user=accessControl dbname=accesscontrol password=agryz2010")
	api := api.Api{
		pgDB,
	}
	api.Init()
}
