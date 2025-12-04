package interfaces

import (
	"context"

	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
)

type IFlightUsecase interface {
	SearchFlight(ctx context.Context, req entity.FlightRequest) (entity.FlightResponse, error)
}
