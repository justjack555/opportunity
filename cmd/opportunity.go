package main

import (
	"log"
	"net/http"

	"github.com/justjack555/opportunity/pkg/api/application"
)

const opportunityPort = ":8080"

func main() {
	err := application.LoadRoutes()
	if err != nil {
		log.Fatalln("Unable to load routes due to error: ", err)
	}

	log.Fatalln(http.ListenAndServe(opportunityPort, nil))
}
