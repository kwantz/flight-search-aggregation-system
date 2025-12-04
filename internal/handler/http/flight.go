package http

import (
	"encoding/json"
	"net/http"

	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
)

func (h *handler) PostFlightSearch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req entity.FlightRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Write([]byte(`invalid request body`))
		return
	}

	flights, err := h.ucFlight.SearchFlight(ctx, req)
	if err != nil {
		w.Write([]byte(`internal server error`))
		return
	}

	b, err := json.Marshal(flights)
	if err != nil {
		w.Write([]byte(`internal server error`))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(b)
}
