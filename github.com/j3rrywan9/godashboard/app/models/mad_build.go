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

func Get_all_builds() []Mad_build {
	sqlConnection := "user=postgres dbname=postgres password=postgres host=10.43.0.157 port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", sqlConnection)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	var records []Mad_build
	db.Table("mad_build").Select("*").Limit(50).Scan(&records)
	return records
}
