package models

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Mad_bbrun struct {
	Id int
	Git_hash string `sql:"size:40"`
	Build_id int
	Database_id int
	Platform_id int
	Start_time time.Time
	Reported_at time.Time
	Time_elapsed int
	Run_type string `sql:"size:50"`
	Remote_archive string `sql:"size:100"`
	Jenkins_build_url string `sql:"size:100"`
	Successes int
	Errors int
	Failures int
	Dependency_failures int
	Skipped int
	Email string `sql:"size:255"`
	Ignore_reason string `sql:"size:255"`
}

func Get_all_bbruns() []Mad_bbrun {
	sqlConnection := "user=postgres dbname=postgres password=postgres host=10.43.0.157 port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", sqlConnection)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	var records []Mad_bbrun
	db.Table("mad_bbrun").Select("*").Limit(100).Order("id desc").Scan(&records)
	return records
}
