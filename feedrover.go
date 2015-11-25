package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type StringSlice []string

func main() {
	resp, err := http.Get("http://hipsterjesus.com/api")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("response Body:", string(body))

	db, err := sql.Open("postgres", "user=feedrover host=/var/run/postgresql dbname=feedrover password=feedrover sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var (
		equip_id int
		color string
		)
	rows, err := db.Query("select equip_id, color from playground")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&equip_id, &color)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(equip_id, color)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}


}



