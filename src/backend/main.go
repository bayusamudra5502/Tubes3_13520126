package main

import (
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/controller/pemeriksaan"
	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/controller/penyakit"
	_ "github.com/bayusamudra5502/Tubes3_13520126/src/backend/db"
	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
)

func main() {
	env.Load()

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  "OK",
			"message": "success",
		})
	})

	penyakit.PenyakitController(r)
	pemeriksaan.PemeriksaanController(r)

	r.Run()
}
