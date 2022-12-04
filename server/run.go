package server

import (
	"net/http"
	"strconv"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/yeldiRium/spotify-rules-based-playlists-backend/server/api/health"
)

func Run(
	httpPort int,
	verbose bool,
	spotifyClientID string,
	spotifyClientSecret string,
	frontendUrl string,
	pool *pgxpool.Pool,
) {
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)

		log.Info().
			Msg("verbose mode enabled")
	}

	sessionManager := scs.New()
	sessionManager.Store = pgxstore.New(pool)

	router := newRouter(
		spotifyClientID,
		spotifyClientSecret,
		frontendUrl,
		pool,
		sessionManager,
	)

	log.Info().
		Int("port", httpPort).
		Msg("starting HTTP API...")
	httpPortAsString := strconv.Itoa(httpPort)
	httpServer := &http.Server{
		Addr:    ":" + httpPortAsString,
		Handler: sessionManager.LoadAndSave(router),
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("server failed")
	}
}

func newRouter(
	spotifyClientID string,
	spotifyClientSecret string,
	frontendUrl string,
	pool *pgxpool.Pool,
	sessionManager *scs.SessionManager,
) *httprouter.Router {
	router := httprouter.New()

	router.Handler(
		"GET",
		"/health",
		health.Handle(),
	)

	return router
}
