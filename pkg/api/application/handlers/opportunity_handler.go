package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/justjack555/opportunity/pkg/domain"

	"github.com/justjack555/opportunity/pkg/service"
)

/*
IndexHandler handles requests to index
*/
type OpportunityHandler struct {
	DB *sql.DB
}

/*
OpportunitiesPostRequest api request to request for all opportunities
*/
type OpportunitiesPostRequest struct {
	Name string `json:"opportunityName"`
}

/*
OpportunitiesPostResponse api response to request for all opportunities
*/
type OpportunitiesPostResponse struct {
	Opportunities []*domain.Opportunity `json:"opportunities"`
}

func (oh *OpportunityHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("Hits opp handler route...")
	requestBody := &OpportunitiesPostRequest{}

	defer req.Body.Close()
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("oh.ServeHttp(): err reading body: ", err)
		return
	}

	err = json.Unmarshal(reqBody, requestBody)
	if err != nil {
		log.Println("oh.ServeHttp(): err unmarshalling request body: ", err)
		return
	}

	opportunities, err := service.AddOpportunities(oh.DB, requestBody.Name)
	if err != nil {
		log.Println("oh.ServeHttp(): err retrieving from db: ", err)
	}

	serializedOpps, err := json.Marshal(OpportunitiesPostResponse{Opportunities: opportunities})
	if err != nil {
		log.Println("oh.ServeHttp(): err marshalling: ", err)
	}

	resp.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	_, err = resp.Write(serializedOpps)
	if err != nil {
		log.Println("oh.ServeHttp(): err: ", err)
	}
}
