package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func listenAndServe(config *Config) {
	http.Handle(config.Path, promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
