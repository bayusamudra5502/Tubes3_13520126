package penyakit

import (
	"net/http"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/controller"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/db"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/model"
	"github.com/gin-gonic/gin"
)

func AddPenyakit(c *gin.Context){
	file, err := c.FormFile("file")
	name := c.PostForm("name")

	// Input Item Validation
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "KO",
			"message": "File is required",
			"data": err,
		})
		return
	}

	if name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "KO",
			"message": "Disease name is required",
			"data": err,
		})
		return
	}

	data, isValid := controller.FileValidation(c, file)

	if !isValid {
		return
	}

	// -- END of validation
	// Adding data to database
	record := model.Penyakit{
		Nama: name,
		DnaSequence: data,
	}

	result := db.GetDatabase().Create(&record)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "KO",
			"message": "Error inserting to database",
			"data": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"message": "Success",
		"data": gin.H{
			"filename": name,
			"filesize": file.Size,
			"id": record.ID,
		},
	})
}