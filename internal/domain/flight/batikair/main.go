package batikair

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/kwantz/flight-search-aggregation-system/internal/constants"
	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
	"github.com/kwantz/flight-search-aggregation-system/utils"
)

type batikair struct {
}

func New() *batikair {
	return &batikair{}
}

func (a *batikair) Search(ctx context.Context) ([]entity.Flight, error) {
	retry := 3

	resp, err := a.httpPostWithRetry(ctx, retry)
	if err != nil {
		log.Printf("[ERROR][domain/batikair][Search] failed httpPost, err: %s", err.Error())
		return nil, err
	}

	return a.toFlightEntity(ctx, resp), nil
}

func (a *batikair) httpPostWithRetry(ctx context.Context, retry int) (FlightResponse, error) {
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

func (*batikair) httpPost(ctx context.Context) (FlightResponse, error) {
	var resp FlightResponse

	err := json.Unmarshal([]byte(mock), &resp)
	if err != nil {
		log.Printf("[ERROR][domain/batikair][httpPost] failed json.Unmarshal, err: %s", err.Error())
		return resp, err
	}

	rand.Seed(time.Now().UnixNano())
	sleep := time.Duration(200 + rand.Intn(200))
	time.Sleep(sleep * time.Millisecond)

	return resp, nil
}

func (a *batikair) toFlightEntity(ctx context.Context, resp FlightResponse) []entity.Flight {
	var flights []entity.Flight

	timeLayout := "2006-01-02T15:04:05-0700"

	for _, flight := range resp.Results {
		airline := strings.TrimSpace(flight.AirlineName)

		data := entity.Flight{
			ID:             fmt.Sprintf("%s_%s", flight.FlightNumber, airline),
			Provider:       airline,
			FlightNumber:   flight.FlightNumber,
			Stops:          flight.NumberOfStops,
			AvailableSeats: flight.SeatsAvailable,
			CabinClass:     mapCabinClass[flight.Fare.Class],
			Aircraft:       &flight.AircraftModel,
			Amenities:      flight.OnboardServices,
		}

		data.Airline = entity.Airline{
			Name: airline,
			Code: flight.AirlineIATA,
		}

		departure, err := time.Parse(timeLayout, flight.DepartureDateTime)
		if err != nil {
			log.Printf("[ERROR][domain/batikair][toFlightEntity] failed time.Parse, err: %s", err.Error())
			return nil
		}

		data.Departure = entity.FlightEndpoint{
			City:      constants.MapCityAirportCode[flight.Origin],
			Airport:   flight.Origin,
			Datetime:  departure,
			Timestamp: departure.Unix(),
		}

		arrival, err := time.Parse(timeLayout, flight.ArrivalDateTime)
		if err != nil {
			log.Printf("[ERROR][domain/batikair][toFlightEntity] failed time.Parse, err: %s", err.Error())
			return nil
		}

		data.Arrival = entity.FlightEndpoint{
			City:      constants.MapCityAirportCode[flight.Destination],
			Airport:   flight.Destination,
			Datetime:  arrival,
			Timestamp: arrival.Unix(),
		}

		duration := arrival.Sub(departure)
		data.Duration = entity.Duration{
			TotalMinutes: int(duration.Minutes()),
			Formatted:    utils.FormatDuration(duration),
		}

		data.Price = entity.Price{
			Amount:    flight.Fare.TotalPrice,
			Currency:  flight.Fare.CurrencyCode,
			Formatted: utils.FormatCurrency(flight.Fare.CurrencyCode, flight.Fare.TotalPrice),
		}

		baggageNotes := strings.Split(flight.BaggageInfo, ", ")
		if len(baggageNotes) == 2 {
			data.Baggage = entity.Baggage{
				CarryOn: baggageNotes[0],
				Checked: baggageNotes[1],
			}
		}

		flights = append(flights, data)
	}

	return flights
}
