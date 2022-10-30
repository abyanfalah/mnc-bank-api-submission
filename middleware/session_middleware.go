package middleware

import (
	response "mnc-bank-api/utils/common_response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerId, err := ctx.Cookie("session")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "unable to get session (middleware)",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if customerId == "" {
			response.JsonUnauthorized(ctx, "Unauthorized")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}
