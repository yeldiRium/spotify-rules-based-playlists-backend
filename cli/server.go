package cli

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/yeldiRium/spotify-rules-based-playlists-backend/server"
)

var serverHTTPPortFlag int
var serverSpotifyClientIDFlag string
var serverSpotifyClientSecretFlag string
var serverSpotifyRedirectURI string
var serverFrontendURLFlag string
var serverDatabaseUrlFlag string

func init() {
	RootCommand.AddCommand(ServerCommand)
	ServerCommand.Flags().IntVar(&serverHTTPPortFlag, "http-port", 3_000, "sets the HTTP API port to listen on")
	ServerCommand.Flags().StringVar(&serverSpotifyClientIDFlag, "spotify-client-id", "", "sets the Spotify Client ID for OAuth2.0")
	ServerCommand.Flags().StringVar(&serverSpotifyClientSecretFlag, "spotify-client-secret", "", "sets the Spotify Client Secret for OAuth2.0")
	ServerCommand.Flags().StringVar(&serverSpotifyRedirectURI, "spotify-redirect-uri", "", "sets the Spotify Redirect URI for OAuth2.0")
	ServerCommand.Flags().StringVar(&serverFrontendURLFlag, "frontend-url", "", "sets the URL under which the frontend is deployed")
	ServerCommand.Flags().StringVar(&serverDatabaseUrlFlag, "database-url", "", "sets the URL for the database")
}

var ServerCommand = &cobra.Command{
	Use:   "server",
	Short: "Runs a spotify-rules-based-playlist server",
	Run: func(command *cobra.Command, args []string) {
		if serverSpotifyClientIDFlag == "" {
			log.Fatal().
				Msg("--spotify-client-id is missing")
		}
		if serverSpotifyClientSecretFlag == "" {
			log.Fatal().
				Msg("--spotify-client-secret is missing")
		}
		if serverSpotifyRedirectURI == "" {
			log.Fatal().
				Msg("--spotify-redirect-uri is missing")
		}
		if serverFrontendURLFlag == "" {
			log.Fatal().
				Msg("--frontend-url is missing")
		}
		if serverDatabaseUrlFlag == "" {
			log.Fatal().
				Msg("--database-url is missing")
		}

		pool, err := pgxpool.Connect(context.Background(), serverDatabaseUrlFlag)
		if err != nil {
			log.
				Fatal().
				Err(err).
				Msg("could not establish a connection to the database")
		}
		defer pool.Close()

		server.Run(
			serverHTTPPortFlag,
			rootVerboseFlag,
			serverSpotifyClientIDFlag,
			serverSpotifyClientSecretFlag,
			serverSpotifyRedirectURI,
			serverFrontendURLFlag,
			pool,
		)
	},
}
