package postgres

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"

	"github.com/account-api/infrastructure/logger"
	"github.com/pressly/goose/v3"
)

func ConnectPostgresDB(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error(logger.ConfigError, err)

		return nil, err
	}

	return db, pingSql(ctx, db)
}

func pingSql(ctx context.Context, db *sql.DB) (err error) {
	// wait until db is ready
	for start := time.Now(); time.Since(start) < (5 * time.Second); {
		err = db.PingContext(ctx)
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}
	return err
}

func MigratePostgresDB(db *sql.DB, migrationsDir string) error {
	goose.SetTableName("public.goose_db_version")
	return goose.Up(db, migrationsDir)
}

type PostgresRepository struct {
	DB *sql.DB
}
