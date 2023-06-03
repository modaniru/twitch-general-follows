package internal

import (
	"log"

	"github.com/BurntSushi/toml"
)

type configauration struct {
	Port         int    `toml:"port"`
	ClientId     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
}

func NewConfiguration(configPath string) *configauration {
	cfg := configauration{}
	_, err := toml.DecodeFile(configPath, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
