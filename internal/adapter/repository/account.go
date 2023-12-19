package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/account-api/infrastructure/postgres"
	"github.com/account-api/internal/core/models"
	"github.com/account-api/internal/core/ports"
)

type accountRepository struct {
	postgres.PostgresRepository
}

func (a *accountRepository) Create(ctx context.Context, entity models.Account) (string, error) {
	q := `
	INSERT INTO accounts (id, name, created_on)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

	acc := entity
	row := a.DB.QueryRowContext(
		ctx, q, acc.Id, acc.Name, time.Now(),
	)

	err := row.Scan(&acc.Id)
	if err != nil {
		return "", err
	}

	return acc.Id, nil
}

// Update implements ports.AccountRepository.
func (*accountRepository) Update(ctx context.Context, entity models.Account) (string, error) {
	panic("unimplemented")
}

func NewAccountRepository(db *sql.DB) ports.AccountRepository {
	return &accountRepository{
		postgres.PostgresRepository{
			DB: db,
		},
	}
}
