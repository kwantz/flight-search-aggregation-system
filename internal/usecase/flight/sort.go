package flight

import (
	"sort"

	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
)

func (u *usecase) sortFlights(flights []entity.Flight, req entity.FlightRequest) []entity.Flight {
	sort.Slice(flights, func(left, right int) bool {
		if req.SortBy.Price == "asc" {
			return flights[left].Price.Amount < flights[right].Price.Amount
		}

		if req.SortBy.Price == "desc" {
			return flights[left].Price.Amount > flights[right].Price.Amount
		}

		if req.SortBy.Duration == "asc" {
			return flights[left].Duration.TotalMinutes < flights[right].Duration.TotalMinutes
		}

		if req.SortBy.Duration == "desc" {
			return flights[left].Duration.TotalMinutes > flights[right].Duration.TotalMinutes
		}

		if req.SortBy.Arrival == "asc" {
			return flights[left].Arrival.Datetime.Before(flights[right].Arrival.Datetime)
		}

		if req.SortBy.Arrival == "desc" {
			return flights[left].Arrival.Datetime.After(flights[right].Arrival.Datetime)
		}

		if req.SortBy.Departure == "asc" {
			return flights[left].Departure.Datetime.Before(flights[right].Departure.Datetime)
		}

		if req.SortBy.Departure == "desc" {
			return flights[left].Departure.Datetime.After(flights[right].Departure.Datetime)
		}

		return false
	})

	return flights
}
