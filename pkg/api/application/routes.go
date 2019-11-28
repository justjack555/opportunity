package application

import (
	"net/http"

	"github.com/justjack555/opportunity/pkg/api/application/handlers"
)

/*
LoadRoutes loads application routes
*/
func LoadRoutes() error {
	http.Handle("/", &handlers.IndexHandler{})

	return nil
}
