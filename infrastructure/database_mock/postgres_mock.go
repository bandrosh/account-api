package databasemock

import (
	"context"
	"database/sql"
)

type PostgresRepositoryMock interface {
	ConnectPostgresDB(context.Context, string) (*sql.DB, error)
}
