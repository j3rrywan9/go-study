package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Mad_platform struct {
	Id int
	Name string `sql:"size:50"`
}

func Get_all_platforms(dbConn gorm.DB) []Mad_platform {
	var records []Mad_platform
	dbConn.Table("mad_platform").Select("*").Order("id").Scan(&records)
	return records
}
