package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MysqlDB db
type MysqlDB struct {
	db *gorm.DB
}

// NewMysqlDB create a db
func NewMysqlDB(uri string) (*MysqlDB, error) {
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		return nil, err
	}

	if !db.HasTable(&Ancient{}) {
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Ancient{}).Error; err != nil {
			panic(err)
		}
	}

	return &MysqlDB{
		db: db,
	}, nil
}
