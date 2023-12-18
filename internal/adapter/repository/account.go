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

// Create implements ports.AccountRepository.
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

// Delete implements ports.AccountRepository.
func (a *accountRepository) Delete(ctx context.Context, ID string) error {
	panic("unimplemented")
}

// Get implements ports.AccountRepository.
func (a *accountRepository) Get(ctx context.Context, filter map[string]interface{}, skip *int, take *int) ([]interface{}, error) {
	panic("unimplemented")
}

// GetByID implements ports.AccountRepository.
func (a *accountRepository) GetByID(ctx context.Context, ID string) (interface{}, error) {
	panic("unimplemented")
}

// Update implements ports.AccountRepository.
func (a *accountRepository) Update(ctx context.Context, ID string, entity interface{}) error {
	panic("unimplemented")
}

func NewAccountRepository(db *sql.DB) ports.AccountRepository {
	return &accountRepository{
		postgres.PostgresRepository{
			DB: db,
		},
	}
}
