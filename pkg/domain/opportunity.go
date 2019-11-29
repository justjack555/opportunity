package domain

import (
	"database/sql"
)

type Opportunity struct {
	DB    *sql.DB `json:"-"`
	ID    int     `json:"id"`
	Title string  `json:"title"`
}

/*
GetAllOpportunities - fetch all opportunities from db
*/
func (opp *Opportunity) GetAllOpportunities() ([]*Opportunity, error) {
	rows, err := opp.DB.Query("SELECT * FROM opportunities")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	opportunities := make([]*Opportunity, 0)
	for rows.Next() {
		opportunity := new(Opportunity)
		err := rows.Scan(&opportunity.ID, &opportunity.Title)
		if err != nil {
			return nil, err
		}
		opportunities = append(opportunities, opportunity)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return opportunities, nil
}
