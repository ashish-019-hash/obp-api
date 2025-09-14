package utils

import (
	"github.com/gin-gonic/gin"
)

func SendJSONResponse(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, data)
}

func SendErrorResponse(ctx *gin.Context, statusCode int, message string, details string) {
	response := gin.H{
		"error": gin.H{
			"message": message,
			"code":    statusCode,
		},
	}
	
	if details != "" {
		response["error"].(gin.H)["details"] = details
	}
	
	ctx.JSON(statusCode, response)
}
