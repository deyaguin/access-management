package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"gitlab/nefco/accessControl/migrations/models"
	"gopkg.in/gormigrate.v1"
)

type Migration struct {
}

var migrate1 = &gormigrate.Migration{
	ID: "0",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(
			&models.Service{},
			&models.Action{},
			&models.User{},
			&models.Group{},
			&models.Policy{},
			&models.Permission{},
		).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.DropTable(
			&models.Service{},
			&models.Action{},
			&models.User{},
			&models.Group{},
			&models.Policy{},
			&models.Permission{},
		).Error
	},
}

func (migration *Migration) Init() {
	//pgDB, err := gorm.Open("postgres", "host=localhost user=accessControl dbname=test password=agryz2010")
	sqliteDB, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	//if err = pgDB.DB().Ping(); err != nil {
	//	log.Fatal(err)
	//}
	if err = sqliteDB.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	//pgDB.LogMode(true)
	sqliteDB.LogMode(true)

	//m := gormigrate.New(pgDB, gormigrate.DefaultOptions, []*gormigrate.Migration{
	//	migrate1,
	//})
	m := gormigrate.New(sqliteDB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migrate1,
	})

	if err = m.Migrate(); err != nil {
		log.Fatal(err)
	}

	log.Print("migration successfully")

}
