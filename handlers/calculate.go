package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bitbucket.org/Sanny_Lebedev/test6/fibb"
	"github.com/satori/go.uuid"
)

type (
	answer struct {
		UID     string `json:"UID"`
		Success bool   `json:"success"`
		Done    bool   `json:"done"`
		Meta    meta   `json:"meta"`
	}

	meta struct {
		Last int64   `json:"last"`
		Nums []int64 `json:"nums"`
	}
)

// home is a simple HTTP handler function which writes a response.
func calculate(w http.ResponseWriter, r *http.Request) {

	u1 := uuid.Must(uuid.NewV4()).String()

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(answer{UID: u1, Success: true, Done: false})
	i, err := strconv.ParseInt(r.FormValue("n"), 10, 64)
	if err != nil {
		i = 0
	}
	go fibb.Calc(u1, i)

}
