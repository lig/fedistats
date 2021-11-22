package collector

import (
	"context"
	"log"

	"github.com/mattn/go-mastodon"

	"gitlab.com/lig/fedistats/internal/app/metrics"
)

func listenTimeline(config *mastodon.Config) {
	c := mastodon.NewClient(config)
	pubChannel, err := c.StreamingPublic(context.Background(), true)
	if err != nil {
		log.Fatal(err)
	}

	for pubEvent := range pubChannel {
		switch event := pubEvent.(type) {
		case *mastodon.UpdateEvent:
			metrics.UpdatesCount.Inc()
			print(".")
		case *mastodon.DeleteEvent:
			metrics.DeletesCount.Inc()
			print("x")
		case *mastodon.ErrorEvent:
			log.Println(event.Error())
		case *mastodon.NotificationEvent:
		default:
		}
	}
}
