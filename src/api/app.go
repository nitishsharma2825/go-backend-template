package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	// Add some basic middleware to the router
	r.Use(middleware.RequestID) // Assigns a request ID to each request
	r.Use(middleware.RealIP)    // Sets the RemoteAddr to the value of X-Real-IP or X-Forwarded-For
	r.Use(middleware.Logger)    // Logs the start and end of each request with the elapsed processing time
	r.Use(middleware.Recoverer) // Recover from panics without crashing server

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux *chi.Mux) error {
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute * 2,
	}

	return server.ListenAndServe()
}
