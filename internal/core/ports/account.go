package ports

import (
	"context"

	"github.com/account-api/internal/core/models"
)

type AccountService interface {
	UpdateAccount(context.Context, models.Account) error
	CreateAccount(context.Context, models.Account) error
}

type AccountRepository interface {
	Create(ctx context.Context, entity models.Account) (string, error)
	Update(ctx context.Context, entity models.Account) (string, error)
}
