package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/justjack555/opportunity/pkg/api/application"

	_ "github.com/lib/pq"
)

const opportunityPort = ":8080"

func createDbConnection() (*sql.DB, error) {
	connStr := "dbname=opportunity user=opportunity_server sslmode=disable"
	return sql.Open("postgres", connStr)
}

func main() {
	db, err := createDbConnection()
	if err != nil {
		log.Fatalln("Could not connect to db: ", err)
	}

	log.Println("Able to connec to db...")

	err = application.LoadRoutes(db)
	if err != nil {
		log.Fatalln("Unable to load routes due to error: ", err)
	}

	log.Fatalln(http.ListenAndServe(opportunityPort, nil))
}
