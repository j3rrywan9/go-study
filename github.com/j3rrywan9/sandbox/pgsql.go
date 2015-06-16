package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type mad_database struct {
	id int
	vendor string
	version string
}

func main() {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres password=postgres host=jw-dashboard-pg.dcenter.d.com port=5432 sslmode=disable")

	if err != nil {
		panic(err)
	}
	
	rows, err := db.Query("select * from mad_database")

	if err != nil {
		panic(err)
	}
	
	defer rows.Close()
	
	for rows.Next() {
		var record mad_database
		if err := rows.Scan(&record.id, &record.vendor, &record.version); err != nil {
			panic(err)
		}
		fmt.Printf("%d %s %s\n", record.id, record.vendor, record.version)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}
