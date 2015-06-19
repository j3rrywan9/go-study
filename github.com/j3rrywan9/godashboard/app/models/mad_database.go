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

func Get_all_databases() []Mad_database {
	sqlConnection := "user=postgres dbname=postgres password=postgres host=10.43.0.157 port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", sqlConnection)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	var records []Mad_database
	db.Table("mad_database").Select("*").Order("id").Scan(&records)
	return records
}
