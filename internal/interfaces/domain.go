package interfaces

import (
	"context"

	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
)

type IFlightProviderDomain interface {
	Search(ctx context.Context) ([]entity.Flight, error)
}
