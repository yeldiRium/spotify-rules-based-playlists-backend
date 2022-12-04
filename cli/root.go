package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootVersionFlag bool
var rootVerboseFlag bool

func init() {
	RootCommand.Flags().BoolVarP(&rootVersionFlag, "version", "v", false, "prints the version")
	RootCommand.PersistentFlags().BoolVarP(&rootVerboseFlag, "verbose", "", false, "enables verbose mode")
}

var RootCommand = &cobra.Command{
	Use:   "spotify-rules-backed-playlists-backend",
	Short: "A services that generates spotify playlists based on rules",
	Long:  "A services that generates spotify playlists based on rules.",
	Run: func(command *cobra.Command, args []string) {
		if rootVersionFlag {
			VersionCommand.Run(command, args)
			return
		}

		fmt.Println(command.UsageString())
	},
}
