package pemeriksaan

import (
	"net/http"
	"strconv"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/db"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/model"
	"github.com/gin-gonic/gin"
)

func DeleteCheckupHistory(ctx *gin.Context) {
	checkId, err := strconv.Atoi(ctx.Params.ByName("id"))

	// Input Item Validation
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "Check up history id is required and must be a number",
			"data":    err,
		})
		return
	}

	// Getting disease
	var checkupHistory model.Pemeriksaan
	var del = db.GetDatabase().Delete(&checkupHistory, checkId)

	if del.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "KO",
			"message": "Internal server error",
			"data":    del.Error,
		})

		return
	}

	if del.RowsAffected < 1 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "Disease is not found",
			"data":    nil,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Success",
		"data": gin.H{
			"id": checkId,
		},
	})
}
