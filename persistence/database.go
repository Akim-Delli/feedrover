package persistence

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type Content struct {
	Id string
	Content_id string
	Content_type string
	Page string
	Date_unix string
	Url string
}

func Persist(content *Content)  {
	db, err := sql.Open("postgres", "user=feedrover host=/var/run/postgresql dbname=feedrover password=feedrover sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	timestamp, err := strconv.Atoi(content.Date_unix)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("INSERT INTO cosmopolitan VALUES ($1, $2, $3, $4, $5, $6)", content.Id, content.Content_id, content.Content_type, timestamp, 1, content.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
//	for rows.Next() {
//		err := rows.Scan(&equip_id, &color)
//		if err != nil {
//			log.Fatal(err)
//		}
//		log.Println(equip_id, color)
//	}
//	err = rows.Err()
//	if err != nil {
//		log.Fatal(err)
//	}
}

