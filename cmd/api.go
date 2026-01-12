package main

import (
	"log"
	"net/http"
	"time"

	"github.com/amdaaraijen/Learn-API/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *api) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)

	r.Use(middleware.Timeout(60*time.Second))
	
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})


	productService := products.NewService()
	productHandler := products.NewHandler(productService)

	r.Get("/products", productHandler.ListProducts)

	return r
}

func (app *api) run(h http.Handler) error {
	svr := http.Server{
	Addr: app.config.addr,
	Handler: h,
	ReadTimeout:  10 * time.Second,
	WriteTimeout: 10 * time.Second,
	IdleTimeout:  120 * time.Second,
	}

	log.Printf("Application has started at addr %s", app.config.addr)

	return svr.ListenAndServe()
}

type api struct {
	config config
}

type config struct {
	addr string
}
