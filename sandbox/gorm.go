package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Mad_database struct {
	Id int
	Vendor string
	Version string
}

func main() {
	connectionString := "user=postgres dbname=postgres password=postgres host=172.16.100.9 port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", connectionString)
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
