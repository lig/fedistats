package collector

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mattn/go-mastodon"
)

func runOnGetAPI(config *mastodon.Config) {
	c := mastodon.NewClient(config)
	timeline, err := c.GetTimelinePublic(
		context.Background(),
		true,
		&mastodon.Pagination{MaxID: "", SinceID: "", MinID: "", Limit: 1},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(timeline[0])
	var minID = timeline[0].ID

	for {
		time.Sleep(2 * time.Second)
		println("*** Requesting 20 more")

		timeline, err = c.GetTimelinePublic(
			context.Background(),
			true,
			&mastodon.Pagination{MaxID: "", SinceID: "", MinID: minID, Limit: 20},
		)
		if err != nil {
			log.Fatal(err)
		}
		println("*** Got", len(timeline), "more")

		for i := len(timeline) - 1; i >= 0; i-- {
			fmt.Println(timeline[i])
			minID = timeline[i].ID
		}
	}
}
