package flight

import "github.com/kwantz/flight-search-aggregation-system/internal/interfaces"

type usecase struct {
	flightProvider []interfaces.IFlightProviderDomain
}

func New(flightProvider []interfaces.IFlightProviderDomain) *usecase {
	return &usecase{
		flightProvider: flightProvider,
	}
}
