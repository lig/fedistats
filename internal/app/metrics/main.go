package metrics

import (
	"gitlab.com/lig/fedistats/internal/pkg/fedistats"
)

func Run(config *fedistats.Config) {
	listenAndServe(&Config{
		Path: config.Metrics.Path,
		Port: config.Metrics.Port,
	})
}
