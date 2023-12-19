package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/account-api/config"
	"github.com/account-api/infrastructure/logger"
	"github.com/account-api/internal/core/models"
	"github.com/account-api/internal/core/ports"
)

func SetAccountRoutes(ctx context.Context, config config.Configuration, r *gin.Engine, s ports.AccountService) {
	r.PUT("/account", updateAccount(ctx, config, s))
	r.POST("/account", createAccount(ctx, config, s))
}

func createAccount(ctx context.Context, cfg config.Configuration, s ports.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var account models.Account

		if err := c.ShouldBindJSON(&account); err != nil {
			logger.Error(logger.HTTPError, "cannot marshal body")
			return
		}

		if err := s.CreateAccount(ctx, account); err != nil {
			logger.Error(logger.HTTPError, "Error creating account")
			c.JSON(http.StatusUnprocessableEntity, map[string]string{
				"error":  "failed create account",
				"reason": err.Error(),
			})
			return
		}

		logger.Info(logger.ServerInfo, fmt.Sprintf("Account Created %v", account))

		c.JSON(http.StatusCreated, map[string]string{"success": "created"})
	}
}

func updateAccount(ctx context.Context, cfg config.Configuration, s ports.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var account models.Account

		if err := c.ShouldBindJSON(&account); err != nil {
			logger.Error(logger.HTTPError, "cannot marshal body")
			return
		}

		if err := s.UpdateAccount(ctx, account); err != nil {
			logger.Error(logger.HTTPError, "Error creating account")
			c.JSON(http.StatusUnprocessableEntity, map[string]string{
				"error":  "failed update account",
				"reason": err.Error(),
			})
			return
		}

		logger.Info(logger.ServerInfo, fmt.Sprintf("Account Updated %v", account))

		c.JSON(http.StatusOK, map[string]string{"success": "updated"})
	}
}
