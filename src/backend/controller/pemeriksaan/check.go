package pemeriksaan

import (
	"errors"
	"net/http"
	"time"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/controller"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/db"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/lib"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckDisease(ctx *gin.Context){
	file, err := ctx.FormFile("file")
	name := ctx.PostForm("name")
	diseaseId := ctx.PostForm("disease_id")
	algorithm := ctx.DefaultPostForm("algorithm", "KMP")

	// Input Item Validation
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "KO",
			"message": "File is required",
			"data": err,
		})
		return
	}

	if name == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "KO",
			"message": "User name is required",
			"data": err,
		})
		return
	}

	if algorithm != "KMP" && algorithm != "BOYER" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "KO",
			"message": "Algorithm is not valid",
			"data": err,
		})
		return
	}

	if diseaseId == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "KO",
			"message": "Disease id is required",
			"data": err,
		})
		return
	}

	// File Validation
	sequence, isValid := controller.FileValidation(ctx, file)

	if !isValid {
		return
	}
	
	// Getting disease
	var disease model.Penyakit
	err = db.GetDatabase().First(&disease, diseaseId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": "KO",
				"message": "Disease is not found",
				"data": err,
			})

			return
		}else{
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": "KO",
				"message": "Internal server error",
				"data": err,
			})

			return
		}
	}

	var similarity float64
	var foundIndex int
	var start time.Time = time.Now()

	// Checking
	if algorithm == "KMP" {
		foundIndex = lib.Kmp(sequence, disease.DnaSequence)
	} else {
		// Boyer-moore algorithm
		foundIndex = lib.BoyerMoore(sequence, disease.DnaSequence)
	}

	if foundIndex == -1 {
		similarity = lib.Similarity(sequence, disease.DnaSequence)
	} else {
		similarity = 1.0
	}

	var stop time.Time = time.Now()
	var duration = stop.Nanosecond() - start.Nanosecond()

	data := model.Pemeriksaan{
		NamaPasien: name,
		SequenceSample: sequence,
		PenyakitID: disease.ID,
		Similarity: similarity,
	}

	err = db.GetDatabase().Create(&data).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "KO",
			"message": "Error inserting to database",
			"data": err,
		})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"message": "Success",
		"data": gin.H{
			"duration": duration,
			"id": data.ID,
			"disease_name": disease.Nama,
			"disease_id": data.PenyakitID,
			"is_match": foundIndex != -1,
			"similarity": similarity,
		},
	})
}