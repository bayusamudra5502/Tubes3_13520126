package penyakit

import (
	"errors"
	"net/http"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/db"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDiseaseById(ctx *gin.Context) {
	diseaseId := ctx.Params.ByName("id")

	// Input Item Validation
	if diseaseId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "Disease id is required",
			"data":    nil,
		})
		return
	}

	// Getting disease
	var disease model.Penyakit
	var err = db.GetDatabase().First(&disease, diseaseId).Error

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
			"id":                disease.ID,
			"name":              disease.Nama,
			"sequence":          disease.DnaSequence,
			"created_timestamp": disease.CreatedAt,
			"updated_timestamp": disease.UpdatedAt,
		},
	})
}
