package garudaindonesia

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

type garuda struct {
}

func New() *garuda {
	return &garuda{}
}

func (a *garuda) Search(ctx context.Context) ([]entity.Flight, error) {
	retry := 3

	resp, err := a.httpPostWithRetry(ctx, retry)
	if err != nil {
		log.Printf("[ERROR][domain/garuda][Search] failed httpPost, err: %s", err.Error())
		return nil, err
	}

	flights, err := a.toFlightEntity(ctx, resp)
	if err != nil {
		log.Printf("[ERROR][domain/garuda][Search] failed toFlightEntity, err: %s", err.Error())
		return nil, err
	}

	return flights, nil
}

func (a *garuda) httpPostWithRetry(ctx context.Context, retry int) (FlightResponse, error) {
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

func (*garuda) httpPost(ctx context.Context) (FlightResponse, error) {
	var resp FlightResponse

	err := json.Unmarshal([]byte(mock), &resp)
	if err != nil {
		log.Printf("[ERROR][domain/garuda][httpPost] failed json.Unmarshal, err: %s", err.Error())
		return resp, err
	}

	time.Sleep(100 * time.Millisecond)
	return resp, nil
}

func (a *garuda) toFlightEntity(ctx context.Context, resp FlightResponse) ([]entity.Flight, error) {
	var flights []entity.Flight

	timeLayout := "2006-01-02T15:04:05-07:00"

	for _, flight := range resp.Flights {
		airline := strings.TrimSpace(flight.Airline)

		data := entity.Flight{
			ID:             fmt.Sprintf("%s_%s", flight.FlightID, airline),
			Provider:       airline,
			FlightNumber:   flight.FlightID,
			Stops:          flight.Stops,
			AvailableSeats: flight.AvailableSeats,
			CabinClass:     flight.FareClass,
			Aircraft:       &flight.Aircraft,
			Amenities:      flight.Amenities,
		}

		data.Airline = entity.Airline{
			Name: airline,
			Code: flight.AirlineCode,
		}

		departure, err := time.Parse(timeLayout, flight.Departure.Time)
		if err != nil {
			log.Printf("[ERROR][domain/garuda][toFlightEntity] failed time.Parse, err: %s", err.Error())
			return nil, err
		}

		data.Departure = entity.FlightEndpoint{
			City:      constants.MapCityAirportCode[flight.Departure.Airport],
			Airport:   flight.Departure.Airport,
			Datetime:  departure,
			Timestamp: departure.Unix(),
		}

		arrival, err := time.Parse(timeLayout, flight.Arrival.Time)
		if err != nil {
			log.Printf("[ERROR][domain/garuda][toFlightEntity] failed time.Parse, err: %s", err.Error())
			return nil, err
		}

		data.Arrival = entity.FlightEndpoint{
			City:      constants.MapCityAirportCode[flight.Arrival.Airport],
			Airport:   flight.Arrival.Airport,
			Datetime:  arrival,
			Timestamp: arrival.Unix(),
		}

		duration := arrival.Sub(departure)
		data.Duration = entity.Duration{
			TotalMinutes: int(duration.Minutes()),
			Formatted:    utils.FormatDuration(duration),
		}

		data.Price = entity.Price{
			Amount:    flight.Price.Amount,
			Currency:  flight.Price.Currency,
			Formatted: utils.FormatCurrency(flight.Price.Currency, flight.Price.Amount),
		}

		data.Baggage = entity.Baggage{
			CarryOn: fmt.Sprintf("%dkg cabin", flight.Baggage.CarryOn),
			Checked: fmt.Sprintf("%dkg checked", flight.Baggage.Checked),
		}

		flights = append(flights, data)
	}

	return flights, nil
}
