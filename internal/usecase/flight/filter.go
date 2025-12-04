package flight

import "github.com/kwantz/flight-search-aggregation-system/internal/entity"

func (u *usecase) filterFlights(flights []entity.Flight, req entity.FlightRequest) []entity.Flight {
	var result []entity.Flight

	for _, flight := range flights {
		if !u.isValidFlight(flight, req) {
			continue
		}
		result = append(result, flight)
	}

	return result
}

func (u *usecase) isValidFlight(flight entity.Flight, req entity.FlightRequest) bool {
	if req.PriceFrom != 0 && flight.Price.Amount < req.PriceFrom {
		return false
	}

	if req.PriceTo != 0 && flight.Price.Amount > req.PriceTo {
		return false
	}

	if req.Stops != 0 && flight.Stops > req.Stops {
		return false
	}

	if req.Airline != "" && flight.Airline.Name != req.Airline {
		return false
	}

	if req.DurationFrom != 0 && flight.Duration.TotalMinutes < req.DurationFrom {
		return false
	}

	if req.DurationTo != 0 && flight.Duration.TotalMinutes > req.DurationTo {
		return false
	}

	return true
}
