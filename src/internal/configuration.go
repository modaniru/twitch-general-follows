package internal

import (
	"log"

	"github.com/BurntSushi/toml"
)

type configauration struct {
	Port string `toml:"port"`
}

func NewConfiguration(configPath string) *configauration {
	cfg := configauration{}
	_, err := toml.DecodeFile(configPath, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
