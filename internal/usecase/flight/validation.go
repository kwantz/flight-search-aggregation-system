package flight

import (
	"context"
	"log"

	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
)

func (uc *usecase) validateRequest(ctx context.Context, req entity.FlightRequest) (entity.FlightRequest, bool) {
	if req.Origin == "" {
		log.Printf("[INFO][usecase/flight][validateRequest] origin is empty")
		return req, false
	}

	if req.Destination == "" {
		log.Printf("[INFO][usecase/flight][validateRequest] destination is empty")
		return req, false
	}

	if req.DepartureDate == "" {
		log.Printf("[INFO][usecase/flight][validateRequest] departureDate is empty")
		return req, false
	}

	if req.Passengers <= 0 {
		log.Printf("[INFO][usecase/flight][validateRequest] passengers is non-positive")
		return req, false
	}

	if req.CabinClass == "" {
		log.Printf("[INFO][usecase/flight][validateRequest] cabinClass is empty")
		return req, false
	}

	departureTime, err := uc.getDepartureTime(ctx, req)
	if err != nil {
		log.Printf("[ERROR][usecase/flight][validateRequest] failed getDepartureTime, err: %s", err.Error())
		return req, false
	}

	returnTime, err := uc.getReturnTime(ctx, req)
	if err != nil {
		log.Printf("[ERROR][usecase/flight][validateRequest] failed getReturnTime, err: %s", err.Error())
		return req, false
	}

	if !returnTime.IsZero() && returnTime.Before(departureTime) {
		log.Printf("[INFO][usecase/flight][validateRequest] returnTime is before departureTime")
		return req, false
	}

	req.Departure = departureTime
	req.Return = returnTime

	return req, true
}
