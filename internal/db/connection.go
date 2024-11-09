package db

import (
	"context"
	"database/sql"
	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/dynamic_config"
	"github.com/dlmiddlecote/sqlstats"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var SQLXDB *sqlx.DB

type Sqlxer interface {
	sqlx.Queryer
	sqlx.QueryerContext
	sqlx.Execer
	sqlx.ExecerContext
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
}

func NewEqtIntegrationDB(config dynamic_config.DynamicConfig) {
	var (
		db          *sqlx.DB
		err         error
		maxAttempts = 3
	)

	for i := 0; i < maxAttempts; i++ {
		db, err = sqlx.Connect("postgres", config.GetAgnusDBConnString())
		if err != nil {
			logrus.Error("unable to connect to postgres DB ", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		db.SetMaxOpenConns(config.GetAgnusDBMaxOpenConnection())
		db.SetMaxIdleConns(config.GetAgnusDBMaxIdleConnection())

		SQLXDB = db
		break
	}

	if err != nil {
		log.Fatal(err)
	}

	collector := sqlstats.NewStatsCollector("{{SERVICE_NAME_SNAKE_CASE}}", db)
	prometheus.MustRegister(collector)
}
