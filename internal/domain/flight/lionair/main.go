package lionair

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/kwantz/flight-search-aggregation-system/internal/constants"
	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
	"github.com/kwantz/flight-search-aggregation-system/utils"
)

type lionair struct {
}

func New() *lionair {
	return &lionair{}
}

func (a *lionair) Search(ctx context.Context) ([]entity.Flight, error) {
	retry := 3

	resp, err := a.httpPostWithRetry(ctx, retry)
	if err != nil {
		log.Printf("[ERROR][domain/lionair][Search] failed httpPost, err: %s", err.Error())
		return nil, err
	}

	flights, err := a.toFlightEntity(ctx, resp)
	if err != nil {
		log.Printf("[ERROR][domain/lionair][Search] failed toFlightEntity, err: %s", err.Error())
		return nil, err
	}

	return flights, nil
}

func (a *lionair) httpPostWithRetry(ctx context.Context, retry int) (FlightResponse, error) {
	var res FlightResponse
	var err error

	delay := time.Duration(100 * time.Millisecond)

	for attempt := 0; attempt < retry; attempt++ {
		res, err = a.httpPost(ctx)
		if err == nil {
			return res, nil
		}

		time.Sleep(delay)
		delay *= 2
	}

	return FlightResponse{}, err
}

func (*lionair) httpPost(ctx context.Context) (FlightResponse, error) {
	var resp FlightResponse

	err := json.Unmarshal([]byte(mock), &resp)
	if err != nil {
		log.Printf("[ERROR][domain/lionair][httpPost] failed json.Unmarshal, err: %s", err.Error())
		return resp, err
	}

	time.Sleep(200 * time.Millisecond)
	return resp, nil
}

func (a *lionair) toFlightEntity(ctx context.Context, resp FlightResponse) ([]entity.Flight, error) {
	var flights []entity.Flight

	for _, flight := range resp.Data.AvailableFlights {
		airline := strings.TrimSpace(flight.Carrier.Name)

		data := entity.Flight{
			ID:             fmt.Sprintf("%s_%s", flight.ID, airline),
			Provider:       airline,
			FlightNumber:   flight.ID,
			Stops:          flight.StopCount,
			AvailableSeats: flight.SeatsLeft,
			CabinClass:     strings.ToLower(flight.Pricing.FareType),
			Aircraft:       &flight.PlaneType,
			Amenities:      []string{},
		}

		if flight.Services.WifiAvailable {
			data.Amenities = append(data.Amenities, "wifi")
		}

		if flight.Services.MealsIncluded {
			data.Amenities = append(data.Amenities, "meals")
		}

		data.Airline = entity.Airline{
			Name: airline,
			Code: flight.Carrier.IATA,
		}

		loc, err := time.LoadLocation(flight.Schedule.DepartureTimezone)
		if err != nil {
			log.Printf("[ERROR][domain/lionair][toFlightEntity] failed time.LoadLocation for %s, err: %s", flight.Schedule.DepartureTimezone, err.Error())
			return nil, err
		}

		departureTime, err := time.ParseInLocation("2006-01-02T15:04:05", flight.Schedule.Departure, loc)
		if err != nil {
			log.Printf("[ERROR][domain/lionair][toFlightEntity] failed time.ParseInLocation for %s, err: %s", flight.Schedule.Departure, err.Error())
			return nil, err
		}

		data.Departure = entity.FlightEndpoint{
			City:      constants.MapCityAirportCode[flight.Route.From.Code],
			Airport:   flight.Route.From.Code,
			Datetime:  departureTime,
			Timestamp: departureTime.Unix(),
		}

		loc, err = time.LoadLocation(flight.Schedule.ArrivalTimezone)
		if err != nil {
			log.Printf("[ERROR][domain/lionair][toFlightEntity] failed time.LoadLocation for %s, err: %s", flight.Schedule.ArrivalTimezone, err.Error())
			return nil, err
		}

		arrivalTime, err := time.ParseInLocation("2006-01-02T15:04:05", flight.Schedule.Arrival, loc)
		if err != nil {
			log.Printf("[ERROR][domain/lionair][toFlightEntity] failed time.ParseInLocation for %s, err: %s", flight.Schedule.Arrival, err.Error())
			return nil, err
		}

		data.Arrival = entity.FlightEndpoint{
			City:      constants.MapCityAirportCode[flight.Route.To.Code],
			Airport:   flight.Route.To.Code,
			Datetime:  arrivalTime,
			Timestamp: arrivalTime.Unix(),
		}

		duration := arrivalTime.Sub(departureTime)
		data.Duration = entity.Duration{
			TotalMinutes: int(duration.Minutes()),
			Formatted:    utils.FormatDuration(duration),
		}

		data.Price = entity.Price{
			Amount:    flight.Pricing.Total,
			Currency:  flight.Pricing.Currency,
			Formatted: utils.FormatCurrency(flight.Pricing.Currency, flight.Pricing.Total),
		}

		data.Baggage = entity.Baggage{
			CarryOn: fmt.Sprintf("%s cabin", strings.TrimSpace(flight.Services.BaggageAllowance.Cabin)),
			Checked: fmt.Sprintf("%s checked", strings.TrimSpace(flight.Services.BaggageAllowance.Hold)),
		}

		flights = append(flights, data)
	}

	return flights, nil
}
