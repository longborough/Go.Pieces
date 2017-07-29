package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
	sa "github.com/longborough/sysaudit"
)

func main() {
	var sdate string
	var dopen string
	tot := 0
	dopen = fmt.Sprintf("host=%s sslmode=disable user=%s dbname=%s password=%s",
						sa.Server, sa.PgUser, sa.DbName, sa.PgPw) 
	db, err := sql.Open("postgres", dopen)
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