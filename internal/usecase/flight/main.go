package flight

import (
	"context"
	"log"
	"time"

	"github.com/kwantz/flight-search-aggregation-system/internal/constants"
	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
	"github.com/kwantz/flight-search-aggregation-system/internal/interfaces"
)

type usecase struct {
	flightProvider []interfaces.IFlightProviderDomain
}

func New(flightProvider []interfaces.IFlightProviderDomain) *usecase {
	return &usecase{
		flightProvider: flightProvider,
	}
}

func (u *usecase) SearchFlight(ctx context.Context, req entity.FlightRequest) (entity.FlightResponse, error) {
	req, valid := u.validateRequest(ctx, req)
	if !valid {
		log.Printf("[INFO][usecase/flight][SearchFlight] invalid request body")
		return entity.FlightResponse{}, constants.ErrBadRequest
	}

	flights := []entity.Flight{}
	flightProviderSuccess := 0
	flightProviderFail := 0
	flightProviderStartTime := time.Now()

	// Improvement: change to async + wait group + lock mutex

	for _, provider := range u.flightProvider {
		flightsFromProvider, err := provider.Search(ctx)
		if err != nil {
			log.Printf("[ERROR][usecase/flight][SearchFlight] failed provider.Search, err: %s", err.Error())

			flightProviderFail++
			continue
		}

		flightProviderSuccess++
		flights = append(flights, flightsFromProvider...)
	}

	flightProviderTimeDuration := time.Since(flightProviderStartTime)

	// Improvement: for better performance, move these inside async process above
	flights = u.searchFlight(flights, req)  // move to async
	flights = u.filterFlights(flights, req) // move to async

	flights = u.sortFlights(flights, req)
	flights, bestValue := u.rankFlights(flights)

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
			CacheHit:           false, // Improvement: set true if implement redis cache
		},
		BestValue: bestValue,
		Flights:   flights,
	}

	return resp, nil
}
