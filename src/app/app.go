package app

import (
	"app/api"
	"app/dataBase"
)

type App struct {

}

func (app *App) Init() {
	pgDB := &dataBase.DB{
		"postgres",
		"host=localhost user=accessControl dbname=accesscontrol password=agryz2010",
	}
	web := &api.Api{
		pgDB.Connect(),
	}
	web.Init()
}
