package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/account-api/internal/core/models"
	"github.com/account-api/internal/core/ports"
)

type PostgresRepositoryMock struct {
	mem map[string]models.Account
}

func (p *PostgresRepositoryMock) Create(ctx context.Context, account models.Account) (string, error) {
	if _, ok := p.mem[account.Id]; !ok {
		p.mem[account.Id] = account
		return account.Id, nil
	}

	return "", errors.New("duplicated Key violation")
}

func (p *PostgresRepositoryMock) Update(ctx context.Context, account models.Account) (string, error) {
	if _, ok := p.mem[account.Id]; !ok {
		return "", errors.New(fmt.Sprintf("account %s does not exist", account.Id))
	}

	p.mem[account.Id] = account
	return "", nil
}

func NewAccountMockRepository() ports.AccountRepository {
	return &PostgresRepositoryMock{
		mem: make(map[string]models.Account),
	}
}
