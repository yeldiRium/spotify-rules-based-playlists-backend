package store

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Setup(ctx context.Context, pool *pgxpool.Pool) error {
	rows, err := pool.Query(
		context.Background(),
		`
CREATE TABLE IF NOT EXISTS sessions (
	token TEXT PRIMARY KEY,
	data BYTEA NOT NULL,
	expiry TIMESTAMPTZ NOT NULL
);
`,
	)
	if err != nil {
		return errors.New("could not create sessions table")
	}
	defer rows.Close()

	rows, err = pool.Query(
		ctx,
		`
CREATE INDEX IF NOT EXISTS sessions_expiry_idx ON sessions (expiry);
`,
	)
	if err != nil {
		return errors.New("could not create index on expiry field on table sessions")
	}
	defer rows.Close()

	return nil
}
