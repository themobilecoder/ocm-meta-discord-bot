package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/themobilecoder/ocm-meta/meta"
)

type Config struct {
	DISCORD_API_KEY string `envconfig:"DISCORD_API_KEY" required:"true"`
	OCM_GUILD_ID    string `envconfig:"OCM_GUILD_ID" required:"true"`
}

func main() {
	var cfg Config
	err := envconfig.Process("METABOT", &cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	monkeys := meta.GetOnChainMonkeys()
	fmt.Print(monkeys[4642-1])
}
