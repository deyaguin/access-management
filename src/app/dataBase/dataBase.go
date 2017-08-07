package dataBase

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	Vendor string
	Url string
}

func (dataBase *DB) Connect() error {
	db, err := gorm.Open(dataBase.Vendor, dataBase.Url)
	defer db.Close()
	return err
}
