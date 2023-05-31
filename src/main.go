package main

import (
	"fmt"

	"github.com/modaniru/twitch-general-follows/src/internal"
)

func main() {
	cfg := internal.NewConfiguration("config.toml")
	fmt.Printf("%+v", *cfg)
}
