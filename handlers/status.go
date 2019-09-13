package handlers

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/Sanny_Lebedev/test6/fibb"
)

// home is a simple HTTP handler function which writes a response.
func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	nums, done := fibb.Status(r.FormValue("UID"))
	json.NewEncoder(w).Encode(answer{UID: r.FormValue("UID"), Success: true, Meta: meta{Nums: nums, Last: nums[len(nums)-1:][0]}, Done: done})
}
