package pemeriksaan

import (
	"net/http"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/controller"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/db"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/model"
	"github.com/gin-gonic/gin"
)

func GetCheckupHistory(ctx *gin.Context) {
	page, limit := controller.PaginationContext(ctx)

	if limit <= 0 || page < 0 {
		return
	}

	offset := page * limit

	var count int64
	var result []model.Pemeriksaan

	db.GetDatabase().Table("pemeriksaan").Count(&count)
	db.GetDatabase().Limit(limit).Offset(offset).Find(&result)

	var data []gin.H

	for _, i := range result {
		data = append(data, gin.H{
			"id":                i.ID,
			"name":              i.NamaPasien,
			"similarity":        i.Similarity,
			"disease_id":        i.PenyakitID,
			"disease_name":      i.Penyakit.Nama,
			"is_match":          i.Similarity == 1.0,
			"created_timestamp": i.CreatedAt,
			"updated_timestamp": i.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Success",
		"data": gin.H{
			"read_number":  len(data),
			"total_record": count,
			"data":         data,
		},
	})
}
