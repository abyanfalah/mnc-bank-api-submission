package middleware

import (
	"mnc-bank-api/utils"

	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerId, err := ctx.Cookie("session")
		if err != nil {
			utils.JsonErrorInternalServerError(ctx, err, "error getting session")
			return
		}

		if customerId == "" {
			utils.JsonUnauthorized(ctx, "Unauthorized")
			return
		}
	}
}
