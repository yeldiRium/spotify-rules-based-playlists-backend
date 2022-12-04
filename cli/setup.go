package cli

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/yeldiRium/spotify-rules-based-playlists-backend/store"
)

var setupDatabaseUrlFlag string

func init() {
	RootCommand.AddCommand(SetupCommand)
	SetupCommand.Flags().StringVar(&setupDatabaseUrlFlag, "database-url", "", "sets the URL for the database")
}

var SetupCommand = &cobra.Command{
	Use:   "setup",
	Short: "Sets up the database for the spotify-rules-based-playlists-backend",
	Run: func(command *cobra.Command, args []string) {
		if setupDatabaseUrlFlag == "" {
			log.Fatal().
				Msg("--database-url is missing")
		}

		pool, err := pgxpool.Connect(context.Background(), setupDatabaseUrlFlag)
		if err != nil {
			log.
				Fatal().
				Err(err).
				Msg("could not establish a connection to the database")
		}
		defer pool.Close()

		err = store.Setup(context.Background(), pool)
		if err != nil {
			log.
				Fatal().
				Err(err).
				Msg("could not create session table and index")
		}
	},
}
