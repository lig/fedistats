package fedistats

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mastodon struct {
		Server       string `yaml:"server"`
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
		AccessToken  string `yaml:"access_token"`
	} `yaml:"mastodon"`
	Metrics struct {
		Path string `yaml:"path,omitempty"`
		Port uint16 `yaml:"port,omitempty"`
	} `yaml:"metrics"`
}

func GetConfig() *Config {
	f, err := os.Open("./config/fedistats.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	config := Config{}
	config.Metrics.Path = "/metrics"
	config.Metrics.Port = 2112

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
