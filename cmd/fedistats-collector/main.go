package main

import (
	"gitlab.com/lig/fedistats/internal/app/collector"
	"gitlab.com/lig/fedistats/internal/app/metrics"
	"gitlab.com/lig/fedistats/internal/pkg/fedistats"
)

func main() {
	config := fedistats.GetConfig()

	go metrics.Run(config)
	collector.Run(config)
}
