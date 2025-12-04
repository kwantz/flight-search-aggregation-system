package main

import (
	"log"
	"net/http"

	"github.com/kwantz/flight-search-aggregation-system/internal/domain/flight/airasia"
	"github.com/kwantz/flight-search-aggregation-system/internal/domain/flight/batikair"
	"github.com/kwantz/flight-search-aggregation-system/internal/domain/flight/garudaindonesia"
	"github.com/kwantz/flight-search-aggregation-system/internal/domain/flight/lionair"
	"github.com/kwantz/flight-search-aggregation-system/internal/interfaces"
	"github.com/kwantz/flight-search-aggregation-system/internal/usecase/flight"

	handlerHttp "github.com/kwantz/flight-search-aggregation-system/internal/handler/http"
)

func main() {
	domainAirAsia := airasia.New()
	domainBatikAir := batikair.New()
	domainGarudaID := garudaindonesia.New()
	domainLionAir := lionair.New()

	flightProviders := []interfaces.IFlightProviderDomain{
		domainAirAsia,
		domainBatikAir,
		domainGarudaID,
		domainLionAir,
	}

	usecaseFlight := flight.New(flightProviders)

	handler := handlerHttp.New(usecaseFlight)
	router := handler.Register()

	log.Printf("[INFO] server run in port :3000")
	http.ListenAndServe(":3000", router)
}
