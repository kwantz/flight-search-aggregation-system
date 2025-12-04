package lionair

type FlightResponse struct {
	Success bool `json:"success"`
	Data    Data `json:"data"`
}

type Data struct {
	AvailableFlights []Flight `json:"available_flights"`
}

type Flight struct {
	ID         string    `json:"id"`
	Carrier    Carrier   `json:"carrier"`
	Route      Route     `json:"route"`
	Schedule   Schedule  `json:"schedule"`
	FlightTime int       `json:"flight_time"`
	IsDirect   bool      `json:"is_direct"`
	StopCount  int       `json:"stop_count,omitempty"`
	Layovers   []Layover `json:"layovers,omitempty"`
	Pricing    Pricing   `json:"pricing"`
	SeatsLeft  int       `json:"seats_left"`
	PlaneType  string    `json:"plane_type"`
	Services   Services  `json:"services"`
}

type Carrier struct {
	Name string `json:"name"`
	IATA string `json:"iata"`
}

type Route struct {
	From Airport `json:"from"`
	To   Airport `json:"to"`
}

type Airport struct {
	Code string `json:"code"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Schedule struct {
	Departure         string `json:"departure"`
	DepartureTimezone string `json:"departure_timezone"`
	Arrival           string `json:"arrival"`
	ArrivalTimezone   string `json:"arrival_timezone"`
}

type Pricing struct {
	Total    int64  `json:"total"`
	Currency string `json:"currency"`
	FareType string `json:"fare_type"`
}

type Services struct {
	WifiAvailable    bool             `json:"wifi_available"`
	MealsIncluded    bool             `json:"meals_included"`
	BaggageAllowance BaggageAllowance `json:"baggage_allowance"`
}

type BaggageAllowance struct {
	Cabin string `json:"cabin"`
	Hold  string `json:"hold"`
}

type Layover struct {
	Airport         string `json:"airport"`
	DurationMinutes int    `json:"duration_minutes"`
}

var mock = `{
  "success": true,
  "data": {
    "available_flights": [
      {
        "id": "JT740",
        "carrier": {
          "name": "Lion Air",
          "iata": "JT"
        },
        "route": {
          "from": {
            "code": "CGK",
            "name": "Soekarno-Hatta International",
            "city": "Jakarta"
          },
          "to": {
            "code": "DPS",
            "name": "Ngurah Rai International",
            "city": "Denpasar"
          }
        },
        "schedule": {
          "departure": "2025-12-15T05:30:00",
          "departure_timezone": "Asia/Jakarta",
          "arrival": "2025-12-15T08:15:00",
          "arrival_timezone": "Asia/Makassar"
        },
        "flight_time": 105,
        "is_direct": true,
        "pricing": {
          "total": 950000,
          "currency": "IDR",
          "fare_type": "ECONOMY"
        },
        "seats_left": 45,
        "plane_type": "Boeing 737-900ER",
        "services": {
          "wifi_available": false,
          "meals_included": false,
          "baggage_allowance": {
            "cabin": "7 kg",
            "hold": "20 kg"
          }
        }
      },
      {
        "id": "JT742",
        "carrier": {
          "name": "Lion Air",
          "iata": "JT"
        },
        "route": {
          "from": {
            "code": "CGK",
            "name": "Soekarno-Hatta International",
            "city": "Jakarta"
          },
          "to": {
            "code": "DPS",
            "name": "Ngurah Rai International",
            "city": "Denpasar"
          }
        },
        "schedule": {
          "departure": "2025-12-15T11:45:00",
          "departure_timezone": "Asia/Jakarta",
          "arrival": "2025-12-15T14:35:00",
          "arrival_timezone": "Asia/Makassar"
        },
        "flight_time": 110,
        "is_direct": true,
        "pricing": {
          "total": 890000,
          "currency": "IDR",
          "fare_type": "ECONOMY"
        },
        "seats_left": 38,
        "plane_type": "Boeing 737-800",
        "services": {
          "wifi_available": false,
          "meals_included": false,
          "baggage_allowance": {
            "cabin": "7 kg",
            "hold": "20 kg"
          }
        }
      },
      {
        "id": "JT650",
        "carrier": {
          "name": "Lion Air",
          "iata": "JT"
        },
        "route": {
          "from": {
            "code": "CGK",
            "name": "Soekarno-Hatta International",
            "city": "Jakarta"
          },
          "to": {
            "code": "DPS",
            "name": "Ngurah Rai International",
            "city": "Denpasar"
          }
        },
        "schedule": {
          "departure": "2025-12-15T16:20:00",
          "departure_timezone": "Asia/Jakarta",
          "arrival": "2025-12-15T21:10:00",
          "arrival_timezone": "Asia/Makassar"
        },
        "flight_time": 230,
        "is_direct": false,
        "stop_count": 1,
        "layovers": [
          {
            "airport": "SUB",
            "duration_minutes": 75
          }
        ],
        "pricing": {
          "total": 780000,
          "currency": "IDR",
          "fare_type": "ECONOMY"
        },
        "seats_left": 52,
        "plane_type": "Boeing 737-800",
        "services": {
          "wifi_available": false,
          "meals_included": false,
          "baggage_allowance": {
            "cabin": "7 kg",
            "hold": "20 kg"
          }
        }
      }
    ]
  }
}`
