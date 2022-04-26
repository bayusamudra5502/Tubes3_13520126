package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginationContext(ctx *gin.Context) (int, int) {
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "50"))

	if err != nil || limit <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "Limit is not valid",
			"data":    err,
		})
		return -1, -1
	}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "0"))

	if err != nil || page < 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "Page is not valid",
			"data":    err,
		})
		return -1, -1
	}

	return page, limit
}
