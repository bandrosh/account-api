package middlewares

import (
	"github.com/gin-gonic/gin"

	"github.com/account-api/infrastructure/logger"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info(logger.ServerInfo, "DOING SOMETHING ON MIDDLEWARE")

		// before request
		c.Next()
	}
}
