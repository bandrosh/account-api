package ports

import (
	"context"

	"github.com/account-api/internal/core/models"
)

type AccountService interface {
	UpdateAccount(context.Context, string, models.Account) error
	CreateAccount(context.Context, models.Account) error
}

type AccountRepository interface {
	Create(ctx context.Context, entity models.Account) (string, error)
}
