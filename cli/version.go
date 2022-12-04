package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yeldiRium/spotify-rules-based-playlists-backend/version"
)

func init() {
	RootCommand.AddCommand(VersionCommand)
}

var VersionCommand = &cobra.Command{
	Use:   "version",
	Short: "Prints the spotify-rules-based-playlists-backend's version",
	Long:  "Prints the spotify-rules-based-playlists-backend's version.",
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("Spotify-rules-based-playlists-backend " + version.Version)
		fmt.Println("Revision " + version.GitVersion)
		fmt.Println()
	},
}
