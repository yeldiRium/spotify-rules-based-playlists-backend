package version

import (
	"runtime/debug"

	"github.com/rs/zerolog/log"
)

var Version = "(version unavailable)"
var GitVersion = "(version unavailable)"

func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatal().
			Msg("failed to read build info")
	}

	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" {
			GitVersion = setting.Value
			break
		}
	}
}
