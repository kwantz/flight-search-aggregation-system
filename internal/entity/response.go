package entity

type FlightResponse struct {
	SearchCriteria SearchCriteria         `json:"search_criteria"`
	Metadata       FlightResponseMetadata `json:"metadata"`
	Flights        []Flight               `json:"flight"`

	// Next: add departure_flights & return_flights for round-trip
}

type SearchCriteria struct {
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departure_date"`
	Passengers    int    `json:"passengers"`
	CabinClass    string `json:"cabin_class"`
}

type FlightResponseMetadata struct {
	TotalResults       int   `json:"total_results"`
	ProvidersQueried   int   `json:"providers_queried"`
	ProvidersSucceeded int   `json:"providers_succeeded"`
	ProvidersFailed    int   `json:"providers_failed"`
	SearchTimeMs       int64 `json:"search_time_ms"`
	CacheHit           bool  `json:"cache_hit"`
}
