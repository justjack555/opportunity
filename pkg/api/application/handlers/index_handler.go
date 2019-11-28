package handlers

import (
	"fmt"
	"net/http"
)

/*
IndexHandler handles requests to index
*/
type IndexHandler struct{}

func (ih *IndexHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("ih.ServeHttp()")
}
