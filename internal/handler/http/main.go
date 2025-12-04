package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kwantz/flight-search-aggregation-system/internal/interfaces"
)

type handler struct {
	ucFlight interfaces.IFlightUsecase
}

func New(ucFlight interfaces.IFlightUsecase) *handler {
	return &handler{
		ucFlight: ucFlight,
	}
}

func (h *handler) Register() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/flight/search", h.PostFlightSearch)
	})

	return r
}
