package flight

import (
	"context"
	"log"
	"time"

	"github.com/kwantz/flight-search-aggregation-system/internal/entity"
)

func (uc *usecase) getDepartureTime(ctx context.Context, req entity.FlightRequest) (time.Time, error) {
	departureDate, err := time.Parse("2006-01-02", req.DepartureDate)
	if err != nil {
		log.Printf("[ERROR][usecase/flight][getDepartureTime] failed time.Parse, err: %s", err.Error())
		return time.Time{}, err
	}

	timezone := "Asia/Jakarta"
	if req.DepartureTimezone != "" {
		timezone = req.DepartureTimezone
	}

	departureTimezone, err := time.LoadLocation(timezone)
	if err != nil {
		log.Printf("[ERROR][usecase/flight][getDepartureTime] failed time.Parse, err: %s", err.Error())
		return time.Time{}, err
	}

	timeStr := "00:00:00"
	if req.DepartureTime != "" {
		timeStr = req.DepartureTime
	}

	departureTime, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		log.Printf("[ERROR][usecase/flight][getDepartureTime] failed time.Parse, err: %s", err.Error())
		return time.Time{}, err
	}

	departure := time.Date(
		departureDate.Year(),
		departureDate.Month(),
		departureDate.Day(),
		departureTime.Hour(),
		departureTime.Minute(),
		departureTime.Second(),
		0,
		departureTimezone,
	)

	return departure, nil
}

func (u *usecase) getReturnTime(ctx context.Context, req entity.FlightRequest) (time.Time, error) {
	if req.ReturnDate == nil {
		return time.Time{}, nil
	}

	returnDate, err := time.Parse("2006-01-02", *req.ReturnDate)
	if err != nil {
		log.Printf("[ERROR][usecase/flight][getReturnTime] failed time.Parse, err: %s", err.Error())
		return time.Time{}, err
	}

	timezone := "Asia/Jakarta"
	if req.ReturnTimezone != nil && *req.ReturnTimezone != "" {
		timezone = *req.ReturnTimezone
	}

	returnTimezone, err := time.LoadLocation(timezone)
	if err != nil {
		log.Printf("[ERROR][usecase/flight][getReturnTime] failed time.Parse, err: %s", err.Error())
		return time.Time{}, err
	}

	timeStr := "00:00:00"
	if req.ReturnTime != nil && *req.ReturnTime != "" {
		timeStr = *req.ReturnTime
	}

	returnTime, err := time.Parse("15:04:05", timeStr)
	if err != nil {
		log.Printf("[ERROR][usecase/flight][getReturnTime] failed time.Parse, err: %s", err.Error())
		return time.Time{}, err
	}

	returnDatetime := time.Date(
		returnDate.Year(),
		returnDate.Month(),
		returnDate.Day(),
		returnTime.Hour(),
		returnTime.Minute(),
		returnTime.Second(),
		0,
		returnTimezone,
	)

	return returnDatetime, nil
}
