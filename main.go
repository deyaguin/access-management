package main

import (
	"gitlab/nefco/accessControl/app"
	"gitlab/nefco/accessControl/migrations"
)

func main() {
	m := new(migrations.Migration)
	m.Init()
	app := new(app.App)
	app.Init()
}
