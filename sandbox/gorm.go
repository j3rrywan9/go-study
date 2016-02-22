package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Mad_database struct {
	Id int
	Vendor string
	Version string
}

func main() {
	sqlConnection := "user=postgres dbname=postgres password=postgres host=10.43.0.157 port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", sqlConnection)
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	
	var records []Mad_database

	db.Table("mad_database").Select("*").Where("vendor = ?", "oracle").Scan(&records)
	
	for i := 0; i < len(records); i++ {
		fmt.Println(records[i])
	}
}
