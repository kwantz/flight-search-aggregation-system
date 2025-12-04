package garudaindonesia

type FlightResponse struct {
	Status  string   `json:"status"`
	Flights []Flight `json:"flights"`
}

type Flight struct {
	FlightID       string    `json:"flight_id"`
	Airline        string    `json:"airline"`
	AirlineCode    string    `json:"airline_code"`
	Departure      Airport   `json:"departure"`
	Arrival        Airport   `json:"arrival"`
	Duration       int       `json:"duration_minutes"`
	Stops          int       `json:"stops"`
	Aircraft       string    `json:"aircraft"`
	Price          Price     `json:"price"`
	AvailableSeats int       `json:"available_seats"`
	FareClass      string    `json:"fare_class"`
	Baggage        Baggage   `json:"baggage"`
	Amenities      []string  `json:"amenities,omitempty"`
	Segments       []Segment `json:"segments,omitempty"`
}

type Airport struct {
	Airport  string `json:"airport"`
	City     string `json:"city"`
	Time     string `json:"time"`
	Terminal string `json:"terminal"`
}

type Price struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type Baggage struct {
	CarryOn int `json:"carry_on"`
	Checked int `json:"checked"`
}

type Segment struct {
	FlightNumber    string  `json:"flight_number"`
	Departure       Airport `json:"departure"`
	Arrival         Airport `json:"arrival"`
	DurationMinutes int     `json:"duration_minutes"`
	LayoverMinutes  int     `json:"layover_minutes,omitempty"`
}

var mock = `{
  "status": "success",
  "flights": [
    {
      "flight_id": "GA400",
      "airline": "Garuda Indonesia",
      "airline_code": "GA",
      "departure": {
        "airport": "CGK",
        "city": "Jakarta",
        "time": "2025-12-15T06:00:00+07:00",
        "terminal": "3"
      },
      "arrival": {
        "airport": "DPS",
        "city": "Denpasar",
        "time": "2025-12-15T08:50:00+08:00",
        "terminal": "I"
      },
      "duration_minutes": 110,
      "stops": 0,
      "aircraft": "Boeing 737-800",
      "price": {
        "amount": 1250000,
        "currency": "IDR"
      },
      "available_seats": 28,
      "fare_class": "economy",
      "baggage": {
        "carry_on": 1,
        "checked": 2
      },
      "amenities": [
        "wifi",
        "meal",
        "entertainment"
      ]
    },
    {
      "flight_id": "GA410",
      "airline": "Garuda Indonesia",
      "airline_code": "GA",
      "departure": {
        "airport": "CGK",
        "city": "Jakarta",
        "time": "2025-12-15T09:30:00+07:00",
        "terminal": "3"
      },
      "arrival": {
        "airport": "DPS",
        "city": "Denpasar",
        "time": "2025-12-15T12:25:00+08:00",
        "terminal": "I"
      },
      "duration_minutes": 115,
      "stops": 0,
      "aircraft": "Airbus A330-300",
      "price": {
        "amount": 1450000,
        "currency": "IDR"
      },
      "available_seats": 15,
      "fare_class": "economy",
      "baggage": {
        "carry_on": 1,
        "checked": 2
      },
      "amenities": [
        "wifi",
        "power_outlet",
        "meal",
        "entertainment"
      ]
    },
    {
      "flight_id": "GA315",
      "airline": "Garuda Indonesia",
      "airline_code": "GA",
      "departure": {
        "airport": "CGK",
        "city": "Jakarta",
        "time": "2025-12-15T14:00:00+07:00",
        "terminal": "3"
      },
      "arrival": {
        "airport": "SUB",
        "city": "Surabaya",
        "time": "2025-12-15T15:30:00+07:00",
        "terminal": "2"
      },
      "duration_minutes": 90,
      "stops": 0,
      "aircraft": "Boeing 737",
      "price": {
        "amount": 1850000,
        "currency": "IDR"
      },
      "segments": [
        {
          "flight_number": "GA315",
          "departure": {
            "airport": "CGK",
            "time": "2025-12-15T14:00:00+07:00"
          },
          "arrival": {
            "airport": "SUB",
            "time": "2025-12-15T15:30:00+07:00"
          },
          "duration_minutes": 90
        },
        {
          "flight_number": "GA332",
          "departure": {
            "airport": "SUB",
            "time": "2025-12-15T17:15:00+07:00"
          },
          "arrival": {
            "airport": "DPS",
            "time": "2025-12-15T18:45:00+08:00"
          },
          "duration_minutes": 90,
          "layover_minutes": 105
        }
      ],
      "available_seats": 22,
      "fare_class": "economy",
      "baggage": {
        "carry_on": 1,
        "checked": 2
      }
    }
  ]
}`
