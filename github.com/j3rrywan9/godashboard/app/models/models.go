package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Mad_platform struct {
	Id int
	Name string `sql:"size:50"`
}

func Get_platform_name_by_id(id int) Mad_platform {
	sqlConnection := "user=postgres dbname=postgres password=postgres host=10.43.0.157 port=5432 sslmode=disable"
	
	db, err := gorm.Open("postgres", sqlConnection)
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)

	var record Mad_platform
	
	db.Table("mad_platform").Select("*").Where("id = ?", id).Scan(&record)
	
	return record
}