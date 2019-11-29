package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/justjack555/opportunity/pkg/service"
)

/*
IndexHandler handles requests to index
*/
type IndexHandler struct {
	DB *sql.DB
}

func (ih *IndexHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("ih.ServeHttp()")
	serverResp, err := service.GetOpportunities(ih.DB)
	if err != nil {
		log.Fatalln("err: ", err)
	}

	serializedOpps, err := json.Marshal(serverResp)
	if err != nil {
		log.Fatalln("ih.ServeHttp(): err: ", err)
	}

	_, err = resp.Write(serializedOpps)
	if err != nil {
		log.Fatalln("ih.ServeHttp(): err: ", err)
	}

	fmt.Println("ih.ServeHttp(): db resp: ", serializedOpps)
}
