package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jean0206/customer-api-golang/Services"
)

func GetRoutes() http.Handler {

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	go r.Mount("/data", dataRoute())
	go r.Mount("/customer", customerRoute())
	go r.Mount("/product", productRoute())
	go r.Mount("/transaction", productTransaction())

	return r

}

func dataRoute() http.Handler {
	r := chi.NewRouter()
	r.Get("/", Services.ImportData)
	return r
}

func customerRoute() http.Handler {
	r := chi.NewRouter()
	r.Get("/", Services.GetCustomers)
	r.Get("/ip/{ip}", Services.GetCustomersWithIp)
	r.Get("/{userId}", Services.GetCustomerById)
	return r
}

func productRoute() http.Handler {
	r := chi.NewRouter()
	r.Get("/", Services.GetAllProducts)
	r.Get("/{productId}", Services.GetProductById)
	r.Get("/suggestion/{min}/{max}", Services.GetSuggestions)
	return r
}

func productTransaction() http.Handler {
	r := chi.NewRouter()
	r.Get("/", Services.GetAllTransactions)
	r.Get("/{userId}", Services.GetCustomerTransactionById)
	return r
}
