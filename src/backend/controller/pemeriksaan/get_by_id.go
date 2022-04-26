package pemeriksaan

import (
	"errors"
	"net/http"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/db"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCheckUpById(ctx *gin.Context) {
	checkupId := ctx.Params.ByName("id")

	// Input Item Validation
	if checkupId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "Disease id is required",
			"data":    nil,
		})
		return
	}

	// Getting patient
	var patient model.Pemeriksaan
	var disease model.Penyakit
	var err = db.GetDatabase().First(&patient, checkupId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "KO",
				"message": "Check up id is not found",
				"data":    err,
			})

			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "KO",
				"message": "Internal server error",
				"data":    err,
			})

			return
		}
	}

	err = db.GetDatabase().First(&disease, patient.PenyakitID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status":  "KO",
				"message": "Disease is not found",
				"data":    err,
			})

			return
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "KO",
				"message": "Internal server error",
				"data":    err,
			})

			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Success",
		"data": gin.H{
			"id":                patient.ID,
			"name":              patient.NamaPasien,
			"created_timestamp": patient.CreatedAt,
			"updated_timestamp": patient.UpdatedAt,
			"disease": gin.H{
				"id":       disease.ID,
				"name":     disease.Nama,
				"sequence": disease.DnaSequence,
			},
			"similarity": patient.Similarity,
			"sample":     patient.SequenceSample,
			"is_match":   patient.Similarity == 1.0,
		},
	})
}
