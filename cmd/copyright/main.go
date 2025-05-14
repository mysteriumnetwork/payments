package main

import (
	"github.com/mysteriumnetwork/go-ci/commands"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := commands.CopyrightD(".", "pb"); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
