package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Gormd struct {
	ID   string `gorm:"column:id"`
	Name string `gorm:"column:name"`
}
type FatemehTest struct {
	ID   string `gorm:"primary_key"`
	Name string
}

var DB *gorm.DB

func OpenTestConnection(constr string) (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", constr)
	db.LogMode(true)
	return
}

func DeleteTable(DB *gorm.DB, table Gormd) error {
	if err := DB.DropTableIfExists(table).Error; err != nil {
		return err
	}
	return nil
}
