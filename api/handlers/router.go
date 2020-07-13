package handlers

import (
	"github.com/eugeneradionov/stocks/api/handlers/auth"
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

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/auth", func(r chi.Router) {
				r.Post("/register", auth.Register)
				r.Post("/login", auth.Login)

				r.Group(func(r chi.Router) {
					r.Use(middlewares.Auth)

					r.Post("/refresh-token", auth.RefreshToken)
					r.Delete("/logout", auth.Logout)
				})
			})

			r.Route("/symbols", func(r chi.Router) {
				r.Use(middlewares.Auth)

				r.Get("/", symbols.GetList)
				r.Get("/{symbolName}", symbols.GetByName)

				r.Post("/{symbolName}/candles", candles.Get)
			})
		})
	})

	return r
}
