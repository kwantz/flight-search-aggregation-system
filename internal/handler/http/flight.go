package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kwantz/flight-search-aggregation-system/internal/constants"
	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
)

func (h *handler) PostFlightSearch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req entity.FlightRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[ERROR][handler/flight][PostFlightSearch] failed json.NewDecoder.Decode, err: %s", err.Error())
		w.Write([]byte(`bad request - invalid request body`))
		return
	}

	flights, err := h.ucFlight.SearchFlight(ctx, req)
	if err == constants.ErrBadRequest {
		log.Printf("[ERROR][handler/flight][PostFlightSearch] invalid request body")
		w.Write([]byte(`bad request - invalid request body`))
		return
	}

	if err != nil {
		log.Printf("[ERROR][handler/flight][PostFlightSearch] failed ucFlight.SearchFlight, err: %s", err.Error())
		w.Write([]byte(`internal server error`))
		return
	}

	b, err := json.Marshal(flights)
	if err != nil {
		log.Printf("[ERROR][handler/flight][PostFlightSearch] failed json.Marshal, err: %s", err.Error())
		w.Write([]byte(`internal server error`))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(b)
}
