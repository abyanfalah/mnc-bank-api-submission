package middleware

import (
	response "mnc-bank-api/utils/common_response"

	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerId, err := ctx.Cookie("session")
		if err != nil {
			response.JsonErrorInternalServerError(ctx, err, "error getting session")
			return
		}

		if customerId == "" {
			response.JsonUnauthorized(ctx, "Unauthorized")
			return
		}

	}
}
