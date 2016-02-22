package models

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Mad_build struct {
	Id int
	Build_date time.Time
	Version string `sql:"size:20"`
	Git_hash string `sql:"size:40"`
	Upgrade_start_version string `sql:"size:20"`
}

func Get_all_builds(dbConn gorm.DB) []Mad_build {
	var records []Mad_build
	dbConn.Table("mad_build").Select("*").Order("id desc").Scan(&records)
	return records
}
