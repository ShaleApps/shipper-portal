package app

import (
	"encoding/json"
	"net/http"

	"github.com/ShaleApps/go-service-utils/chi/http/controllers"
	utilMiddlewares "github.com/ShaleApps/go-service-utils/chi/http/middlewares"
	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/config"
	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/http/controlers/api/drivers"
	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/http/controlers/api/loads"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func loadServiceRoutes(cnf config.SvcConfig) {
	routerService.Use(middleware.Logger)
	routerService.Use(utilMiddlewares.MetricsReporter())
	routerService.Use(middleware.Recoverer)

	// Needed for the ingress health
	routerService.Get("/health", controllers.HealthCheck)
	routerService.Get("/ready", controllers.ReadyCheck)

	driversHandler := new(drivers.Handler)
	driversHandler.SvcConfig = &cnf

	loadsHandler := new(loads.Handler)
	loadsHandler.SvcConfig = &cnf

	routerService.Route("/v1", func(r chi.Router) {

		r.Route("/drivers/{driver_id}", func(r chi.Router) {
			r.Post("/email", driversHandler.EmailDriverController)
			r.Route("/loads/{load_id}", func(r chi.Router) {
				r.Get("/verify", loadsHandler.VerifyLoadDriverController)
				r.Put("/pickup", loadsHandler.PickupLoadController)
				r.Put("/dropoff", loadsHandler.DropoffLoadController)
			})
		})
	})

	routerService.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"message": "Not found",
		})
	})
}

func loadAdminRoutes() {
	routerAdmin.Use(middleware.Recoverer)

	routerAdmin.Route("/", func(r chi.Router) {
		r.Get("/metrics", promhttp.Handler().ServeHTTP)
		r.Get("/health", controllers.HealthCheck)
		r.Get("/ready", controllers.ReadyCheck)
	})

	routerService.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Not found",
		})
	})
}
