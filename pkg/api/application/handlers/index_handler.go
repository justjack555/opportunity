package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/justjack555/opportunity/pkg/domain"

	"github.com/justjack555/opportunity/pkg/service"
)

/*
IndexHandler handles requests to index
*/
type IndexHandler struct {
	DB *sql.DB
}

/*
OpportunitiesResponse api response to request for all opportunities
*/
type OpportunitiesResponse struct {
	Opportunities []*domain.Opportunity `json:"opportunities"`
}

func (ih *IndexHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	opportunities, err := service.GetOpportunities(ih.DB)
	if err != nil {
		log.Println("ih.ServeHttp(): err retrieving from db: ", err)
	}

	serializedOpps, err := json.Marshal(OpportunitiesResponse{Opportunities: opportunities})
	if err != nil {
		log.Println("ih.ServeHttp(): err marshalling: ", err)
	}

	resp.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	_, err = resp.Write(serializedOpps)
	if err != nil {
		log.Println("ih.ServeHttp(): err: ", err)
	}
}
