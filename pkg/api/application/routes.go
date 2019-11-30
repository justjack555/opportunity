package application

import (
	"database/sql"
	"net/http"

	"github.com/justjack555/opportunity/pkg/api/application/handlers"
)

/*
LoadRoutes loads application routes
*/
func LoadRoutes(db *sql.DB) error {
	http.Handle("/", &handlers.IndexHandler{DB: db})
	http.Handle("/opportunities", &handlers.OpportunityHandler{DB: db})

	return nil
}
