package persistence

// Database interaction
// https://godoc.org/github.com/lib/pq


import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"github.com/Akim-Delli/feedrover/config"
)

type Content struct {
	Id string
	Content_id string
	Content_type string
	Page int
	Date_unix string
	Url string
}

func Persist(content *Content)  {
	db, err := sql.Open("postgres",
						fmt.Sprintf("user=%s host=%s dbname=%v password=%v sslmode=%s",
							         config.DB_USER, config.DB_HOST, config.DB_NAME, config.DB_PASSWORD, config.DB_SSL_MODE))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	timestamp, err := strconv.Atoi(content.Date_unix)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query(`INSERT INTO elle SELECT $1, $2, $3, $4, $5, $6, Now() WHERE NOT EXISTS (SELECT 1 FROM elle WHERE id=$1);`,
						   content.Id, content.Content_id, content.Content_type, timestamp, 1, content.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

