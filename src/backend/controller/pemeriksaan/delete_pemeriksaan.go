package pemeriksaan

import (
	"net/http"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/db"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/model"
	"github.com/gin-gonic/gin"
)

func DeleteCheckupHistory(ctx *gin.Context) {
	checkId := ctx.Params.ByName("id")

	// Input Item Validation
	if checkId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "Check up history id is required",
			"data":    nil,
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
			"id":                checkupHistory.ID,
			"name":              checkupHistory.NamaPasien,
			"similarity":        checkupHistory.Similarity,
			"disease_id":        checkupHistory.PenyakitID,
			"disease_name":      checkupHistory.Penyakit.Nama,
			"is_match":          checkupHistory.Similarity == 1.0,
			"created_timestamp": checkupHistory.CreatedAt,
			"updated_timestamp": checkupHistory.UpdatedAt,
		},
	})
}
