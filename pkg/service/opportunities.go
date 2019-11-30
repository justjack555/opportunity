package service

import (
	"database/sql"
	"fmt"

	"github.com/justjack555/opportunity/pkg/domain"
)

/*
AddOpportunities retrieves all opportunities for the user
*/
func AddOpportunities(db *sql.DB, name string) ([]*domain.Opportunity, error) {
	fmt.Println("AddOpportunities()")

	oppModel := &domain.Opportunity{DB: db}
	return oppModel.CreateOpportunities(name)
}
