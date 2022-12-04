package main

import (
	"os"
	"strconv"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/yeldiRium/spotify-rules-based-playlists-backend/cli"
)

func main() {
	configureLogging()

	if strconv.IntSize != 64 {
		log.Fatal().
			Msg("64-bit architecture required")
	}

	err := cli.RootCommand.Execute()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to execute root command")
	}
}

func configureLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if isatty.IsTerminal(os.Stderr.Fd()) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}
