package domain

import (
	"database/sql"
	"fmt"
	"log"
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
	fmt.Println("getAllOpportunities()")
	rows, err := opp.DB.Query("SELECT * FROM opportunities")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	opportunities := make([]*Opportunity, 0)
	for rows.Next() {
		opportunity := new(Opportunity)
		err := rows.Scan(&opportunity.ID, &opportunity.Title)
		if err != nil {
			log.Fatal(err)
		}
		opportunities = append(opportunities, opportunity)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, opportunity := range opportunities {
		fmt.Printf("%d, %s\n", opportunity.ID, opportunity.Title)
	}

	return opportunities, nil
}
