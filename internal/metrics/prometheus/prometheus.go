package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var metricNamespace = "{{SERVICE_NAME_SNAKE_CASE}}"

type Collectors struct {
	DriverDetailsUpdatedCount prometheus.CounterVec
}

func NewMetricsCollector() Collectors {
	return Collectors{
		DriverDetailsUpdatedCount: *prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: metricNamespace,
			Name:      "handled_messages_count",
			Help:      "The total number of driver updates handled",
		}, []string{"some_useful_label"}),
	}
}

func (c Collectors) Register() {
	prometheus.MustRegister(c.DriverDetailsUpdatedCount)
}
