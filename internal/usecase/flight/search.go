package flight

import (
	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
)

func (u *usecase) searchFlight(flights []entity.Flight, req entity.FlightRequest) []entity.Flight {
	result := []entity.Flight{}

	for _, flight := range flights {
		if flight.Departure.Airport != req.Origin {
			continue
		}

		if flight.Arrival.Airport != req.Destination {
			continue
		}

		if flight.Departure.Datetime.Before(req.Departure) {
			continue
		}

		if flight.AvailableSeats < req.Passengers {
			continue
		}

		if flight.CabinClass != req.CabinClass {
			continue
		}

		result = append(result, flight)
	}

	return result
}
