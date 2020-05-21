package handlers

import (
	"github.com/eugeneradionov/stocks/api/handlers/candles"
	"github.com/eugeneradionov/stocks/api/handlers/middlewares"
	"github.com/eugeneradionov/stocks/api/handlers/symbols"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middlewares.RequestID)

	r.Route("/symbols", func(r chi.Router) {
		r.Get("/", symbols.GetList)
		r.Get("/{symbolName}", symbols.GetByName)

		r.Post("/{symbolName}/candles", candles.Get)
	})

	return r
}
