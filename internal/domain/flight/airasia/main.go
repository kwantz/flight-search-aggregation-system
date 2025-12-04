package airasia

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/kwantz/flight-search-aggregation-system/internal/constants"
	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
	"github.com/kwantz/flight-search-aggregation-system/utils"
)

type airasia struct {
}

func New() *airasia {
	return &airasia{}
}

func (a *airasia) Search(ctx context.Context) ([]entity.Flight, error) {
	// Next: add retry + exponential backoff
	// Next: add timeout with context

	resp, err := a.httpPost(ctx)
	if err != nil {
		log.Printf("[ERROR][domain/airasia][Search] failed httpPost, err: %s", err.Error())
		return nil, err
	}

	flights, err := a.toFlightEntity(ctx, resp)
	if err != nil {
		log.Printf("[ERROR][domain/airasia][Search] failed toFlightEntity, err: %s", err.Error())
		return nil, err
	}

	return flights, nil
}

func (*airasia) httpPost(ctx context.Context) (SearchResponse, error) {
	var resp SearchResponse

	// simulate success rate 90%
	if time.Now().Unix()%100 >= 90 {
		return resp, errors.New("internal server error")
	}

	err := json.Unmarshal([]byte(mock), &resp)
	if err != nil {
		log.Printf("[ERROR][domain/airasia][httpPost] failed json.Unmarshal, err: %s", err.Error())
		return resp, err
	}

	time.Sleep(150 * time.Millisecond)
	return resp, nil
}

func (a *airasia) toFlightEntity(ctx context.Context, resp SearchResponse) ([]entity.Flight, error) {
	var flights []entity.Flight

	timeLayout := "2006-01-02T15:04:05-07:00"

	for _, flight := range resp.Flights {
		data := entity.Flight{
			ID:             fmt.Sprintf("%s_%s", flight.Code, flight.Airline),
			Provider:       flight.Airline,
			FlightNumber:   flight.Code,
			Stops:          len(flight.Stops),
			AvailableSeats: flight.Seats,
			CabinClass:     flight.CabinClass,
		}

		data.Airline = entity.Airline{
			Name: flight.Airline,
			Code: flight.Code[:2],
		}

		departure, err := time.Parse(timeLayout, flight.DepartTime)
		if err != nil {
			log.Printf("[ERROR][domain/airasia][toFlightEntity] failed time.Parse, err: %s", err.Error())
			return nil, err
		}

		data.Departure = entity.FlightEndpoint{
			City:      constants.MapCityAirportCode[flight.FromAirport],
			Airport:   flight.FromAirport,
			Datetime:  departure,
			Timestamp: departure.Unix(),
		}

		arrival, err := time.Parse(timeLayout, flight.ArriveTime)
		if err != nil {
			log.Printf("[ERROR][domain/airasia][toFlightEntity] failed time.Parse, err: %s", err.Error())
			return nil, err
		}

		data.Arrival = entity.FlightEndpoint{
			City:      constants.MapCityAirportCode[flight.ToAirport],
			Airport:   flight.ToAirport,
			Datetime:  arrival,
			Timestamp: arrival.Unix(),
		}

		duration := arrival.Sub(departure)
		data.Duration = entity.Duration{
			TotalMinutes: int(duration.Minutes()),
			Formatted:    utils.FormatDuration(duration),
		}

		data.Price = entity.Price{
			Amount:    flight.PriceIDR,
			Currency:  constants.CurrencyIDR,
			Formatted: utils.FormatCurrency(constants.CurrencyIDR, flight.PriceIDR),
		}

		baggageNotes := strings.Split(flight.BaggageNote, ", ")
		if len(baggageNotes) == 2 {
			data.Baggage = entity.Baggage{
				CarryOn: baggageNotes[0],
				Checked: baggageNotes[1],
			}
		}

		flights = append(flights, data)
	}

	return flights, nil
}
