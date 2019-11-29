package service

import (
	"database/sql"
	"fmt"

	"github.com/justjack555/opportunity/pkg/domain"
)

/*
GetOpportunities retrieves all opportunities for the user
*/
func GetOpportunities(db *sql.DB) ([]*domain.Opportunity, error) {
	fmt.Println("GetOpportunities()")

	oppModel := &domain.Opportunity{DB: db}
	return oppModel.GetAllOpportunities()
}
