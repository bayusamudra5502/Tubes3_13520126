package penyakit

import "github.com/gin-gonic/gin"

func PenyakitController(r *gin.Engine) {
	r.POST("/disease", AddPenyakit)
	r.GET("/disease", GetDisease)
	r.GET("/disease/:id", GetDiseaseById)
	r.DELETE("/disease/:id", DeleteDisease)
}
