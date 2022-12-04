package server

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

func Run(
	httpPort int,
	verbose bool,
	spotifyClientID string,
	spotifyClientSecret string,
	frontendUrl string,
	pool *pgxpool.Pool,
) {

}
