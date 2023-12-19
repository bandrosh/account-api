package services

import (
	"context"
	"fmt"

	"github.com/account-api/config"
	"github.com/account-api/infrastructure/logger"
	"github.com/account-api/internal/core/models"
	"github.com/account-api/internal/core/ports"
)

type AccountService struct {
	config            config.Configuration
	accountRepository ports.AccountRepository
}

func (a *AccountService) CreateAccount(ctx context.Context, account models.Account) error {
	_, err := a.accountRepository.Create(ctx, account)
	if err != nil {
		logger.Error(logger.HTTPError, fmt.Sprintf("cannot create account error: %v", err.Error()))
		return err
	}

	return nil
}

func (a *AccountService) UpdateAccount(ctx context.Context, account models.Account) error {
	_, err := a.accountRepository.Update(ctx, account)
	if err != nil {
		logger.Error(logger.HTTPError, fmt.Sprintf("cannot update account error: %v", err.Error()))
		return err
	}

	return nil
}

func NewAccountService(config config.Configuration, accountRepository ports.AccountRepository) ports.AccountService {
	return &AccountService{config: config, accountRepository: accountRepository}
}
