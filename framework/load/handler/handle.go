package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	usecase "github.com/cledupe/cross-commerce-sort-number/application/use_case"
	"github.com/cledupe/cross-commerce-sort-number/framework/infra"
)

var ITEM_PER_PAGE = 100

type ShowNumberResponse struct {
	Numbers []float64 `json:"listNumbers"`
}

func GetNumbers(w http.ResponseWriter, r *http.Request) {
	var showNumberResponse ShowNumberResponse
	var dbJson infra.JsonDb
	page := r.URL.Query().Get("page")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		w.Write([]byte("{[]}"))
		return
	}

	showNumberResponse.Numbers = usecase.ShowNumbers(ITEM_PER_PAGE, pageInt, dbJson)
	jsonResponse, jsonError := json.Marshal(showNumberResponse)

	if jsonError != nil {
		log.Print(jsonError)
	}

	buildResponse(w, r, jsonResponse)

}

func buildResponse(w http.ResponseWriter, r *http.Request, content []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(content)

	if r.Method == http.MethodOptions {
		return
	}

}
