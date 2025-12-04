package entity

import "time"

type FlightRequest struct {
	// To Search
	Origin        string  `json:"origin"`
	Destination   string  `json:"destination"`
	DepartureDate string  `json:"departure_date"`
	ReturnDate    *string `json:"return_date"` // Using a pointer to handle null values
	Passengers    int     `json:"passengers"`
	CabinClass    string  `json:"cabin_class"`

	// To Filter
	Airline           string  `json:"Airline"`
	DepartureTime     string  `json:"departure_time"`
	DepartureTimezone string  `json:"departure_timezone"`
	ReturnTime        *string `json:"return_time"`
	ReturnTimezone    *string `json:"return_timezone"`
	PriceFrom         int64   `json:"price_from"`
	PriceTo           int64   `json:"price_to"`
	DurationFrom      int     `json:"duration_from"` // in minute
	DurationTo        int     `json:"duration_to"`   // in minute
	Stops             int     `json:"stops"`

	SortBy FlightSortRequest `json:"sort_by"`

	Departure time.Time `json:"-"`
	Return    time.Time `json:"-"`
}

type FlightSortRequest struct {
	Price     string `json:"price"`     // asc or desc
	Duration  string `json:"duration"`  // asc or desc
	Arrival   string `json:"arrival"`   // asc or desc
	Departure string `json:"departure"` // asc or desc
}
