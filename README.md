# flight-search-aggregation-system

# Setup
1. Install Go 1.16 or higher
2. Clone the repository
3. Run `go mod tidy` to install dependencies
4. Run `go run cmd/http/main.go` to start the server

# Usage
1. Send a POST request to `http://localhost:3000/api/v1/flight/search` with the following JSON body:
```json
{
    "origin": "CGK", 
    "destination": "DPS", 
    "departure_date": "2025-12-15", 
    "return_date": null, 
    "passengers": 1, 
    "cabin_class": "economy"
}
```
2. The server will return a JSON response with the aggregated flight results from multiple airlines.

# Functionality

1. Aggregate Flight Data from Multiple Sources
    * `func toFlightEntity()`
        * [AirAsia](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/airasia/main.go#L79)
        * [BatikAir](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/batikair/main.go#L67)
        * [Garuda](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/garudaindonesia/main.go#L73)
        * [LionAir](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/lionair/main.go#L73)

2. Search Capabilities
    * `func searchFlight()` [Link](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/usecase/flight/search.go#L7)

3. Filter Capabilities
    * `func filterFlights()` [Link](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/usecase/flight/filter.go#L5)

4. Scoring Ranking
    * `func rankFlights()` [Link](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/usecase/flight/scoring.go#L9)

5. Handle Data Inconsistencies
    * `func validateRequest()` [Link](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/usecase/flight/validation.go#L10)

6. Handle timezone conversions
    * [AirAsia](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/airasia/main.go#L99-L123)
    * [BatikAir](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/batikair/main.go#L91-L115)
    * [Garuda](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/garudaindonesia/main.go#L97-L121)
    * [LionAir](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/lionair/main.go#L103-L139)
    * [Request Body](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/usecase/flight/utils.go)

7. Implement retry logic with exponential backoff
    * `func httpPostWithRetry()`
        * [AirAsia](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/airasia/main.go#L42)
        * [BatikAir](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/batikair/main.go#L35)
        * [Garuda](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/garudaindonesia/main.go#L41)
        * [LionAir](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/lionair/main.go#L41)

8. Support for currency display
    * `func FormatCurrency()` [Link](https://github.com/kwantz/flight-search-aggregation-system/blob/main/utils/format.go#L28)
        * [AirAsia](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/airasia/main.go#L134)
        * [BatikAir](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/batikair/main.go#L126)
        * [Garuda](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/garudaindonesia/main.go#L132)
        * [LionAir](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/domain/flight/lionair/main.go#L150)

9. Parallel provider queries [Link](https://github.com/kwantz/flight-search-aggregation-system/blob/main/internal/usecase/flight/main.go#L36-L62)
