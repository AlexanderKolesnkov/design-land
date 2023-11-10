package main

import (
	"github.com/AlexanderKolesnkov/disign-land/pkg/config"
	"github.com/AlexanderKolesnkov/disign-land/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	return mux
}
