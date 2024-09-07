package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(time.Minute))

	fileServe := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServe))

	// display our test page
	mux.Get("/test-patterns", app.TestPatterns)

	// factory routes
	mux.Get("/dog-from-factory", app.CreateDogFromFactory)
	mux.Get("/cat-from-factory", app.CreateCatFromFactory)

	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)

	return mux
}
