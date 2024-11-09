package app

import (
	"log"
	"net/http"
	"time"

	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/config"
	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/db"
	"github.com/ShaleApps/go-service-utils/helpers"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var (
	routerService = chi.NewRouter()
	routerAdmin   = chi.NewRouter()
	g             errgroup.Group
)

func StartApp(cnf config.SvcConfig) {
	servicePort := helpers.GetEnv("SERVICE_PORT", "8080")
	adminPort := helpers.GetEnv("ADMIN_PORT", "8081")

	// Establish Db connection
	db.NewEqtIntegrationDB(cnf.DynamicConfig)

	loadServiceRoutes(cnf)

	loadAdminRoutes()

	g.Go(func() error {
		logrus.Infof("Loading service routes on port: %s", servicePort)

		err := (&http.Server{
			Addr:         ":" + servicePort,
			Handler:      routerService,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		}).ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		logrus.Infof("Loading admin routes on port: %s", adminPort)
		err := (&http.Server{
			Addr:         ":" + adminPort,
			Handler:      routerAdmin,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		}).ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
