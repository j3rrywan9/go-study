package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type mad_database struct {
	Id int
	Vendor string
	Version string
}

func main() {
	sqlConnection := "user=postgres dbname=postgres password=postgres host=jw-dashboard-pg.dcenter.d.com port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", sqlConnection)
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	
	var record mad_database
	
	var records []mad_database

	db.First(&record)

	db.Table("mad_database").Select("*").Where("vendor = ?", "oracle").Scan(&records)

	fmt.Printf("%d %s %s\n", record.Id, record.Vendor, record.Version)
	
	for i := 0; i < len(records); i++ {
		fmt.Println(records[i])
	}
}
