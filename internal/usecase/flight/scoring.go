package flight

import (
	"math"

	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
)

func (uc *usecase) rankFlights(flights []entity.Flight) ([]entity.Flight, entity.Flight) {
	// Reference : https://www.geeksforgeeks.org/dsa/weighted-sum-method-multi-criteria-decision-making/
	// note:
	// lower price      = good
	// lower duration   = good
	// lower stops      = good
	// higher amenities = good

	var result []entity.Flight
	var bestValue entity.Flight

	weights := map[string]float64{
		"price":     0.4,
		"duration":  0.3,
		"stops":     0.2,
		"amenities": 0.1,
	}

	minStops := math.Inf(1)
	minPrice := math.Inf(1)
	minDuration := math.Inf(1)
	maxAmenities := math.Inf(-1)

	for _, flight := range flights {
		if float64(flight.Stops) < minStops {
			minStops = float64(flight.Stops)
		}
		if float64(flight.Price.Amount) < minPrice {
			minPrice = float64(flight.Price.Amount)
		}
		if float64(flight.Duration.TotalMinutes) < minDuration {
			minDuration = float64(flight.Duration.TotalMinutes)
		}
		if float64(len(flight.Amenities)) > maxAmenities {
			maxAmenities = float64(len(flight.Amenities))
		}
	}

	highestRank := float64(0)

	for _, flight := range flights {
		normalizeStops := float64(1)
		if flight.Stops > 0 {
			normalizeStops = minStops / float64(flight.Stops)
		}

		normalizePrice := float64(1)
		if flight.Price.Amount > 0 {
			normalizePrice = minPrice / float64(flight.Price.Amount)
		}

		normalizeDuration := float64(1)
		if flight.Duration.TotalMinutes > 0 {
			normalizeDuration = minDuration / float64(flight.Duration.TotalMinutes)
		}

		normalizeAmenities := float64(1)
		if maxAmenities > 0 {
			normalizeAmenities = float64(len(flight.Amenities)) / maxAmenities
		}

		scoreStops := normalizeStops * weights["stops"]
		scorePrice := normalizePrice * weights["price"]
		scoreDuration := normalizeDuration * weights["duration"]
		scoreAmenities := normalizeAmenities * weights["amenities"]

		flight.Rank = scoreStops + scorePrice + scoreDuration + scoreAmenities
		result = append(result, flight)

		if flight.Rank > highestRank {
			highestRank = flight.Rank
			bestValue = flight
		}
	}

	return result, bestValue
}
