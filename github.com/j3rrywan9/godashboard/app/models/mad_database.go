package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Mad_database struct {
	Id int
	Vendor string `sql:"size:50"`
	Version string `sql:"size:50"`
}

func Get_all_databases(dbConn gorm.DB) []Mad_database {
	var records []Mad_database
	dbConn.Table("mad_database").Select("*").Order("id").Scan(&records)
	return records
}
