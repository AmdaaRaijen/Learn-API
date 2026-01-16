package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/amdaaraijen/Learn-API/internal/adapters/pgsql/sqlc"
	"github.com/amdaaraijen/Learn-API/internal/orders"
	"github.com/amdaaraijen/Learn-API/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
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

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)

	r.Route("/products", func(r chi.Router) {
		r.Get("/", productHandler.ListProducts)
		r.Get("/{id}", productHandler.GetProductById)
	})

	orderService := orders.NewService(repo.New(app.db), app.db)
	orderHandler := orders.NewHandler(orderService)

	r.Post("/orders", orderHandler.PlaceOrder)

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
	db *pgx.Conn
}

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	dsn string
}
