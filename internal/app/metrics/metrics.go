package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	UpdatesCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "updates_total",
		Help: "The total number of update events",
	})
	DeletesCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "deletes_total",
		Help: "The total number of delete events",
	})
)
