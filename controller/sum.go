package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"swagger-sample/models"
)

// SumHandler godoc
// @Summary Returns the sum of two numbers
// @Description Get the sum of two query parameters a and b
// @Tags math
// @Accept  json
// @Produce  json
// @Param a query number true "First number"
// @Param b query number true "Second number"
// @Success 200 {object} models.OpResponse
// @Failure 400 {string} string "Invalid input"
// @Router /sum [get]
func SumHandler(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, errA := strconv.ParseFloat(aStr, 64)
	b, errB := strconv.ParseFloat(bStr, 64)

	if errA != nil || errB != nil {
		http.Error(w, "Invalid query parameters 'a' or 'b'", http.StatusBadRequest)
		return
	}

	resp := models.OpResponse{
		A:   a,
		B:   b,
		Sol: a + b,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
