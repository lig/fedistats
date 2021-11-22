package collector

import (
	"github.com/mattn/go-mastodon"

	"gitlab.com/lig/fedistats/internal/pkg/fedistats"
)

func Run(config *fedistats.Config) {
	listenTimeline(&mastodon.Config{
		Server:       config.Mastodon.Server,
		ClientID:     config.Mastodon.ClientID,
		ClientSecret: config.Mastodon.ClientSecret,
		AccessToken:  config.Mastodon.AccessToken,
	})
}
