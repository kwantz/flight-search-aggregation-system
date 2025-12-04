package flight

import (
	"context"
	"log"
	"time"

	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
)

func (u *usecase) SearchFlight(ctx context.Context, req entity.FlightRequest) (entity.FlightResponse, error) {
	var flights []entity.Flight

	// Next: Handle Data Inconsistencies / validation
	// sort field is mandatory
	// check return date is after departure date

	departureTime, err := u.getDepartureTime(ctx, req)
	if err != nil {
		log.Printf("[ERROR][usecase/flight][Search] failed getDepartureTime, err: %s", err.Error())
		return entity.FlightResponse{}, err
	}

	returnTime, err := u.getReturnTime(ctx, req)
	if err != nil {
		log.Printf("[ERROR][usecase/flight][Search] failed getReturnTime, err: %s", err.Error())
		return entity.FlightResponse{}, err
	}

	req.Departure = departureTime
	req.Return = returnTime

	flightProviderSuccess := 0
	flightProviderFail := 0
	flightProviderStartTime := time.Now()

	// Next: change to async + wait group + lock mutex
	for _, provider := range u.flightProvider {
		flightsFromProvider, err := provider.Search(ctx)
		if err != nil {
			log.Printf("[ERROR][usecase/flight][Search] failed provider.Search, err: %s", err.Error())

			flightProviderFail++
			continue
		}

		flightProviderSuccess++
		flights = append(flights, flightsFromProvider...)
	}

	flightProviderTimeDuration := time.Since(flightProviderStartTime)

	// Next: for better performance, move these inside async process above
	flights = u.searchFlight(flights, req)
	flights = u.filterFlights(flights, req)
	flights = u.sortFlights(flights, req)

	// Next: find Price Comparison & Ranking

	resp := entity.FlightResponse{
		SearchCriteria: entity.SearchCriteria{
			Origin:        req.Origin,
			Destination:   req.Destination,
			DepartureDate: req.DepartureDate,
			Passengers:    req.Passengers,
			CabinClass:    req.CabinClass,
		},
		Metadata: entity.FlightResponseMetadata{
			TotalResults:       len(flights),
			ProvidersQueried:   len(u.flightProvider),
			ProvidersSucceeded: flightProviderSuccess,
			ProvidersFailed:    flightProviderFail,
			SearchTimeMs:       flightProviderTimeDuration.Milliseconds(),
			CacheHit:           false, // Next: set true if implement redis cache
		},
		Flights: flights,
	}

	return resp, nil
}

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
