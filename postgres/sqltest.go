package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func main() {
    var sdate string
    tot := 0     
	db, err := sql.Open("postgres", "host=10.11.75.214 sslmode=disable user=datalex dbname=datalex password=datalex")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT to_char(ldate,'YYYY-MM-DD') as sdate, sum(total) as tot FROM syserr group by sdate order by sdate")
	if err != nil {
		log.Fatal(err)
	}
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&sdate, &tot)
        if err != nil {
            log.Fatal(err)
        }
	log.Println(sdate,tot)
}
err = rows.Err()
if err != nil {
	log.Fatal(err)
}
}